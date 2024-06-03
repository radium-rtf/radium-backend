package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/wave/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
	"github.com/uptrace/bun"
)

type Dialogue struct {
	db *bun.DB
}

func NewDialogueRepo(pg *postgres.Postgres) Dialogue {
	return Dialogue{db: pg.DB}
}

func (r Dialogue) Get(ctx context.Context, dialogueId uuid.UUID) (*entity.Dialogue, error) {
	var dialogue entity.Dialogue
	err := r.db.NewSelect().Model(&dialogue).
		Where("id = ?", dialogueId).
		Scan(ctx)
	return &dialogue, err
}

func (r Dialogue) GetByUsers(ctx context.Context, firstUserId, secondUserId uuid.UUID) (*entity.Dialogue, error) {
	var dialogue entity.Dialogue
	err := r.db.NewSelect().Model(&dialogue).
		Where("(first_user_id = ? AND second_user_id = ?) OR (first_user_id = ? AND second_user_id = ?)", firstUserId, secondUserId, secondUserId, firstUserId).
		Scan(ctx)
	return &dialogue, err
}

func (r Dialogue) GetAllByUserId(ctx context.Context, userId uuid.UUID) ([]*entity.Dialogue, error) {
	// TODO: потом нужно будет доставать это из юзера через join(Relation("Dialogues"))
	var dialogues []*entity.Dialogue
	err := r.db.NewSelect().Model(&dialogues).
		Where("first_user_id = ? OR second_user_id = ?", userId, userId).
		Scan(ctx)
	return dialogues, err
}

func (r Dialogue) Create(ctx context.Context, dialogue *entity.Dialogue) error {
	return r.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		_, err := tx.NewInsert().Model(dialogue).Exec(ctx)
		return err
	})
}
