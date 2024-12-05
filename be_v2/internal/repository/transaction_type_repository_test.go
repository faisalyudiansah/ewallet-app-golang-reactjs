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

func TestNewTransactionTypeRepository(t *testing.T) {
	db, _, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		DSN:                  "db",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	gormDB, _ := gorm.Open(dialector, &gorm.Config{})

	dep := NewTransactionTypeRepository(database.NewGormWrapper(gormDB))
	assert.NotNil(t, dep)
	assert.NotNil(t, dep.db)
}

func Test_transactionTypeRepositoryPostgreSQL_GetAll(t *testing.T) {
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
		want    []model.TransactionType
		wantErr error
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
			},
			mock: func(m sqlmock.Sqlmock) {
				rows := sqlmock.
					NewRows([]string{"transaction_type_id", "type_name"}).
					AddRow(1, "transfer").
					AddRow(2, "topup")
				m.ExpectQuery("SELECT (.+)").WillReturnRows(rows)
			},
			want: []model.TransactionType{
				{
					TransactionTypeId: 1,
					TypeName:          "transfer",
				},
				{
					TransactionTypeId: 2,
					TypeName:          "topup",
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
			r := &transactionTypeRepositoryPostgreSQL{
				db: database.NewGormWrapper(gormDB),
			}

			tt.mock(mockDB)

			got, err := r.GetAll(tt.args.ctx)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
