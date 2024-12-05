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

func TestNewWalletRepository(t *testing.T) {
	db, _, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		DSN:                  "db",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	gormDB, _ := gorm.Open(dialector, &gorm.Config{})

	dep := NewWalletRepository(database.NewGormWrapper(gormDB))
	assert.NotNil(t, dep)
	assert.NotNil(t, dep.db)
}

func Test_walletRepositoryPostgreSQL_CreateOne(t *testing.T) {
	db, mockDB, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		DSN:                  "db",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	gormDB, _ := gorm.Open(dialector, &gorm.Config{})

	type args struct {
		ctx    context.Context
		wallet model.Wallet
	}
	tests := []struct {
		name    string
		args    args
		mock    func(sqlmock.Sqlmock)
		want    *model.Wallet
		wantErr error
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				wallet: model.Wallet{
					WalletNumber: "777001",
				},
			},
			mock: func(s sqlmock.Sqlmock) {
				row := sqlmock.
					NewRows([]string{"wallet_id", "wallet_number", "created_at", "updated_at"}).
					AddRow(
						"1",
						"777001",
						time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
						time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
					)

				s.ExpectBegin()
				s.ExpectQuery("INSERT (.+)").WillReturnRows(row)
				s.ExpectCommit()
			},
			want: &model.Wallet{
				WalletId:     1,
				WalletNumber: "777001",
				CreatedAt:    time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt:    time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			wantErr: nil,
		},
		{
			name: "error",
			args: args{
				ctx: context.Background(),
				wallet: model.Wallet{
					WalletNumber: "777001",
				},
			},
			mock: func(s sqlmock.Sqlmock) {
				s.ExpectBegin()
				s.ExpectQuery("INSERT (.+)").WillReturnError(fmt.Errorf("error"))
				s.ExpectRollback()
			},
			want:    nil,
			wantErr: fmt.Errorf("error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &walletRepositoryPostgreSQL{
				db: database.NewGormWrapper(gormDB),
			}

			tt.mock(mockDB)

			got, err := r.CreateOne(tt.args.ctx, tt.args.wallet)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_walletRepositoryPostgreSQL_SaveOne(t *testing.T) {
	db, mockDB, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		DSN:                  "db",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	gormDB, _ := gorm.Open(dialector, &gorm.Config{})

	type args struct {
		ctx    context.Context
		wallet model.Wallet
	}
	tests := []struct {
		name    string
		args    args
		mock    func(sqlmock.Sqlmock)
		update  bool
		want    *model.Wallet
		wantErr error
	}{
		{
			name: "success - update",
			args: args{
				ctx: context.Background(),
				wallet: model.Wallet{
					WalletId:     1,
					WalletNumber: "777001",
					CreatedAt:    time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			mock: func(s sqlmock.Sqlmock) {
				s.ExpectBegin()
				s.ExpectExec("UPDATE (.+)").WillReturnResult(sqlmock.NewResult(1, 1))
				s.ExpectCommit()
			},
			update: true,
			want: &model.Wallet{
				WalletId:     1,
				WalletNumber: "777001",
				CreatedAt:    time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			wantErr: nil,
		},
		{
			name: "success - insert",
			args: args{
				ctx: context.Background(),
				wallet: model.Wallet{
					WalletNumber: "777001",
					Amount:       decimal.NewFromInt(5000),
				},
			},
			mock: func(s sqlmock.Sqlmock) {
				row := sqlmock.
					NewRows([]string{"wallet_id", "wallet_number", "amount", "created_at", "updated_at"}).
					AddRow(
						1,
						"777001",
						decimal.NewFromInt(5000),
						time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
						time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
					)

				s.ExpectBegin()
				s.ExpectQuery("INSERT (.+)").WillReturnRows(row)
				s.ExpectCommit()
			},
			update: false,
			want: &model.Wallet{
				WalletId:     1,
				WalletNumber: "777001",
				Amount:       decimal.NewFromInt(5000),
				CreatedAt:    time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt:    time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			wantErr: nil,
		},
		{
			name: "error when insert",
			args: args{
				ctx: context.Background(),
				wallet: model.Wallet{
					WalletNumber: "777001",
				},
			},
			mock: func(s sqlmock.Sqlmock) {
				s.ExpectBegin()
				s.ExpectQuery("INSERT (.+)").WillReturnError(fmt.Errorf("error"))
				s.ExpectRollback()
			},
			update:  false,
			want:    nil,
			wantErr: fmt.Errorf("error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &walletRepositoryPostgreSQL{
				db: database.NewGormWrapper(gormDB),
			}

			tt.mock(mockDB)

			got, err := r.SaveOne(tt.args.ctx, tt.args.wallet)

			if tt.update {
				assert.Equal(t, tt.want.WalletId, got.WalletId)
				assert.Equal(t, tt.want.WalletNumber, got.WalletNumber)
				assert.Equal(t, tt.want.CreatedAt, got.CreatedAt)
			} else {
				assert.Equal(t, tt.want, got)
			}
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_walletRepositoryPostgreSQL_GetOneByIdWithLock(t *testing.T) {
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
		want    *model.Wallet
		wantErr error
	}{
		{
			name: "success",
			args: args{
				ctx:      context.Background(),
				walletId: 1,
			},
			mock: func(s sqlmock.Sqlmock) {
				row := sqlmock.NewRows([]string{"wallet_id", "amount"}).AddRow(1, 5000)

				s.ExpectQuery("SELECT (.+)").WillReturnRows(row)
			},
			want: &model.Wallet{
				WalletId: 1,
				Amount:   decimal.NewFromInt(5000),
			},
			wantErr: nil,
		},
		{
			name: "error record not found",
			args: args{
				ctx:      context.Background(),
				walletId: 1,
			},
			mock: func(s sqlmock.Sqlmock) {
				s.ExpectQuery("SELECT (.+)").WillReturnError(gorm.ErrRecordNotFound)
			},
			want:    nil,
			wantErr: nil,
		},
		{
			name: "other error",
			args: args{
				ctx:      context.Background(),
				walletId: 1,
			},
			mock: func(s sqlmock.Sqlmock) {
				s.ExpectQuery("SELECT (.+)").WillReturnError(fmt.Errorf("error"))
			},
			want:    nil,
			wantErr: fmt.Errorf("error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &walletRepositoryPostgreSQL{
				db: database.NewGormWrapper(gormDB),
			}

			tt.mock(mockDB)

			got, err := r.GetOneByIdWithLock(tt.args.ctx, tt.args.walletId)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_walletRepositoryPostgreSQL_GetOneByNumberWithLock(t *testing.T) {
	db, mockDB, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		DSN:                  "db",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	gormDB, _ := gorm.Open(dialector, &gorm.Config{})

	type args struct {
		ctx          context.Context
		walletNumber string
	}
	tests := []struct {
		name    string
		args    args
		mock    func(sqlmock.Sqlmock)
		want    *model.Wallet
		wantErr error
	}{
		{
			name: "success",
			args: args{
				ctx:          context.Background(),
				walletNumber: "777001",
			},
			mock: func(s sqlmock.Sqlmock) {
				row := sqlmock.NewRows([]string{"wallet_number", "amount"}).AddRow("777001", 5000)
				s.ExpectQuery("SELECT (.+)").WillReturnRows(row)
			},
			want: &model.Wallet{
				WalletNumber: "777001",
				Amount:       decimal.NewFromInt(5000),
			},
			wantErr: nil,
		},
		{
			name: "error no record found",
			args: args{
				ctx:          context.Background(),
				walletNumber: "777001",
			},
			mock: func(s sqlmock.Sqlmock) {
				s.ExpectQuery("SELECT (.+)").WillReturnError(gorm.ErrRecordNotFound)
			},
			want:    nil,
			wantErr: nil,
		},
		{
			name: "other error",
			args: args{
				ctx:          context.Background(),
				walletNumber: "777001",
			},
			mock: func(s sqlmock.Sqlmock) {
				s.ExpectQuery("SELECT (.+)").WillReturnError(fmt.Errorf("error"))
			},
			want:    nil,
			wantErr: fmt.Errorf("error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &walletRepositoryPostgreSQL{
				db: database.NewGormWrapper(gormDB),
			}

			tt.mock(mockDB)

			got, err := r.GetOneByNumberWithLock(tt.args.ctx, tt.args.walletNumber)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_walletRepositoryPostgreSQL_GetOneByUserId(t *testing.T) {
	db, mockDB, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		DSN:                  "db",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	gormDB, _ := gorm.Open(dialector, &gorm.Config{})

	type args struct {
		ctx    context.Context
		userId int64
	}
	tests := []struct {
		name    string
		args    args
		mock    func(sqlmock.Sqlmock)
		want    *model.Wallet
		wantErr error
	}{
		{
			name: "success",
			args: args{
				ctx:    context.Background(),
				userId: 4,
			},
			mock: func(s sqlmock.Sqlmock) {
				row := sqlmock.NewRows([]string{"wallet_id", "wallet_number"}).AddRow(1, "777001")
				s.ExpectQuery("SELECT (.+)").WillReturnRows(row)
			},
			want: &model.Wallet{
				WalletId:     1,
				WalletNumber: "777001",
			},
			wantErr: nil,
		},
		{
			name: "error no record found",
			args: args{
				ctx:    context.Background(),
				userId: 4,
			},
			mock: func(s sqlmock.Sqlmock) {
				s.ExpectQuery("SELECT (.+)").WillReturnError(gorm.ErrRecordNotFound)
			},
			want:    nil,
			wantErr: nil,
		},
		{
			name: "other error",
			args: args{
				ctx:    context.Background(),
				userId: 4,
			},
			mock: func(s sqlmock.Sqlmock) {
				s.ExpectQuery("SELECT (.+)").WillReturnError(fmt.Errorf("error"))
			},
			want:    nil,
			wantErr: fmt.Errorf("error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &walletRepositoryPostgreSQL{
				db: database.NewGormWrapper(gormDB),
			}

			tt.mock(mockDB)

			got, err := r.GetOneByUserId(tt.args.ctx, tt.args.userId)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
