package database

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/pbkdf2"
)

const (
	SaltSize   = 16
	Iterations = 10000
	KeyLength  = 32
)

func hashPassword(password string) (string, error) {
	salt := make([]byte, SaltSize)
	_, err := rand.Read(salt)
	if err != nil {
		return "", fmt.Errorf("failed to generate salt: %v", err)
	}
	hash := pbkdf2.Key([]byte(password), salt, Iterations, KeyLength, sha256.New)
	return fmt.Sprintf("%x:%x", salt, hash), nil // Store as "salt:hash"
}

func verifyPassword(storedPassword, providedPassword string) (bool, error) {
	parts := [2]string{}
	_, err := fmt.Sscanf(storedPassword, "%x:%x", &parts[0], &parts[1])
	if err != nil {
		return false, fmt.Errorf("failed to parse stored password: %v", err)
	}

	salt, err := hex.DecodeString(parts[0])
	if err != nil {
		return false, fmt.Errorf("failed to decode salt: %v", err)
	}

	storedHash, err := hex.DecodeString(parts[1])
	if err != nil {
		return false, fmt.Errorf("failed to decode stored hash: %v", err)
	}

	providedHash := pbkdf2.Key([]byte(providedPassword), salt, Iterations, len(storedHash), sha256.New)

	if subtle.ConstantTimeCompare(storedHash, providedHash) == 1 {
		return true, nil
	}
	return false, nil
}
