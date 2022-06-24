package repo

import (
	"backend-test-template/app/common"
	"backend-test-template/app/model"
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserTokenRepo struct {
	db *gorm.DB
}

func (r *UserTokenRepo) DBWithTimeout(ctx context.Context) (*gorm.DB, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(ctx, common.GeneralQueryTimeout)
	return r.db.WithContext(ctx), cancel
}

func NewUserTokenRepo(db *gorm.DB) *UserTokenRepo {
	return &UserTokenRepo{db: db}
}

func (r *UserTokenRepo) Create(ctx context.Context, ob *model.UserToken) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()

	return tx.Create(ob).Error
}
func (r *UserTokenRepo) GetOne(ctx context.Context, id uuid.UUID) (*model.UserToken, error) {
	o := &model.UserToken{}

	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()

	err := tx.First(&o, id).Error
	return o, err
}

func (r *UserTokenRepo) GetOneFlexible(ctx context.Context, field string, value interface{}) (model.UserToken, error) {
	o := model.UserToken{}

	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()

	err := tx.Where(field+" = ? ", value).First(&o).Error
	return o, err
}

func (r *UserTokenRepo) ValidateToken(ctx context.Context, token string) (model.UserToken, error) {
	o := model.UserToken{}

	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()

	err := tx.Where("token = ? AND expire_at >= current_timestamp", token).First(&o).Error
	return o, err
}

func (r *UserTokenRepo) Update(ctx context.Context, update *model.UserToken) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	return tx.Where("id = ?", update.ID).Save(&update).Error
}

func (r *UserTokenRepo) Delete(ctx context.Context, d *model.UserToken) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	return tx.WithContext(ctx).Unscoped().Delete(&d).Error
}
