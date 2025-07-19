package generate_password

import (
	"crypto/rand"
	"math/big"
)

// letters - список допустимых символов в пароле
const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GeneratePassword(n int) (string, error) {
	password := make([]byte, n)

	for i := range password {
		val, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		idx := int(val.Int64())

		password[i] = letters[idx]
	}

	return string(password), nil
}
