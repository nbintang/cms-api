package category

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type categoryRepositoryImpl struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepositoryImpl{db}
}

func (r *categoryRepositoryImpl) ExistsByID(ctx context.Context, id string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&Category{}).
		Where("id = ?", id).
		Count(&count).Error
	return count > 0, err
}

func (r *categoryRepositoryImpl) FindAll(ctx context.Context, limit, offset int) ([]Category, int64, error) {
	var categories []Category
	var total int64
	var category Category
	db := r.db.WithContext(ctx).Model(&category)
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := db.Scopes(
		Paginate(limit, offset),
		SelectedFields,
	).Find(&categories).Error; err != nil {
		return nil, 0, err
	}
	return categories, total, nil
}

func (r *categoryRepositoryImpl) FindByID(ctx context.Context, id string) (*Category, error) {
	var category Category
	if err := r.db.WithContext(ctx).
		Scopes(WhereID(id), SelectedFields).
		Take(&category).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepositoryImpl) Create(ctx context.Context, category *Category) (uuid.UUID, error) {
	if err := r.db.WithContext(ctx).Create(&category).Error; err != nil {
		return uuid.Nil, nil
	}
	return category.ID, nil
}
func (r *categoryRepositoryImpl) Update(ctx context.Context, id string, category *Category) (uuid.UUID, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return uuid.Nil, err
	}

	tx := r.db.WithContext(ctx).
		Model(&Category{}).
		Scopes(WhereID(id)).
		Updates(category)

	if tx.Error != nil {
		return uuid.Nil, tx.Error
	}

	if tx.RowsAffected == 0 {
		return uuid.Nil, gorm.ErrRecordNotFound
	}

	return uid, nil
}
func (r *categoryRepositoryImpl) Delete(ctx context.Context, id string) error {
	if err := r.db.WithContext(ctx).
		Scopes(WhereID(id)).
		Delete(&Category{}).
		Error; err != nil {
		return err
	}
	return nil
}
