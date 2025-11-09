package cryptography

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"

	"golang.org/x/crypto/argon2"
)

func GenerateSalt(length int) ([]byte, error) {
	salt := make([]byte, length)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}
	return salt, nil
}

func HashPassword(password string, salt []byte) string {
	hash := argon2.IDKey([]byte(password), salt, 3, 64*1024, 4, 32)
	return base64.RawStdEncoding.EncodeToString(hash)
}

func VerifyPassword(password string, salt []byte, hash string) bool {
	newHash := HashPassword(password, salt)
	decodedHash, _ := base64.RawStdEncoding.DecodeString(hash)
	decodedNewHash, _ := base64.RawStdEncoding.DecodeString(newHash)
	return subtle.ConstantTimeCompare(decodedHash, decodedNewHash) == 1
}
