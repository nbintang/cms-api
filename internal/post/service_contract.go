package post

import "context"

type PostService interface {
	FindAllPosts(ctx context.Context, page, limit, offset int) ([]PaginatedPostResponseDTO, int64, error)
	FindPostByID(ctx context.Context, id string) (*PostResponseDTO, error)
	CreatePost(ctx context.Context, dto PostRequestDTO, userID string) (*PostResponseDTO, error)
	UpdatePostByID(ctx context.Context, id string, dto PostRequestDTO, userID string) (*PostResponseDTO, error)
	DeletePostByID(ctx context.Context, id string) error
}
