package handler

import (
	"backend-test-template/app/common"
	"backend-test-template/app/model"
	"backend-test-template/app/repo"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

type PostResponseHandlers struct {
	obRepo   *repo.PostResponseRepo
	userRepo *repo.UserRepo
	postRepo *repo.PostRepo
}

func NewPostResponseHandlers(repo *repo.PostResponseRepo, userRepo *repo.UserRepo, postRepo *repo.PostRepo) *PostResponseHandlers {
	return &PostResponseHandlers{obRepo: repo, userRepo: userRepo, postRepo: postRepo}
}

func (h *PostResponseHandlers) GetPostResponseById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	PostResponseID, err := uuid.Parse(params["id"])
	if err != nil {
		common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "While parsing id got error : "+err.Error(), nil))
		return
	}

	PostResponse, err := h.obRepo.Get(request.Context(), PostResponseID)
	if err != nil {
		common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "While getting id got error : "+err.Error(), nil))
		return
	}
	common.ResponseWithJson(writer, http.StatusNotFound, PostResponse)
}

func (h *PostResponseHandlers) GetAllResponseOfAPost(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	PostID, err := uuid.Parse(params["id"])
	if err != nil {
		common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "While parsing id got error : "+err.Error(), nil))
		return
	}
	PostResponses, err := h.obRepo.GetAll(request.Context(), PostID)
	if err != nil {
		common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "While getting id got error : "+err.Error(), nil))
		return
	}
	common.ResponseWithJson(writer, http.StatusNotFound, PostResponses)
}

func (h *PostResponseHandlers) Create(writer http.ResponseWriter, request *http.Request) {
	var PostResponseCreateRequest model.PostResponseCreate
	if err := json.NewDecoder(request.Body).Decode(&PostResponseCreateRequest); err != nil {
		common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "While parsing request got error : "+err.Error(), nil))
		return
	}
	if PostResponseCreateRequest.UserID == nil || PostResponseCreateRequest.Message == nil || PostResponseCreateRequest.PostID == nil {
		common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "While parsing request got error : user_id, post_id and message must not be empty", nil))
		return
	}
	PostResponse := model.PostResponse{
		Message: *PostResponseCreateRequest.Message,
		UserID:  *PostResponseCreateRequest.UserID,
		PostID:  *PostResponseCreateRequest.PostID,
	}

	if err := h.obRepo.Create(request.Context(), &PostResponse); err != nil {
		common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "While creating PostResponse got error : "+err.Error(), nil))
		return
	}
	common.ResponseWithJson(writer, http.StatusNotFound, common.ResponseStandard(http.StatusOK, "Create success", PostResponse))
}

func (h *PostResponseHandlers) UpdatePostResponseByID(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	var PostResponseUpdate model.PostResponseUpdate
	if err := json.NewDecoder(request.Body).Decode(&PostResponseUpdate); err != nil {
		common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "While parsing request got error : "+err.Error(), nil))
		return
	}

	params := mux.Vars(request)
	PostResponseID, err := uuid.Parse(params["id"])
	if err != nil {
		common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "While parsing request id got error : "+err.Error(), nil))
		return
	}
	PostResponse, err := h.obRepo.Get(ctx, PostResponseID)
	if err != nil {
		common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "While validating PostResponse got error : "+err.Error(), nil))
		return
	}
	common.Sync(PostResponseUpdate, PostResponse)
	if err = h.obRepo.Update(ctx, PostResponse); err != nil {
		common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "While updating PostResponse got error : "+err.Error(), nil))
		return
	}
	common.ResponseWithJson(writer, http.StatusNotFound, common.ResponseStandard(http.StatusOK, "Update success", nil))
}

func (h *PostResponseHandlers) DeletePostResponseByID(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	params := mux.Vars(request)
	PostResponseID, err := uuid.Parse(params["id"])
	if err != nil {
		common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "While parsing id error : "+err.Error(), nil))
		return
	}
	PostResponse, err := h.obRepo.Get(ctx, PostResponseID)
	if err != nil {
		common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "While validating PostResponse got error : "+err.Error(), nil))
		return
	}
	if err = h.obRepo.DeletePostResponse(request.Context(), PostResponse); err != nil {
		common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "While deleting PostResponse got error : "+err.Error(), nil))
		return
	}
	common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "Delete success", nil))
}
