package handler

import (
	"backend-test-template/app/common"
	"backend-test-template/app/model"
	"gorm.io/gorm"
	"net/http"
)

type MigrationHandler struct {
	db *gorm.DB
}

func NewMigrationHandler(db *gorm.DB) *MigrationHandler {
	return &MigrationHandler{db: db}
}

func (h *MigrationHandler) MigrateDatabase(writer http.ResponseWriter, request *http.Request) {
	var err error
	tx := h.db.Begin()

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	_ = tx.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	models := []interface{}{
		&model.User{},
		&model.Post{},
		&model.PostResponse{},
		&model.UserToken{},
	}
	for _, m := range models {
		err = tx.AutoMigrate(m)
		if err != nil {
			common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "While migrating database got error : "+err.Error(), nil))
		}
	}
	common.ResponseWithJson(writer, http.StatusBadRequest, common.ResponseStandard(http.StatusBadRequest, "Migrate success", nil))
}
