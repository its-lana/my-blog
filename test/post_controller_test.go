package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"my-blog/app"
	"my-blog/controller"
	"my-blog/helper"
	"my-blog/middleware"
	"my-blog/model/domain"
	"my-blog/repository"
	"my-blog/service"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() *sql.DB {
	// change <Password> with your MySql password
	db, err := sql.Open("mysql", "root:<Password>@tcp(localhost:3306)/myblog_test")
	helper.PanicIfError(err)


	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	postRepository := repository.NewPostRepository()
	postService := service.NewPostService(postRepository, db, validate)
	postController := controller.NewPostController(postService)
	router := app.NewRouter(postController)

	return middleware.NewAuthMiddleware(router)
}

func truncatePost(db *sql.DB) {
	db.Exec("TRUNCATE post")
}

func TestCreatePostSuccess(t *testing.T) {
	db := setupTestDB()
	truncatePost(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"title" : "Review Smartphone Samsung S21 Ultra", "category": "Smartphone", "Content":"This is review about Smartphone Samsung S21 Ultra"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/posts", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "Review Smartphone Samsung S21 Ultra", responseBody["data"].(map[string]interface{})["title"])
	assert.Equal(t, "Smartphone", responseBody["data"].(map[string]interface{})["category"])
	assert.Equal(t, "This is review about Smartphone Samsung S21 Ultra", responseBody["data"].(map[string]interface{})["content"])
}

func TestCreatePostFailed(t *testing.T) {
	db := setupTestDB()
	truncatePost(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"title" : "", "category" : "", "content" : ""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/posts", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestUpdatePostSuccess(t *testing.T) {
	db := setupTestDB()
	truncatePost(db)

	tx, _ := db.Begin()
	postRepository := repository.NewPostRepository()
	post := postRepository.Save(context.Background(), tx, domain.Post{
		Title: "Review Laptop Legion Slim 5",
		Category: "Laptop",
		Content: "This is review about Legion Slim 5",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"title" : "Review Laptop Legion Slim 7", "category": "Laptop", "content": "This is review about Legion Slim 7"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/posts/"+strconv.Itoa(post.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, post.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "Review Laptop Legion Slim 7", responseBody["data"].(map[string]interface{})["title"])
	assert.Equal(t, "Laptop", responseBody["data"].(map[string]interface{})["category"])
	assert.Equal(t, "This is review about Legion Slim 7", responseBody["data"].(map[string]interface{})["content"])
}

func TestUpdatePostFailed(t *testing.T) {
	db := setupTestDB()
	truncatePost(db)

	tx, _ := db.Begin()
	postRepository := repository.NewPostRepository()
	post := postRepository.Save(context.Background(), tx, domain.Post{
		Title: "Review Laptop Legion Slim 7",
		Category: "Laptop",
		Content: "This is review about Legion Slim 7",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"title" : "", "category": "", "content": ""}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/posts/"+strconv.Itoa(post.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestGetPostSuccess(t *testing.T) {
	db := setupTestDB()
	truncatePost(db)

	tx, _ := db.Begin()
	postRepository := repository.NewPostRepository()
	post := postRepository.Save(context.Background(), tx, domain.Post{
		Title: "Review Laptop Legion Slim 7",
		Category: "Laptop",
		Content: "This is review about Legion Slim 7",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/posts/"+strconv.Itoa(post.Id), nil)
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, post.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, post.Title, responseBody["data"].(map[string]interface{})["title"])
	assert.Equal(t, post.Category, responseBody["data"].(map[string]interface{})["category"])
	assert.Equal(t, post.Content, responseBody["data"].(map[string]interface{})["content"])
}

func TestGetPostFailed(t *testing.T) {
	db := setupTestDB()
	truncatePost(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/posts/404", nil)
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

func TestDeletePostSuccess(t *testing.T) {
	db := setupTestDB()
	truncatePost(db)

	tx, _ := db.Begin()
	postRepository := repository.NewPostRepository()
	post := postRepository.Save(context.Background(), tx, domain.Post{
		Title: "Review Laptop Legion Slim 7",
		Category: "Laptop",
		Content: "This is review about Legion Slim 7",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/posts/"+strconv.Itoa(post.Id), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
}

func TestDeletePostFailed(t *testing.T) {
	db := setupTestDB()
	truncatePost(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/posts/404", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

func TestListPostsSuccess(t *testing.T) {
	db := setupTestDB()
	truncatePost(db)

	tx, _ := db.Begin()
	postRepository := repository.NewPostRepository()
	post1 := postRepository.Save(context.Background(), tx, domain.Post{
		Title: "Review Laptop Legion Slim 7",
		Category: "Laptop",
		Content: "This is review about Legion Slim 7",
	})
	post2 := postRepository.Save(context.Background(), tx, domain.Post{
		Title: "Review Laptop Acer Nitro 5",
		Category: "Laptop",
		Content: "This is review about Acer Nitro 5",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/posts", nil)
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])

	fmt.Println(responseBody)

	var posts = responseBody["data"].([]interface{})

	postResponse1 := posts[0].(map[string]interface{})
	postResponse2 := posts[1].(map[string]interface{})

	assert.Equal(t, post1.Id, int(postResponse1["id"].(float64)))
	assert.Equal(t, post1.Title, postResponse1["title"])
	assert.Equal(t, post1.Category, postResponse1["category"])
	assert.Equal(t, post1.Content, postResponse1["content"])

	assert.Equal(t, post2.Id, int(postResponse2["id"].(float64)))
	assert.Equal(t, post2.Title, postResponse2["title"])
	assert.Equal(t, post2.Category, postResponse2["category"])
	assert.Equal(t, post2.Content, postResponse2["content"])
}

func TestUnauthorized(t *testing.T) {
	db := setupTestDB()
	truncatePost(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/posts", nil)
	request.Header.Add("X-API-Key", "SALAH")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 401, int(responseBody["code"].(float64)))
	assert.Equal(t, "UNAUTHORIZED", responseBody["status"])
}
