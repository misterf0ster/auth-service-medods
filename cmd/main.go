package main

import (
	"auth-service-medods/internal/handler"
	psql "auth-service-medods/internal/storage"
	"auth-service-medods/pkg/config"
	"auth-service-medods/pkg/logger"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	logger.LoggerInit()

	logger.LogInfo("Starting the service...")

	config.LoadEnv()
	cfg := config.Config()

	url := cfg.DBaseURL()
	conn, err := psql.Open(url)
	if err != nil {
		logger.Log.Fatalf("Unable to connect to db: %v\n", err)
	}
	defer conn.Close()

	h := handler.CreateToken(conn)
	g := gin.Default()

	g.GET("/token", h.GetAccessRefresh)
	g.POST("/refresh", h.PostRefresh)

	port := os.Getenv("PORT")
	if port == "" {
		logger.LogInfo("Port not found")
	}

	logger.Log.Println("Starting server on port", port)
	if err := g.Run(":" + port); err != nil {
		logger.Log.Fatal("Server startup error: " + err.Error())
	}
}
