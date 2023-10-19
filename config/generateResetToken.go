package config

import "github.com/google/uuid"

func GenerateUniqueToken() (string, error) {
	// Buat token unik, misalnya dengan menggunakan UUID
	token := uuid.New()
	return token.String(), nil
}
