package config

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

var JwtKey = []byte(os.Getenv("JWT_KEY"))

type Claims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	UserID   int64  `json:"user_id"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

// Membuat token JWT
func CreateJWTToken(username, email string, userID int64, role string) (string, error) {
	expirationTime := time.Now().Add(time.Minute * 60)

	// Atur payload token
	claims := &Claims{
		Username: username,
		Email:    email,
		UserID:   userID,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Buat token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Simpan token dalam string dengan mengenkripsi menggunakan secret key
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
