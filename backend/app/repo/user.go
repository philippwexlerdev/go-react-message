package repo

import (
	"backend-test-template/app/common"
	"backend-test-template/app/model"
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func (r *UserRepo) DBWithTimeout(ctx context.Context) (*gorm.DB, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(ctx, common.GeneralQueryTimeout)
	return r.db.WithContext(ctx), cancel
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(ctx context.Context, ob *model.User) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()

	return tx.Create(ob).Error
}
func (r *UserRepo) GetUser(ctx context.Context, id uuid.UUID) (*model.User, error) {
	o := &model.User{}

	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()

	err := tx.First(&o, id).Error
	return o, err
}

func (r *UserRepo) ValidateLogin(ctx context.Context, username string, password string) (*model.User, error) {
	o := &model.User{}

	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()

	err := tx.Where("username = ? and password = ?", username, password).First(&o).Error
	return o, err
}

func (r *UserRepo) GetOneFlexible(ctx context.Context, field string, value interface{}) (model.User, error) {
	o := model.User{}

	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()

	err := tx.Where(field+" = ? ", value).First(&o).Error
	return o, err
}

func (r *UserRepo) Update(ctx context.Context, update *model.User) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	return tx.Where("id = ?", update.ID).Save(&update).Error
}

func (r *UserRepo) DeleteUser(ctx context.Context, d *model.User) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	return tx.WithContext(ctx).Unscoped().Delete(&d).Error
}
