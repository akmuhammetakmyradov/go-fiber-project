package posts

import (
	"context"
	"time"

	postscache "github.com/akmuhammetakmyradov/test/internal/posts/cache"
	postsdb "github.com/akmuhammetakmyradov/test/internal/posts/db"
	"github.com/akmuhammetakmyradov/test/internal/posts/models"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/redis/go-redis/v9"
)

type Db interface {
	GetUser(ctx context.Context, login string) (models.User, error)
	CreateUser(ctx context.Context, user models.UserDTO) (models.ID, error)
	CreatePost(ctx context.Context, user models.PostDTO) (models.Post, error)
	DeletePost(ctx context.Context, postID int) error
	GetPosts(ctx context.Context, limit, page int) ([]models.Post, error)
	GetPost(ctx context.Context, postID int) (models.Post, error)
}

type Cache interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
	PaginationAdd(ctx context.Context, key string, score float64, data interface{}) error
	PaginationGet(ctx context.Context, key string, start, end int) ([]string, error)
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
