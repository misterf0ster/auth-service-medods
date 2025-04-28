package service

import (
	m "auth-service-medods/internal/models"
	"auth-service-medods/internal/storage"
	t "auth-service-medods/internal/token"
	"fmt"
)

// генерация токенов access и refresh
func GenerateTokens(userID, ip string, db *storage.DB) (*m.Tokens, error) {
	// генерация access токена
	accessToken, err := t.GetAccessToken(userID, ip)
	if err != nil {
		return nil, err
	}

	// генерация refresh токена
	refreshToken, hashedRefreshToken, err := generateRefreshToken(userID, ip)
	if err != nil {
		return nil, err
	}

	// cохраняю хэш refresh токен в базу
	err = db.StoreRefreshToken(userID, hashedRefreshToken, ip)
	if err != nil {
		return nil, err
	}

	return &m.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// генерация refresh токена
func generateRefreshToken(userID, ip string) (string, string, error) {
	refreshToken := generateRandomString(64)
	hashedRefreshToken, err := HashRefreshToken(refreshToken)
	if err != nil {
		return "", "", err
	}

	return refreshToken, hashedRefreshToken, nil
}

// обновление токенов
func RefreshTokens(accessToken, refreshToken, ip string, db *storage.DB) (*m.Tokens, bool, error) {
	claims, err := t.ParseAccessToken(accessToken)
	if err != nil {
		return nil, false, fmt.Errorf("could not parse access token: %w", err)
	}

	userID := claims["user_id"].(string) // Получаем user_id из токена

	//проверка refresh токена
	valid, storedIP, err := db.VerifyRefreshToken(userID, refreshToken)
	if err != nil {
		return nil, false, fmt.Errorf("could not verify refresh token: %w", err)
	}
	if !valid {
		return nil, false, fmt.Errorf("invalid refresh token")
	}

	// IP изменился - отправить сообщение
	ipChanged := false
	if storedIP != ip {
		ipChanged = true
	}

	//генерация новых токенов
	tokens, err := GenerateTokens(userID, ip, db)
	if err != nil {
		return nil, ipChanged, fmt.Errorf("could not generate new tokens: %w", err)
	}

	return tokens, ipChanged, nil
}
