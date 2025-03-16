package main

import (
	"context"
	"errors"
	"github.com/dnevsky/restaurant-back/internal/pkg/auth"
	"github.com/dnevsky/restaurant-back/internal/pkg/config"
	"github.com/dnevsky/restaurant-back/internal/pkg/logger"
	"github.com/dnevsky/restaurant-back/internal/repository"
	"github.com/dnevsky/restaurant-back/internal/repository/postgres"
	"github.com/dnevsky/restaurant-back/internal/service"
	"github.com/dnevsky/restaurant-back/internal/transport/rest"
	"github.com/dnevsky/restaurant-back/internal/transport/rest/helpers"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger.Init()
	env := os.Getenv("ENV")

	if env == "dev" {
		if err := godotenv.Load(); err != nil {
			log.Fatal("error loading .env file")
		}
	}

	config.InitConfig()

	db, err := postgres.NewPostgres()
	if err != nil {
		log.Fatalf("error connecting to database: %s", err)
	}
	defer postgres.CloseDB(db)

	store := initStore(db)

	tokenManager, err := auth.NewManager(config.Config.JwtSecret)
	if err != nil {
		log.Fatal(err)
		return
	}

	services, err := service.NewService(service.Deps{
		Repository:   store,
		TokenManager: tokenManager,
	})

	helpersManager := helpers.NewManager(
		store.UserRepo,
	)

	httpServerInstance := new(rest.Server)

	handlers := rest.NewHandler(services, helpersManager)

	go func() {
		if err := httpServerInstance.RunHttp(handlers.InitRoutes()); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("error http server: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	log.Println("Started.")
	<-quit

	if err := httpServerInstance.Shutdown(context.Background()); err != nil {
		log.Fatalf("error occurated on shuting down server: %s", err.Error())
	}

	postgres.CloseDB(db)

	log.Println("Shutdown...")
}

func initStore(db *gorm.DB) *repository.Repository {
	repo := repository.New(db)
	return repo
}
