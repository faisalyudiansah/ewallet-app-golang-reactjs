package repository

import (
	"context"
	"fmt"
	"testing"

	"ewallet-server-v2/internal/model"
	"ewallet-server-v2/internal/pkg/database"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestNewSourceOfFundRepository(t *testing.T) {
	db, _, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		DSN:                  "db",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	gormDB, _ := gorm.Open(dialector, &gorm.Config{})

	dep := NewSourceOfFundRepository(database.NewGormWrapper(gormDB))
	assert.NotNil(t, dep)
	assert.NotNil(t, dep.db)
}

func Test_sourceOfFundRepositoryPostgreSQL_GetAll(t *testing.T) {
	db, mockDB, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		DSN:                  "db",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	gormDB, _ := gorm.Open(dialector, &gorm.Config{})

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		mock    func(sqlmock.Sqlmock)
		want    []model.SourceOfFund
		wantErr error
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
			},
			mock: func(m sqlmock.Sqlmock) {
				rows := sqlmock.
					NewRows([]string{"source_of_fund_id", "source_name"}).
					AddRow(1, "Bank").AddRow(2, "Minimart")
				m.ExpectQuery("SELECT (.+)").WillReturnRows(rows)
			},
			want: []model.SourceOfFund{
				{
					SourceOfFundId: 1,
					Name:           "Bank",
				},
				{
					SourceOfFundId: 2,
					Name:           "Minimart",
				},
			},
			wantErr: nil,
		},
		{
			name: "error",
			args: args{
				ctx: context.Background(),
			},
			mock: func(m sqlmock.Sqlmock) {
				m.ExpectQuery("SELECT (.+)").WillReturnError(fmt.Errorf("error"))
			},
			want:    nil,
			wantErr: fmt.Errorf("error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &sourceOfFundRepositoryPostgreSQL{
				db: database.NewGormWrapper(gormDB),
			}

			tt.mock(mockDB)

			got, err := r.GetAll(tt.args.ctx)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_sourceOfFundRepositoryPostgreSQL_GetOneById(t *testing.T) {
	db, mockDB, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		DSN:                  "db",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	gormDB, _ := gorm.Open(dialector, &gorm.Config{})

	type args struct {
		ctx      context.Context
		sourceId int64
	}
	tests := []struct {
		name    string
		args    args
		mock    func(sqlmock.Sqlmock)
		want    *model.SourceOfFund
		wantErr error
	}{
		{
			name: "success",
			args: args{
				ctx:      context.Background(),
				sourceId: 3,
			},
			mock: func(m sqlmock.Sqlmock) {
				row := sqlmock.NewRows([]string{"source_of_fund_id", "source_name"}).AddRow(3, "Bank")
				m.ExpectQuery("SELECT (.+)").WillReturnRows(row)
			},
			want: &model.SourceOfFund{
				SourceOfFundId: 3,
				Name:           "Bank",
			},
			wantErr: nil,
		},
		{
			name: "error no record found",
			args: args{
				ctx:      context.Background(),
				sourceId: 3,
			},
			mock: func(m sqlmock.Sqlmock) {
				m.ExpectQuery("SELECT (.+)").WillReturnError(gorm.ErrRecordNotFound)
			},
			want:    nil,
			wantErr: nil,
		},
		{
			name: "other error",
			args: args{
				ctx:      context.Background(),
				sourceId: 3,
			},
			mock: func(m sqlmock.Sqlmock) {
				m.ExpectQuery("SELECT (.+)").WillReturnError(fmt.Errorf("error"))
			},
			want:    nil,
			wantErr: fmt.Errorf("error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &sourceOfFundRepositoryPostgreSQL{
				db: database.NewGormWrapper(gormDB),
			}

			tt.mock(mockDB)

			got, err := r.GetOneById(tt.args.ctx, tt.args.sourceId)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
