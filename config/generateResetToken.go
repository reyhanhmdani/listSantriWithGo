package config

import "github.com/google/uuid"

func generateResetToken() (string, error) {
	// Buat token unik, misalnya dengan menggunakan UUID
	token := uuid.NewString()
	return token, nil
}
