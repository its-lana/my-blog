package repository

import (
	"context"
	"database/sql"
	"errors"
	"my-blog/helper"
	"my-blog/model/domain"
)

type PostRepositoryImpl struct {
}

func NewPostRepository() PostRepository {
	return &PostRepositoryImpl{}
}

func (repository *PostRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, post domain.Post) domain.Post {
	SQL := "insert into posts(title, category, content) values (?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, post.Title, post.Category, post.Content)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	post.Id = int(id)
	return post
}

func (repository *PostRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, post domain.Post) domain.Post {
	SQL := "update posts set title=?, category=?, content=? where id=?"
	_, err := tx.ExecContext(ctx, SQL, post.Title, post.Category, post.Content, post.Id)
	helper.PanicIfError(err)

	return post
}

func (repository *PostRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, post domain.Post) {
	SQL := "delete from posts where id = ?"
	_, err := tx.ExecContext(ctx, SQL, post.Id)
	helper.PanicIfError(err)
}

func (repository *PostRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, postId int) (domain.Post, error) {
	SQL := "select id, title, category, content from posts where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, postId)
	helper.PanicIfError(err)
	defer rows.Close()

	post := domain.Post{}
	if rows.Next() {
		err := rows.Scan(&post.Id, &post.Title, &post.Category, &post.Content)
		helper.PanicIfError(err)
		return post, nil
	} else {
		return post, errors.New("post is not found")
	}
}

func (repository *PostRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Post {
	SQL := "select id, title, category, content from posts"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var categories []domain.Post
	for rows.Next() {
		post := domain.Post{}
		err := rows.Scan(&post.Id, &post.Title, &post.Category, &post.Content)
		helper.PanicIfError(err)
		categories = append(categories, post)
	}
	return categories
}
