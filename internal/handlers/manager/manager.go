package manager

import (
	"github.com/akmuhammetakmyradov/test/internal/posts"
	"github.com/akmuhammetakmyradov/test/pkg/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/redis/go-redis/v9"
)

const (
	baseURL = "/api/test/v1"

	healthcheckerURL = "/healthchecker"
	postsURL         = "/posts"
)

func Manager(db *pgxpool.Pool, cache *redis.Client, cfg *config.Configs) *fiber.App {
	app := fiber.New(fiber.Config{
		BodyLimit: 100 * 1024 * 1024,
	})

	app.Use(logger.New())
	app.Use(cors.New())
	router := app.Group(baseURL)

	// test for server is working or not
	router.Get(healthcheckerURL, func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "The Service is running",
		})
	})

	// posts
	postRouterManager := router.Group(postsURL)
	postRouterRepository := posts.NewRepository(db, cache)
	postRouterHandler := posts.NewHandler(postRouterRepository, cfg)
	postRouterHandler.Register(postRouterManager)

	return app
}
