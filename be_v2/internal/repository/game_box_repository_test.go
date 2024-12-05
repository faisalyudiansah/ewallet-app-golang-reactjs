package repository

import (
	"context"
	"fmt"
	"testing"

	"ewallet-server-v2/internal/model"
	"ewallet-server-v2/internal/pkg/database"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestNewGameBoxRepository(t *testing.T) {
	db, _, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		DSN:                  "db",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	gormDB, _ := gorm.Open(dialector, &gorm.Config{})

	dep := NewGameBoxRepository(database.NewGormWrapper(gormDB))
	assert.NotNil(t, dep)
	assert.NotNil(t, dep.db)
}

func Test_gameBoxRepositoryPostgreSQL_GetAll(t *testing.T) {
	db, mockDB, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		DSN:                  "db",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	gormDB, _ := gorm.Open(dialector, &gorm.Config{})

	type args struct {
		ctx   context.Context
		limit int
	}
	tests := []struct {
		name    string
		args    args
		mock    func(sqlmock.Sqlmock)
		want    []model.GameBox
		wantErr error
	}{
		{
			name: "success",
			args: args{
				ctx:   context.Background(),
				limit: 2,
			},
			mock: func(m sqlmock.Sqlmock) {
				row := sqlmock.NewRows([]string{"game_boxes_id", "amount"}).
					AddRow(1, 5000).
					AddRow(2, 3400)
				m.ExpectQuery("SELECT (.+)").WillReturnRows(row)
			},
			want: []model.GameBox{
				{
					GameBoxId: 1,
					Amount:    decimal.NewFromInt(5000),
				},
				{
					GameBoxId: 2,
					Amount:    decimal.NewFromInt(3400),
				},
			},
			wantErr: nil,
		},
		{
			name: "error",
			args: args{
				ctx:   context.Background(),
				limit: 2,
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
			r := &gameBoxRepositoryPostgreSQL{
				db: database.NewGormWrapper(gormDB),
			}

			tt.mock(mockDB)

			got, err := r.GetAll(tt.args.ctx, tt.args.limit)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_gameBoxRepositoryPostgreSQL_GetOneById(t *testing.T) {
	db, mockDB, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		DSN:                  "db",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	gormDB, _ := gorm.Open(dialector, &gorm.Config{})

	type args struct {
		ctx   context.Context
		boxId int64
	}
	tests := []struct {
		name    string
		args    args
		mock    func(sqlmock.Sqlmock)
		want    *model.GameBox
		wantErr error
	}{
		{
			name: "success",
			args: args{
				ctx:   context.Background(),
				boxId: 3,
			},
			mock: func(m sqlmock.Sqlmock) {
				row := sqlmock.NewRows([]string{"game_boxes_id", "amount"}).AddRow(3, 3500)
				m.ExpectQuery("SELECT (.+)").WithArgs(3).WillReturnRows(row)
			},
			want: &model.GameBox{
				GameBoxId: 3,
				Amount:    decimal.NewFromInt(3500),
			},
			wantErr: nil,
		},
		{
			name: "error no record",
			args: args{
				ctx:   context.Background(),
				boxId: 3,
			},
			mock: func(m sqlmock.Sqlmock) {
				m.ExpectQuery("SELECT (.+)").WithArgs(3).WillReturnError(gorm.ErrRecordNotFound)
			},
			want:    nil,
			wantErr: nil,
		},
		{
			name: "other error",
			args: args{
				ctx:   context.Background(),
				boxId: 3,
			},
			mock: func(m sqlmock.Sqlmock) {
				m.ExpectQuery("SELECT (.+)").WithArgs(3).WillReturnError(fmt.Errorf("error"))
			},
			want:    nil,
			wantErr: fmt.Errorf("error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &gameBoxRepositoryPostgreSQL{
				db: database.NewGormWrapper(gormDB),
			}

			tt.mock(mockDB)

			got, err := r.GetOneById(tt.args.ctx, tt.args.boxId)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
