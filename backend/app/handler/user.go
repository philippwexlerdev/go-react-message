package handler

import (
	"backend-test-template/app/common"
	"backend-test-template/app/model"
	"backend-test-template/app/repo"
	"encoding/base64"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type UserHandlers struct {
	obRepo        *repo.UserRepo
	userTokenRepo *repo.UserTokenRepo
}

func NewUserHandlers(repo *repo.UserRepo, userTokenRepo *repo.UserTokenRepo) *UserHandlers {
	return &UserHandlers{obRepo: repo, userTokenRepo: userTokenRepo}
}

func (h *UserHandlers) GetUserById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	userID, err := uuid.Parse(params["id"])
	if err != nil {
		common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "While parsing id got error : "+err.Error(), nil))
		return
	}
	user, err := h.obRepo.GetUser(request.Context(), userID)
	if err != nil {
		common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "While getting id got error : "+err.Error(), nil))
		return
	}
	common.ResponseWithJson(writer, http.StatusNotFound, user)
}

func (h *UserHandlers) Login(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	var userCreateRequest model.UserCreateOrLogin
	if err := json.NewDecoder(request.Body).Decode(&userCreateRequest); err != nil {
		common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "While parsing request got error : "+err.Error(), nil))
		return
	}
	if userCreateRequest.Username == "" || userCreateRequest.Password == "" {
		common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "While parsing request got error : username and password must not be empty", nil))
		return
	}

	data, err := base64.StdEncoding.DecodeString(userCreateRequest.Password)
	if err != nil {
		common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "Password must be base64 encode, try this: https://www.base64encode.org/", nil))
		return
	}
	user, err := h.obRepo.ValidateLogin(ctx, userCreateRequest.Username, string(data))

	if err != nil {
		common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "Password must be base64 encode, try this: https://www.base64encode.org/", nil))
		return
	}
	newToken := common.RandStringGenerator(30)
	newExpireTime := time.Now().Local().Add(time.Minute * 10)
	userToken, err := h.userTokenRepo.GetOneFlexible(ctx, "user_id", user.ID)
	if err != nil {
		newUserToken := &model.UserToken{
			Token:    newToken,
			UserID:   user.ID,
			ExpireAt: newExpireTime,
		}
		if err := h.userTokenRepo.Create(ctx, newUserToken); err != nil {
			common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "While creating a new token got err : "+err.Error(), nil))
			return
		}
		common.ResponseWithJson(writer, http.StatusNotFound, common.ResponseStandard(http.StatusOK, "Login success", map[string]interface{}{"token": newToken, "user_id": user.ID}))
		return
	}
	userToken.ExpireAt = newExpireTime
	userToken.Token = newToken
	if err = h.userTokenRepo.Update(ctx, &userToken); err != nil {
		common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "While creating a new token got err : "+err.Error(), nil))
		return
	}
	common.ResponseWithJson(writer, http.StatusNotFound, common.ResponseStandard(http.StatusOK, "Login success", map[string]interface{}{"token": newToken, "user_id": user.ID}))
}

func (h *UserHandlers) Create(writer http.ResponseWriter, request *http.Request) {
	var userCreateRequest model.UserCreateOrLogin
	if err := json.NewDecoder(request.Body).Decode(&userCreateRequest); err != nil {
		common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "While parsing request got error : "+err.Error(), nil))
		return
	}
	if userCreateRequest.Username == "" || userCreateRequest.Password == "" {
		common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "While parsing request got error : username and password must not be empty", nil))
		return
	}
	data, err := base64.StdEncoding.DecodeString(userCreateRequest.Password)
	if err != nil {
		common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "Password must be base64 encode, try this: https://www.base64encode.org/", nil))
		return
	}
	user := model.User{
		Username: userCreateRequest.Username,
		Password: string(data),
	}

	if err := h.obRepo.Create(request.Context(), &user); err != nil {
		common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "While creating user got error : "+err.Error(), nil))
		return
	}
	common.ResponseWithJson(writer, http.StatusNotFound, common.ResponseStandard(http.StatusOK, "Create success", nil))
}

func (h *UserHandlers) UpdateUserByID(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	var userUpdate model.UseUpdate
	if err := json.NewDecoder(request.Body).Decode(&userUpdate); err != nil {
		common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "While parsing request got error : "+err.Error(), nil))
		return
	}

	params := mux.Vars(request)
	userID, err := uuid.Parse(params["id"])
	if err != nil {
		common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "While parsing request id got error : "+err.Error(), nil))
		return
	}
	user, err := h.obRepo.GetUser(ctx, userID)
	if err != nil {
		common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "While validating user got error : "+err.Error(), nil))
		return
	}
	common.Sync(userUpdate, user)
	if err = h.obRepo.Update(ctx, user); err != nil {
		common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "While updating user got error : "+err.Error(), nil))
		return
	}
	common.ResponseWithJson(writer, http.StatusNotFound, common.ResponseStandard(http.StatusOK, "Update success", nil))
}

func (h *UserHandlers) DeleteUserByID(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	params := mux.Vars(request)
	userID, err := uuid.Parse(params["id"])
	if err != nil {
		common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "While parsing id error : "+err.Error(), nil))
		return
	}
	user, err := h.obRepo.GetUser(ctx, userID)
	if err != nil {
		common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "While validating user got error : "+err.Error(), nil))
		return
	}
	if err = h.obRepo.DeleteUser(request.Context(), user); err != nil {
		common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "While deleting user got error : "+err.Error(), nil))
		return
	}
	common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "Delete success", nil))
}
