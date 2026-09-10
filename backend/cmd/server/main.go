package main

import (
	"log"
	"net/http"

	"github.com/Dew-F/v-chat/internal/auth"
	"github.com/Dew-F/v-chat/internal/config"
	"github.com/Dew-F/v-chat/internal/database"
	"github.com/Dew-F/v-chat/internal/id"
	"github.com/Dew-F/v-chat/internal/middleware"
	"github.com/Dew-F/v-chat/internal/router"
	"github.com/Dew-F/v-chat/internal/user"
)

func main() {

	cfg := config.Load()

	if err := id.Init(cfg.SnowflakeWorkerID); err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect(cfg.PostgresURL)
	if err != nil {
		log.Fatal(err)
	}

	userRepository := user.NewRepository(db)

	authService := auth.NewService(userRepository, auth.NewJWTManager(cfg.JWTSecret))

	authHandler := auth.NewHandler(authService)

	r := router.New(authHandler)

	handler := middleware.Middleware(r)

	log.Println("server started :8080")

	log.Fatal(
		http.ListenAndServe(":8080", handler),
	)
}
