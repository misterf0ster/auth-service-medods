package handler

import (
	m "auth-service-medods/internal/models"
	auth "auth-service-medods/internal/service"
	psql "auth-service-medods/internal/storage"
	"auth-service-medods/pkg/logger"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TokenHandler struct {
	DB *psql.DB
}

func CreateToken(db *psql.DB) *TokenHandler {
	return &TokenHandler{DB: db}
}

// GET
func (h *TokenHandler) GetAccessRefresh(c *gin.Context) {
	id := c.Param("id")
	sql := "SELECT ip, token_hash, user_id FROM refresh_token"

	rows, err := h.DB.Psql.Query(context.Background(), sql)
	if err != nil {
		logger.LogError("Database connect fals", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": "Database connect fals",
			"Status":  "Error",
		})
		return
	}
	defer rows.Close()

	var refresh []m.Tokens
	for rows.Next() {
		var u m.Tokens
		if err := rows.Scan(&u.AccessToken, &u.RefreshToken, id); err == nil {
			refresh = append(refresh, u)
		}
		return
	}

	logger.LogInfo("Tokens get successfully")
	c.JSON(http.StatusOK, refresh)
}

// POST
func (h *TokenHandler) PostRefresh(c *gin.Context) {
	accessToken := c.Param("access_token")
	refreshToken := c.Param("refresh_token")
	ip := c.Param("ip")

	tokens, ipChanged, err := auth.RefreshTokens(accessToken, refreshToken, ip)
	if err != nil {
		c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "invalid refresh token",
		})
	}

	if ipChanged {
		//уведомление по email
		println("Warning: IP address changed. Email notification sent.")
	}

	c.JSON(http.StatusOK, tokens)
}
