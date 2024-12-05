package repository

import (
	"context"
	"fmt"
	"testing"
	"time"

	"ewallet-server-v2/internal/model"
	"ewallet-server-v2/internal/pkg/database"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestNewUserRepository(t *testing.T) {
	db, _, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		DSN:                  "db",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	gormDB, _ := gorm.Open(dialector, &gorm.Config{})

	dep := NewUserRepository(database.NewGormWrapper(gormDB))
	assert.NotNil(t, dep)
	assert.NotNil(t, dep.db)
}

func Test_userRepositoryPostgreSQL_CreateOne(t *testing.T) {
	db, mockDB, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		DSN:                  "db",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	gormDB, _ := gorm.Open(dialector, &gorm.Config{})

	type args struct {
		ctx  context.Context
		user model.User
	}
	tests := []struct {
		name    string
		args    args
		mock    func(sqlmock.Sqlmock)
		want    *model.User
		wantErr error
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				user: model.User{
					Name:  "lala",
					Email: "lala@mail.com",
				},
			},
			mock: func(m sqlmock.Sqlmock) {
				row := sqlmock.
					NewRows([]string{"user_id", "created_at", "updated_at"}).
					AddRow(
						1,
						time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
						time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
					)
				m.ExpectBegin()
				m.ExpectQuery("INSERT (.+)").WillReturnRows(row)
				m.ExpectCommit()
			},
			want: &model.User{
				UserId:    1,
				Name:      "lala",
				Email:     "lala@mail.com",
				CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			wantErr: nil,
		},
		{
			name: "error",
			args: args{
				ctx: context.Background(),
				user: model.User{
					Name:  "lala",
					Email: "lala@mail.com",
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
			r := &userRepositoryPostgreSQL{
				db: database.NewGormWrapper(gormDB),
			}

			tt.mock(mockDB)

			got, err := r.CreateOne(tt.args.ctx, tt.args.user)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_userRepositoryPostgreSQL_GetOneById(t *testing.T) {
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
		want    *model.User
		wantErr error
	}{
		{
			name: "success",
			args: args{
				ctx:    context.Background(),
				userId: 1,
			},
			mock: func(s sqlmock.Sqlmock) {
				row := sqlmock.NewRows([]string{"user_id", "user_name"}).AddRow(1, "lala")
				s.ExpectQuery("SELECT (.+)").WillReturnRows(row)
			},
			want: &model.User{
				UserId: 1,
				Name:   "lala",
			},
			wantErr: nil,
		},
		{
			name: "error record not found",
			args: args{
				ctx:    context.Background(),
				userId: 1,
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
				userId: 1,
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
			r := &userRepositoryPostgreSQL{
				db: database.NewGormWrapper(gormDB),
			}

			tt.mock(mockDB)

			got, err := r.GetOneById(tt.args.ctx, tt.args.userId)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_userRepositoryPostgreSQL_GetOneByEmail(t *testing.T) {
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
		email string
	}
	tests := []struct {
		name    string
		args    args
		mock    func(sqlmock.Sqlmock)
		want    *model.User
		wantErr error
	}{
		{
			name: "success",
			args: args{
				ctx:   context.Background(),
				email: "lala@mail.com",
			},
			mock: func(s sqlmock.Sqlmock) {
				row := sqlmock.
					NewRows([]string{"user_id", "user_name", "email"}).
					AddRow(1, "lala", "lala@mail.com")
				s.ExpectQuery("SELECT (.+)").WillReturnRows(row)
			},
			want: &model.User{
				UserId: 1,
				Name:   "lala",
				Email:  "lala@mail.com",
			},
			wantErr: nil,
		},
		{
			name: "error record not found",
			args: args{
				ctx:   context.Background(),
				email: "lala@mail.com",
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
				ctx:   context.Background(),
				email: "lala@mail.com",
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
			r := &userRepositoryPostgreSQL{
				db: database.NewGormWrapper(gormDB),
			}

			tt.mock(mockDB)

			got, err := r.GetOneByEmail(tt.args.ctx, tt.args.email)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_userRepositoryPostgreSQL_GetOneByWalletId(t *testing.T) {
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
		want    *model.User
		wantErr error
	}{
		{
			name: "success",
			args: args{
				ctx:      context.Background(),
				walletId: 3,
			},
			mock: func(m sqlmock.Sqlmock) {
				row := sqlmock.NewRows([]string{"user_id", "user_name"}).AddRow(1, "lala")
				m.ExpectQuery("SELECT (.+)").WillReturnRows(row)
			},
			want: &model.User{
				UserId: 1,
				Name:   "lala",
			},
			wantErr: nil,
		},
		{
			name: "error no record found",
			args: args{
				ctx:      context.Background(),
				walletId: 3,
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
				walletId: 3,
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
			r := &userRepositoryPostgreSQL{
				db: database.NewGormWrapper(gormDB),
			}

			tt.mock(mockDB)

			got, err := r.GetOneByWalletId(tt.args.ctx, tt.args.walletId)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_userRepositoryPostgreSQL_SaveOne(t *testing.T) {
	db, mockDB, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		DSN:                  "db",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	gormDB, _ := gorm.Open(dialector, &gorm.Config{})

	type args struct {
		ctx  context.Context
		user model.User
	}
	tests := []struct {
		name    string
		args    args
		mock    func(sqlmock.Sqlmock)
		update  bool
		want    *model.User
		wantErr error
	}{
		{
			name: "success - update",
			args: args{
				ctx: context.Background(),
				user: model.User{
					UserId: 1,
					Name:   "lala",
					Email:  "lala@mail.com",
				},
			},
			mock: func(s sqlmock.Sqlmock) {
				s.ExpectBegin()
				s.ExpectExec("UPDATE (.+)").WillReturnResult(sqlmock.NewResult(1, 1))
				s.ExpectCommit()
			},
			update: true,
			want: &model.User{
				UserId:    1,
				Name:      "lala",
				Email:     "lala@mail.com",
				CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			wantErr: nil,
		},
		{
			name: "success - insert",
			args: args{
				ctx: context.Background(),
				user: model.User{
					Name:  "lala",
					Email: "lala@mail.com",
				},
			},
			mock: func(s sqlmock.Sqlmock) {
				row := sqlmock.
					NewRows([]string{"user_id", "user_name", "email", "created_at", "updated_at"}).
					AddRow(
						1,
						"lala",
						"lala@mail.com",
						time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
						time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
					)
				s.ExpectBegin()
				s.ExpectQuery("INSERT (.+)").WillReturnRows(row)
				s.ExpectCommit()
			},
			update: false,
			want: &model.User{
				UserId:    1,
				Name:      "lala",
				Email:     "lala@mail.com",
				CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			wantErr: nil,
		},
		{
			name: "error when insert",
			args: args{
				ctx: context.Background(),
				user: model.User{
					Name:  "lala",
					Email: "lala@mail.com",
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
			r := &userRepositoryPostgreSQL{
				db: database.NewGormWrapper(gormDB),
			}

			tt.mock(mockDB)

			got, err := r.SaveOne(tt.args.ctx, tt.args.user)

			if tt.update {
				assert.Equal(t, tt.want.UserId, got.UserId)
				assert.Equal(t, tt.want.Name, got.Name)
				assert.Equal(t, tt.want.Email, got.Email)
				assert.NotEqual(t, time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC), got.UserId)
			} else {
				assert.Equal(t, tt.want, got)
			}
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
