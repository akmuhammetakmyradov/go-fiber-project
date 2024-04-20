package posts

import "context"

type Repository interface {
	GetUser(ctx context.Context, login string) (User, error)
	CreateUser(ctx context.Context, user UserDTO) (ID, error)
	CreatePost(ctx context.Context, user PostDTO) (ID, error)
	DeletePost(ctx context.Context, postID int) error
	GetPosts(ctx context.Context, limit, page int) ([]Post, error)
	GetPost(ctx context.Context, postID int) (Post, error)
}
