package main

import (
	h "auth-service-medods/internal/handler"
	psql "auth-service-medods/internal/storage"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	URL := os.Getenv("DATABASE_URL")
	if URL == "" {
		panic("DATABASE_URL not found")
	}

	PORT := os.Getenv("PORT")
	if PORT == "" {
		panic("PORT not found")
	}

	db, err := psql.Open(URL)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	e := echo.New()

	e.POST("/token", h.GetAccessRefresh)
	e.POST("/refresh", h.PostRefresh)

	e.Start(PORT)
}
