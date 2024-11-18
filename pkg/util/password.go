package util

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/argon2"
)

func GenerateSalt() (string, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return base64.RawStdEncoding.EncodeToString(salt), nil
}

func HashPassword(password, salt string) (string, error) {
	saltbyte, err := base64.RawStdEncoding.DecodeString(salt)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), saltbyte, 1, 64*1024, 4, 32)

	return base64.RawStdEncoding.EncodeToString(hash), nil
}

func CheckPassword(passwordHash, password, salt string) bool {

	saltbyte, err := base64.RawStdEncoding.DecodeString(salt)
	if err != nil {
		return false
	}

	oldHash, err := base64.RawStdEncoding.DecodeString(password)
	if err != nil {
		return false
	}

	// 重新生成哈希
	newHash := argon2.IDKey([]byte(password), saltbyte, 1, 64*1024, 4, 32)

	// 比较哈希
	if len(oldHash) != len(newHash) {
		return false
	}

	for i := range oldHash {
		if oldHash[i] != newHash[i] {
			return false
		}
	}

	return true
}
