package handler

import (
	"backend-test-template/app/common"
	"backend-test-template/app/repo"
	"net/http"
)

type Middleware struct {
	userTokenRepo *repo.UserTokenRepo
}

func NewMiddleware(userTokenRepo *repo.UserTokenRepo) *Middleware {
	return &Middleware{userTokenRepo: userTokenRepo}
}

func (h *Middleware) Authen(f func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			common.ResponseWithJson(w, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "Invalid Authorization", nil))
			return
		}

		if _, err := h.userTokenRepo.ValidateToken(r.Context(), token); err != nil {
			common.ResponseWithJson(w, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "Invalid Authorization", nil))
			return
		}

		f(w, r)
	}
}
