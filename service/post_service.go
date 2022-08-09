package service

import (
	"context"
	"my-blog/model/web"
)

type PostService interface {
	Create(ctx context.Context, request web.PostCreateRequest) web.PostResponse
	Update(ctx context.Context, request web.PostUpdateRequest) web.PostResponse
	Delete(ctx context.Context, postId int)
	FindById(ctx context.Context, postId int) web.PostResponse
	FindAll(ctx context.Context) []web.PostResponse
}
