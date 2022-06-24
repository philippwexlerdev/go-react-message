package repo

import (
	"backend-test-template/app/common"
	"backend-test-template/app/model"
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostRepo struct {
	db *gorm.DB
}

func (r *PostRepo) DBWithTimeout(ctx context.Context) (*gorm.DB, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(ctx, common.GeneralQueryTimeout)
	return r.db.WithContext(ctx), cancel
}

func NewPostRepo(db *gorm.DB) *PostRepo {
	return &PostRepo{db: db}
}

func (r *PostRepo) Create(ctx context.Context, ob *model.Post) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()

	return tx.Create(ob).Error
}
func (r *PostRepo) Get(ctx context.Context, id uuid.UUID) (*model.Post, error) {
	o := &model.Post{}

	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()

	err := tx.First(&o, id).Error
	return o, err
}

func (r *PostRepo) GetAll(ctx context.Context) ([]model.Post, error) {
	o := []model.Post{}

	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	err := tx.Find(&o).Error
	return o, err
}

func (r *PostRepo) Update(ctx context.Context, update *model.Post) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	return tx.WithContext(ctx).Where("id = ?", update.ID).Save(&update).Error
}

func (r *PostRepo) DeletePost(ctx context.Context, d *model.Post) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	return tx.WithContext(ctx).Unscoped().Delete(&d).Error
}
