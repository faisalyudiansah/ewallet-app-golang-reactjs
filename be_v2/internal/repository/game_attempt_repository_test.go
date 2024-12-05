package repository

import (
	"context"
	"fmt"
	"testing"
	"time"

	"ewallet-server-v2/internal/model"
	"ewallet-server-v2/internal/pkg/database"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestNewGameAttemptRepository(t *testing.T) {
	db, _, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		DSN:                  "db",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	gormDB, _ := gorm.Open(dialector, &gorm.Config{})

	dep := NewGameAttemptRepository(database.NewGormWrapper(gormDB))
	assert.NotNil(t, dep)
	assert.NotNil(t, dep.db)
}

func Test_gameAttemptRepositoryPostgreSQL_GetCountByWalletId(t *testing.T) {
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
		walletId int64
	}
	tests := []struct {
		name    string
		args    args
		mock    func(sqlmock.Sqlmock)
		want    int64
		wantErr error
	}{
		{
			name: "success",
			args: args{
				ctx:      context.Background(),
				walletId: 5,
			},
			mock: func(m sqlmock.Sqlmock) {
				row := sqlmock.NewRows([]string{"count"}).AddRow(2)
				m.ExpectQuery("SELECT (.+)").WithArgs(5).WillReturnRows(row)
			},
			want:    2,
			wantErr: nil,
		},
		{
			name: "error",
			args: args{
				ctx:      context.Background(),
				walletId: 5,
			},
			mock: func(m sqlmock.Sqlmock) {
				m.ExpectQuery("SELECT (.+)").WillReturnError(fmt.Errorf("error"))
			},
			want:    0,
			wantErr: fmt.Errorf("error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &gameAttemptRepositoryPostgreSQL{
				db: database.NewGormWrapper(gormDB),
			}

			tt.mock(mockDB)

			got, err := r.GetCountByWalletId(tt.args.ctx, tt.args.walletId)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_gameAttemptRepositoryPostgreSQL_CreateOne(t *testing.T) {
	db, mockDB, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		DSN:                  "db",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	gormDB, _ := gorm.Open(dialector, &gorm.Config{})

	type args struct {
		ctx         context.Context
		gameAttempt model.GameAttempt
	}
	tests := []struct {
		name    string
		args    args
		mock    func(sqlmock.Sqlmock)
		want    *model.GameAttempt
		wantErr error
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				gameAttempt: model.GameAttempt{
					GameAttemptId: 1,
					WalletId:      1,
					Amount:        decimal.NewFromInt(5000),
				},
			},
			mock: func(m sqlmock.Sqlmock) {
				m.ExpectBegin()
				row := sqlmock.NewRows(
					[]string{
						"game_attempt_id",
						"wallet_id",
						"amount",
						"game_boxes_id",
						"created_at",
						"updated_at",
						"deleted_at",
					},
				).AddRow(
					1,
					1,
					5000,
					1,
					time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
				)
				m.ExpectQuery("INSERT (.+)").WillReturnRows(row)
				m.ExpectCommit()
			},
			want: &model.GameAttempt{
				GameAttemptId: 1,
				WalletId:      1,
				Amount:        decimal.NewFromInt(5000),
				GameBoxesid:   1,
				CreatedAt:     time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt:     time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
				DeletedAt: gorm.DeletedAt{
					Time:  time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
					Valid: true,
				},
			},
			wantErr: nil,
		},
		{
			name: "error",
			args: args{
				ctx: context.Background(),
				gameAttempt: model.GameAttempt{
					GameAttemptId: 1,
					WalletId:      1,
					Amount:        decimal.NewFromInt(5000),
				},
			},
			mock: func(m sqlmock.Sqlmock) {
				m.ExpectBegin()
				m.ExpectQuery("INSERT (.+)").WillReturnError(fmt.Errorf("error"))
				m.ExpectRollback()
			},
			want:    nil,
			wantErr: fmt.Errorf("error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &gameAttemptRepositoryPostgreSQL{
				db: database.NewGormWrapper(gormDB),
			}

			tt.mock(mockDB)

			got, err := r.CreateOne(tt.args.ctx, tt.args.gameAttempt)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
