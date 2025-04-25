package storage

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

// Хранение хэшированного Refresh токена
func (db *DB) StoreRefreshToken(userID, hashedToken, ip string) error {
	_, err := db.Psql.Exec(context.Background(), `
		INSERT INTO refresh_tokens (user_id, token, ip)
		VALUES ($1, $2, $3)`,
		userID, hashedToken, ip)
	return err
}

// Проверка Refresh токена
func (db *DB) VerifyRefreshToken(userID, refreshToken string) (bool, string, error) {
	var storedToken, storedIP string
	err := db.Psql.QueryRow(context.Background(), `
		SELECT token, ip FROM refresh_tokens WHERE user_id = $1`,
		userID).Scan(&storedToken, &storedIP)
	if err != nil {
		return false, "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedToken), []byte(refreshToken))
	if err != nil {
		return false, "", nil
	}

	return true, storedIP, nil
}
