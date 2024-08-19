package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/akmuhammetakmyradov/test/internal/handlers/manager"
	"github.com/akmuhammetakmyradov/test/pkg/config"
	"github.com/akmuhammetakmyradov/test/pkg/postgresql"
	"github.com/akmuhammetakmyradov/test/pkg/redis"
)

func InitApp(cfg *config.Configs) error {
	db, err := postgresql.NewPostgres(cfg)
	if err != nil {
		return err
	}

	defer func() {
		db.Close()
	}()

	redisClient, err := redis.NewRedisClient(cfg)
	if err != nil {
		return err
	}

	app := manager.Manager(db, redisClient, cfg)

	go func() {
		if err := app.Listen(fmt.Sprintf(":%s", cfg.Listen.Port)); err != nil {
			log.Fatalln(err)
		}
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)
	<-signals
	log.Println("Shutdown Server ...")

	if err := app.ShutdownWithTimeout(5 * time.Second); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	return nil
}
