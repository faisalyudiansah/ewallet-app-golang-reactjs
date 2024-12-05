package repositories

import (
	"context"
	"database/sql"
	"fmt"

	helpercontext "git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/helpers/helper_context"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/models"
)

type GameRepositoryInterface interface {
	CreateGameBox(context.Context, int64, []int) error
	GetGameBox(context.Context, int64) (*models.GameBox, error)
	ChangeStatusBox(context.Context, int64) error
}

type GameRepositoryImplementation struct {
	db *sql.DB
}

func NewGameRepositoryImplementation(db *sql.DB) *GameRepositoryImplementation {
	return &GameRepositoryImplementation{
		db: db,
	}
}

func (gr *GameRepositoryImplementation) CreateGameBox(ctx context.Context, userID int64, boxes []int) error {
	sql := `
		INSERT INTO game_boxs 
			(user_id, is_open, box1, box2, box3, box4, box5, box6, box7, box8, box9, created_at, updated_at) 
		VALUES ($1, TRUE, $2, $3, $4, $5, $6, $7, $8, $9, $10, NOW(), NOW())
	`
	var err error
	txFromCtx := helpercontext.GetTx(ctx)
	if txFromCtx != nil {
		_, err = txFromCtx.ExecContext(ctx, sql, userID, boxes[0], boxes[1], boxes[2], boxes[3], boxes[4], boxes[5], boxes[6], boxes[7], boxes[8])
	} else {
		_, err = gr.db.ExecContext(ctx, sql, userID, boxes[0], boxes[1], boxes[2], boxes[3], boxes[4], boxes[5], boxes[6], boxes[7], boxes[8])
	}
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (gr *GameRepositoryImplementation) GetGameBox(ctx context.Context, userId int64) (*models.GameBox, error) {
	sql := `
		SELECT 
			id, 
			user_id, 
			is_open, 
			box1, box2, box3, box4, box5, box6, box7, box8, box9 
		FROM game_boxs 
		WHERE user_id = $1 AND is_open = TRUE AND deleted_at IS NULL
		ORDER BY created_at DESC LIMIT 1
	`
	var box models.GameBox
	var err error
	txFromCtx := helpercontext.GetTx(ctx)
	if txFromCtx != nil {
		err = txFromCtx.QueryRowContext(ctx, sql, userId).Scan(
			&box.ID,
			&box.UserID,
			&box.IsOpen,
			&box.Box1,
			&box.Box2,
			&box.Box3,
			&box.Box4,
			&box.Box5,
			&box.Box6,
			&box.Box7,
			&box.Box8,
			&box.Box9,
		)
	} else {
		err = gr.db.QueryRowContext(ctx, sql, userId).Scan(
			&box.ID,
			&box.UserID,
			&box.IsOpen,
			&box.Box1,
			&box.Box2,
			&box.Box3,
			&box.Box4,
			&box.Box5,
			&box.Box6,
			&box.Box7,
			&box.Box8,
			&box.Box9,
		)
	}
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &box, nil
}

func (gr *GameRepositoryImplementation) ChangeStatusBox(ctx context.Context, boxId int64) error {
	query := `
		UPDATE game_boxs SET 
		is_open = FALSE,
		updated_at = NOW()
		WHERE id = $1`
	var err error
	txFromCtx := helpercontext.GetTx(ctx)
	if txFromCtx != nil {
		_, err = txFromCtx.ExecContext(ctx, query, boxId)
	} else {
		_, err = gr.db.ExecContext(ctx, query, boxId)
	}
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
