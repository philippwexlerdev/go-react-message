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

type PostHandlers struct {
	obRepo   *repo.PostRepo
	userRepo *repo.UserRepo
}

func NewPostHandlers(repo *repo.PostRepo, userRepo *repo.UserRepo) *PostHandlers {
	return &PostHandlers{obRepo: repo, userRepo: userRepo}
}

func (h *PostHandlers) GetPostById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	PostID, err := uuid.Parse(params["id"])
	if err != nil {
		common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "While parsing id got error : "+err.Error(), nil))
		return
	}

	Post, err := h.obRepo.Get(request.Context(), PostID)
	if err != nil {
		common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "While getting id got error : "+err.Error(), nil))
		return
	}
	common.ResponseWithJson(writer, http.StatusNotFound, Post)
}

func (h *PostHandlers) GetAll(writer http.ResponseWriter, request *http.Request) {
	Posts, err := h.obRepo.GetAll(request.Context())
	if err != nil {
		common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "While getting id got error : "+err.Error(), nil))
		return
	}
	common.ResponseWithJson(writer, http.StatusNotFound, Posts)
}

func (h *PostHandlers) Create(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	var PostCreateRequest model.PostCreate
	if err := json.NewDecoder(request.Body).Decode(&PostCreateRequest); err != nil {
		common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "While parsing request got error : "+err.Error(), nil))
		return
	}
	if PostCreateRequest.UserID == nil || PostCreateRequest.Message == nil {
		common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "While parsing request got error : user_id and message must not be empty", nil))
		return
	}

	if _, err := h.userRepo.GetUser(ctx, *PostCreateRequest.UserID); err != nil {
		common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "While validating user got error : "+err.Error(), nil))
		return
	}
	Post := model.Post{
		Message: *PostCreateRequest.Message,
		UserID:  *PostCreateRequest.UserID,
	}

	if err := h.obRepo.Create(request.Context(), &Post); err != nil {
		common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "While creating Post got error : "+err.Error(), nil))
		return
	}
	common.ResponseWithJson(writer, http.StatusNotFound, common.ResponseStandard(http.StatusOK, "Create success", Post))
}

func (h *PostHandlers) UpdatePostByID(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	var PostUpdate model.PostUpdate
	if err := json.NewDecoder(request.Body).Decode(&PostUpdate); err != nil {
		common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "While parsing request got error : "+err.Error(), nil))
		return
	}

	params := mux.Vars(request)
	PostID, err := uuid.Parse(params["id"])
	if err != nil {
		common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "While parsing request id got error : "+err.Error(), nil))
		return
	}
	Post, err := h.obRepo.Get(ctx, PostID)
	if err != nil {
		common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "While validating Post got error : "+err.Error(), nil))
		return
	}
	common.Sync(PostUpdate, Post)
	if err = h.obRepo.Update(ctx, Post); err != nil {
		common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "While updating Post got error : "+err.Error(), nil))
		return
	}
	common.ResponseWithJson(writer, http.StatusNotFound, common.ResponseStandard(http.StatusOK, "Update success", nil))
}

func (h *PostHandlers) DeletePostByID(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	params := mux.Vars(request)
	PostID, err := uuid.Parse(params["id"])
	if err != nil {
		common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "While parsing id error : "+err.Error(), nil))
		return
	}
	Post, err := h.obRepo.Get(ctx, PostID)
	if err != nil {
		common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "While validating Post got error : "+err.Error(), nil))
		return
	}
	if err = h.obRepo.DeletePost(request.Context(), Post); err != nil {
		common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "While deleting Post got error : "+err.Error(), nil))
		return
	}
	common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "Delete success", nil))
}
