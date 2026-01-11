package category

import (
	"context"

	"github.com/google/uuid"
)

type CategoryRepository interface {
	ExistsByID(ctx context.Context, id string) (bool, error)
	FindAll(ctx context.Context, limit, offset int) ([]Category, int64, error)
	FindByID(ctx context.Context, id string) (*Category, error)
	Create(ctx context.Context, category *Category) (uuid.UUID, error)
	Update(ctx context.Context,id string, category *Category) (uuid.UUID, error) 
	Delete(ctx context.Context, id string) error 
}
