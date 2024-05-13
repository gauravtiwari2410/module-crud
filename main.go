package main

import (
	"fmt"
	"log"
	"module-crud/config"
	"module-crud/novel/controller"
	"module-crud/novel/domain/model"
	"module-crud/novel/repo"
	"module-crud/novel/usecase"

	"github.com/go-redis/redis"
	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("hello world")

	cfg, err := config.LoadConfig(".") // Rename to avoid shadowing
	if err != nil {
		log.Fatalf("cannot load environment variables: %v", err)
	}

	db, err := database.ConnectionDB(&cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&model.Novel{}); err != nil {
		log.Fatalf("Failed to auto migrate models: %v", err)
	}

	rdb := database.ConnectionRedisDB(&cfg)
	startServer(db, rdb)
}

func startServer(db *gorm.DB, rdb *redis.Client) {
	app := fiber.New()
	novelRepo := repo.NewNovelRepo(db, rdb)
	novelUseCase := usecase.NewNovelUseCase(novelRepo)
	novelController := controller.NewNovelController(novelUseCase)
	router.SetupRoutes(app, novelController) // Assume SetupRoutes is a correct method

	if err := app.Listen(":3400"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
