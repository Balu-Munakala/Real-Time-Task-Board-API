package main

import (
	lgr "log"

	"github.com/Balu-Munakala/Real-Time-Task-Board-API/internal/config"
	"github.com/Balu-Munakala/Real-Time-Task-Board-API/internal/platform/database"
	"github.com/Balu-Munakala/Real-Time-Task-Board-API/internal/platform/logger"
	"github.com/Balu-Munakala/Real-Time-Task-Board-API/internal/server"
	"github.com/Balu-Munakala/Real-Time-Task-Board-API/internal/user"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		lgr.Fatal("error occured while loading config")
	}
	lgr.Print("Configurations Loaded Successfully")

	if cfg.AppEnv == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	log := logger.NewLogger(cfg.AppEnv)
	defer log.Sync()
	if log == nil {
		lgr.Fatal("Couldn't Load Logger")
	}

	log.Debug("Import Logger Successfully")

	db, err := database.Connect(cfg.DbURL)
	if err != nil {
		log.Fatal("Error Connecting with DB : " + err.Error())
	}
	defer db.Close()
	log.Debug("DB Successfully Connected")

	userRepo := user.NewUserRepository(db)
	userService := user.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService)

	router := server.NewRouteHandler(log, userHandler)

	// s := server.New(cfg, log, router)
	s := server.New(cfg, router)

	s.Run()
}
