package postgres

import (
	"context"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
	"github.com/uptrace/bun"
)

type File struct {
	db *bun.DB
}

func NewFileRepo(pg *postgres.Postgres) File {
	return File{db: pg.DB}
}

func (r File) Create(ctx context.Context, file *entity.File) error {
	_, err := r.db.NewInsert().Model(file).Exec(ctx)
	return err
}

func (r File) Get(ctx context.Context, url string) (*entity.File, error) {
	var file = new(entity.File)
	err := r.db.NewSelect().
		Model(file).
		Where("url = ?", url).
		Scan(ctx)
	return file, err
}
