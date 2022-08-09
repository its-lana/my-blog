package helper

import (
	"my-blog/model/domain"
	"my-blog/model/web"
)

func ToPostResponse(post domain.Post) web.PostResponse {
	return web.PostResponse{
		Id:   post.Id,
		Title: post.Title,
		Category: post.Category,
		Content: post.Content,
	}
}

func ToPostResponses(categories []domain.Post) []web.PostResponse {
	var postResponses []web.PostResponse
	for _, post := range categories {
		postResponses = append(postResponses, ToPostResponse(post))
	}
	return postResponses
}
