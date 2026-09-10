package auth

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

func HashPassword(password string) (string, error) {

	salt := make([]byte, 16)

	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey(
		[]byte(password),
		salt,
		1,
		64*1024,
		4,
		32,
	)

	return fmt.Sprintf(
		"$argon2id$v=19$m=65536,t=1,p=4$%s$%s",
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(hash),
	), nil
}

func VerifyPassword(password, encodedHash string) bool {

	parts := strings.Split(
		encodedHash,
		"$",
	)

	if len(parts) != 6 {
		return false
	}

	salt, err := base64.RawStdEncoding.DecodeString(
		parts[4],
	)

	if err != nil {
		return false
	}

	hashFromDB, err := base64.RawStdEncoding.DecodeString(
		parts[5],
	)

	if err != nil {
		return false
	}

	hashToCompare := argon2.IDKey(
		[]byte(password),
		salt,
		1,
		64*1024,
		4,
		32,
	)

	return subtle.ConstantTimeCompare(
		hashToCompare,
		hashFromDB,
	) == 1
}
