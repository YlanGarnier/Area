package authentication

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func GenerateHashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to generate hash: %w", err)
	}
	return string(hash), nil
}

func GenerateOAT(hash string) (string, error) {
	hashMd5 := md5.New()
	_, err := hashMd5.Write([]byte(hash))
	if err != nil {
		return "", fmt.Errorf("failed to write hash: %w", err)
	}
	token := hex.EncodeToString(hashMd5.Sum(nil))
	return token, nil
}

func HashOAT(hash string) (string, error) {
	hashMd5 := md5.New()
	_, err := hashMd5.Write([]byte(hash))
	if err != nil {
		return "", fmt.Errorf("failed to write hash: %w", err)
	}
	token := hex.EncodeToString(hashMd5.Sum(nil))
	return token, nil
}

func CompareHashPassword(hash string, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err == nil {
		return true, nil
	} else if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return false, nil
	} else {
		return false, fmt.Errorf("failed to compare hash and password: %w", err)
	}
}
