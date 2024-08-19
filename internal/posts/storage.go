package posts

import (
	"context"

	postscache "github.com/akmuhammetakmyradov/test/internal/posts/cache"
	postsdb "github.com/akmuhammetakmyradov/test/internal/posts/db"
	"github.com/akmuhammetakmyradov/test/internal/posts/models"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/redis/go-redis/v9"
)

type Db interface {
	GetUser(ctx context.Context, login string) (models.User, error)
	CreateUser(ctx context.Context, user models.UserDTO) (models.ID, error)
	CreatePost(ctx context.Context, user models.PostDTO) (models.ID, error)
	DeletePost(ctx context.Context, postID int) error
	GetPosts(ctx context.Context, limit, page int) ([]models.Post, error)
	GetPost(ctx context.Context, postID int) (models.Post, error)
}

type Cache interface {
}

type Repository struct {
	Db
	Cache
}

func NewRepository(db *pgxpool.Pool, client *redis.Client) *Repository {
	return &Repository{
		Db:    postsdb.NewDbRepo(db),
		Cache: postscache.NewRedisRepo(client),
	}
}
