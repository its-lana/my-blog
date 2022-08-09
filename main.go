package main

import (
	"my-blog/app"
	"my-blog/controller"
	"my-blog/helper"
	"my-blog/middleware"
	"my-blog/repository"
	"my-blog/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	postRepository := repository.NewPostRepository()
	postService := service.NewPostService(postRepository, db, validate)
	postController := controller.NewPostController(postService)
	router := app.NewRouter(postController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
