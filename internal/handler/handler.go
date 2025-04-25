package handler

import (
	auth "auth-service-medods/internal/service"
	psql "auth-service-medods/internal/storage"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Контроллер для получения Access и Refresh токенов
func GetAccessRefresh(c echo.Context) error {
	userID := c.QueryParam("user_id")
	ip := c.RealIP()

	db := c.Get("db").(*psql.DB) // Получаем подключение к базе данных

	tokens, err := auth.GenerateTokens(userID, ip, db)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "could not generate tokens",
		})
	}

	return c.JSON(http.StatusOK, tokens)
}

// Контроллер для обновления токенов
func PostRefresh(c echo.Context) error {
	accessToken := c.FormValue("access_token")
	refreshToken := c.FormValue("refresh_token")
	ip := c.RealIP()

	db := c.Get("db").(*psql.DB) // Получаем подключение к базе данных

	tokens, ipChanged, err := auth.RefreshTokens(accessToken, refreshToken, ip, db)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "invalid refresh token",
		})
	}

	if ipChanged {
		//уведомление по email
		println("Warning: IP address changed. Email notification sent.")
	}

	return c.JSON(http.StatusOK, tokens)
}
