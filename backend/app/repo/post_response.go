package repo

import (
	"backend-test-template/app/common"
	"backend-test-template/app/model"
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostResponseRepo struct {
	db *gorm.DB
}

func (r *PostResponseRepo) DBWithTimeout(ctx context.Context) (*gorm.DB, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(ctx, common.GeneralQueryTimeout)
	return r.db.WithContext(ctx), cancel
}

func NewPostResponseRepo(db *gorm.DB) *PostResponseRepo {
	return &PostResponseRepo{db: db}
}

func (r *PostResponseRepo) Create(ctx context.Context, ob *model.PostResponse) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()

	return tx.Create(ob).Error
}
func (r *PostResponseRepo) Get(ctx context.Context, id uuid.UUID) (*model.PostResponse, error) {
	o := &model.PostResponse{}

	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()

	err := tx.First(&o, id).Error
	return o, err
}

func (r *PostResponseRepo) GetAll(ctx context.Context, postID uuid.UUID) ([]model.PostResponse, error) {
	o := []model.PostResponse{}

	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	err := tx.Where("post_id = ?", postID).Find(&o).Error
	return o, err
}

func (r *PostResponseRepo) Update(ctx context.Context, update *model.PostResponse) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	return tx.WithContext(ctx).Where("id = ?", update.ID).Save(&update).Error
}

func (r *PostResponseRepo) DeletePostResponse(ctx context.Context, d *model.PostResponse) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	return tx.WithContext(ctx).Unscoped().Delete(&d).Error
}
