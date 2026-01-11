package category

import "context"

type CategoryService interface {
	FindAllCategories(ctx context.Context, page, limit, offset int) ([]CategoryResponseDTO, int64, error)
	FindCategoryByID(ctx context.Context, id string) (*CategoryResponseDTO, error)
	CreateCategory(ctx context.Context, dto *CategoryRequestDTO) (*CategoryResponseDTO, error)
	UpdateCategoryByID(ctx context.Context, id string, dto *CategoryRequestDTO) (*CategoryResponseDTO, error)
	DeleteCategoryByID(ctx context.Context, id string) error
}
