package controller

import (
	"my-blog/helper"
	"my-blog/model/web"
	"my-blog/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type PostControllerImpl struct {
	PostService service.PostService
}

func NewPostController(postService service.PostService) PostController {
	return &PostControllerImpl{
		PostService: postService,
	}
}

func (controller *PostControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	postCreateRequest := web.PostCreateRequest{}
	helper.ReadFromRequestBody(request, &postCreateRequest)

	postResponse := controller.PostService.Create(request.Context(), postCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   postResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PostControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	postUpdateRequest := web.PostUpdateRequest{}
	helper.ReadFromRequestBody(request, &postUpdateRequest)

	postId := params.ByName("postId")
	id, err := strconv.Atoi(postId)
	helper.PanicIfError(err)

	postUpdateRequest.Id = id

	postResponse := controller.PostService.Update(request.Context(), postUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   postResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PostControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	postId := params.ByName("postId")
	id, err := strconv.Atoi(postId)
	helper.PanicIfError(err)

	controller.PostService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PostControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	postId := params.ByName("postId")
	id, err := strconv.Atoi(postId)
	helper.PanicIfError(err)

	postResponse := controller.PostService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   postResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PostControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	postResponses := controller.PostService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   postResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
