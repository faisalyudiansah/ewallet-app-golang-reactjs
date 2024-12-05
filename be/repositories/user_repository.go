package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/dtos"
	helpercontext "git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/helpers/helper_context"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/models"
)

type UserRepository interface {
	IsEmailAlreadyRegistered(context.Context, string) bool
	PostUser(context.Context, dtos.RequestRegisterUser, string) (*models.User, error)
	GetUserByEmail(context.Context, string) (*models.User, error)
	GetUserById(context.Context, int64) (*models.User, error)
	PutResetPassword(context.Context, string, int64) (*models.User, error)
	PutAttemptGame(context.Context, int, int64) (*models.User, error)
}

type UserRepositoryImplementation struct {
	db *sql.DB
}

func NewUserRepositoryImplementation(db *sql.DB) *UserRepositoryImplementation {
	return &UserRepositoryImplementation{
		db: db,
	}
}

func (us *UserRepositoryImplementation) IsEmailAlreadyRegistered(ctx context.Context, emailInput string) bool {
	sql := `
	SELECT
		u.id
	FROM users u
	WHERE email = $1 AND deleted_at IS NULL;
`
	var user models.User
	txFromCtx := helpercontext.GetTx(ctx)
	if txFromCtx != nil {
		txFromCtx.QueryRowContext(ctx, sql, emailInput).Scan(
			&user.ID,
		)
	} else {
		us.db.QueryRowContext(ctx, sql, emailInput).Scan(
			&user.ID,
		)
	}
	return user.ID != 0
}

func (us *UserRepositoryImplementation) GetUserById(ctx context.Context, userId int64) (*models.User, error) {
	sql := `
	SELECT
		u.id,
		u.email,
		u.password,
		u.chance_game,
		u.fullname,
		u.birthdate,
		u.created_at,
		u.updated_at,
		u.deleted_at
	FROM users u
	WHERE id = $1 AND deleted_at IS NULL;
	`
	var user models.User
	var err error
	txFromCtx := helpercontext.GetTx(ctx)
	if txFromCtx != nil {
		err = txFromCtx.QueryRowContext(ctx, sql, userId).Scan(
			&user.ID,
			&user.Email,
			&user.Password,
			&user.ChanceGame,
			&user.FullName,
			&user.BirthDate,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeleteAt,
		)
	} else {
		err = us.db.QueryRowContext(ctx, sql, userId).Scan(
			&user.ID,
			&user.Email,
			&user.Password,
			&user.ChanceGame,
			&user.FullName,
			&user.BirthDate,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeleteAt,
		)
	}
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &user, nil
}

func (us *UserRepositoryImplementation) PostUser(ctx context.Context, reqBody dtos.RequestRegisterUser, hashPassword string) (*models.User, error) {
	sql := `
		INSERT INTO users (email, password, chance_game, fullname, birthdate, created_at, updated_at) VALUES 
		($1, $2, 0, $3, $4, NOW(), NOW())
		RETURNING id, email, password, chance_game, fullname, birthdate, created_at, updated_at, deleted_at;
	`
	var user models.User
	var err error
	txFromCtx := helpercontext.GetTx(ctx)
	if txFromCtx != nil {
		err = txFromCtx.QueryRowContext(ctx, sql, reqBody.Email, hashPassword, reqBody.FullName, reqBody.BirthDate).Scan(
			&user.ID,
			&user.Email,
			&user.Password,
			&user.ChanceGame,
			&user.FullName,
			&user.BirthDate,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeleteAt,
		)
	} else {
		err = us.db.QueryRowContext(ctx, sql, reqBody.Email, hashPassword, reqBody.FullName, reqBody.BirthDate).Scan(
			&user.ID,
			&user.Email,
			&user.Password,
			&user.ChanceGame,
			&user.FullName,
			&user.BirthDate,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeleteAt,
		)
	}
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &user, nil
}

func (us *UserRepositoryImplementation) GetUserByEmail(ctx context.Context, emailInput string) (*models.User, error) {
	sql := `
		SELECT
			u.id,
			u.email,
			u.password,
			u.chance_game,
			u.fullname,
			u.birthdate,
			u.created_at,
			u.updated_at,
			u.deleted_at
		FROM users u
		WHERE email = $1 AND deleted_at IS NULL;
	`
	var user models.User
	var err error
	txFromCtx := helpercontext.GetTx(ctx)
	if txFromCtx != nil {
		err = txFromCtx.QueryRowContext(ctx, sql, emailInput).Scan(
			&user.ID,
			&user.Email,
			&user.Password,
			&user.ChanceGame,
			&user.FullName,
			&user.BirthDate,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeleteAt,
		)
	} else {
		err = us.db.QueryRowContext(ctx, sql, emailInput).Scan(
			&user.ID,
			&user.Email,
			&user.Password,
			&user.ChanceGame,
			&user.FullName,
			&user.BirthDate,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeleteAt,
		)
	}
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &user, nil
}

func (us *UserRepositoryImplementation) PutResetPassword(ctx context.Context, newPassword string, userId int64) (*models.User, error) {
	sql := `
		UPDATE users SET 
		password = $1,
		updated_at = NOW()
		WHERE id = $2
		RETURNING id, email, password, chance_game, fullname, birthdate, created_at, updated_at, deleted_at;
	`
	var user models.User
	var err error
	txFromCtx := helpercontext.GetTx(ctx)
	if txFromCtx != nil {
		err = txFromCtx.QueryRowContext(ctx, sql, newPassword, userId).Scan(
			&user.ID,
			&user.Email,
			&user.Password,
			&user.ChanceGame,
			&user.FullName,
			&user.BirthDate,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeleteAt,
		)
	} else {
		err = us.db.QueryRowContext(ctx, sql, newPassword, userId).Scan(
			&user.ID,
			&user.Email,
			&user.Password,
			&user.ChanceGame,
			&user.FullName,
			&user.BirthDate,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeleteAt,
		)
	}
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &user, nil
}

func (us *UserRepositoryImplementation) PutAttemptGame(ctx context.Context, chance int, userId int64) (*models.User, error) {
	sql := `
		UPDATE users SET 
			chance_game = $1,
			updated_at = NOW()
		WHERE id = $2
		RETURNING id, email, password, chance_game, fullname, birthdate, created_at, updated_at, deleted_at;
	`
	var user models.User
	var err error
	txFromCtx := helpercontext.GetTx(ctx)
	if txFromCtx != nil {
		err = txFromCtx.QueryRowContext(ctx, sql, chance, userId).Scan(
			&user.ID,
			&user.Email,
			&user.Password,
			&user.ChanceGame,
			&user.FullName,
			&user.BirthDate,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeleteAt,
		)
	} else {
		err = us.db.QueryRowContext(ctx, sql, chance, userId).Scan(
			&user.ID,
			&user.Email,
			&user.Password,
			&user.ChanceGame,
			&user.FullName,
			&user.BirthDate,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeleteAt,
		)
	}
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &user, nil
}
