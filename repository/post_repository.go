package repository

import (
	"context"
	"database/sql"
	"my-blog/model/domain"
)

type PostRepository interface {
	Save(ctx context.Context, tx *sql.Tx, post domain.Post) domain.Post
	Update(ctx context.Context, tx *sql.Tx, post domain.Post) domain.Post
	Delete(ctx context.Context, tx *sql.Tx, post domain.Post)
	FindById(ctx context.Context, tx *sql.Tx, postId int) (domain.Post, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Post
}
