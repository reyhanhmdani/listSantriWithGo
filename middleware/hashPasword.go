package middleware

import "golang.org/x/crypto/bcrypt"

// HashPassword menghasilkan hash dari password yang diberikan.
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPasswordHash memeriksa apakah password cocok dengan hash yang disimpan.
func CheckPasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
