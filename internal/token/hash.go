package token

import (
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// генерация случайного UUID
func GenerateUUID(length int) string {
	return uuid.New().String()[:length]
}

// хэширую refresh токен
func HashRefreshToken(refreshToken string) (string, error) {
	hashToken, err := bcrypt.GenerateFromPassword([]byte(refreshToken), bcrypt.MaxCost)
	if err != nil {
		return "", fmt.Errorf("could not hash refresh token: %w", err)
	}
	return string(hashToken), nil
}

// сравниваю хэшированный токен с переданным
func CompareHashAndPassword(hashToken, refreshToken string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashToken), []byte(refreshToken))
	return err == nil
}
