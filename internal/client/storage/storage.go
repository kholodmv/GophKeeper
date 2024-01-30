package storage

import (
	"fmt"
	"os"
)

const FilePath = "temp/auth_token.txt"

type TokenService interface {
	SaveToken(token string) error
	LoadToken(token string) error
	GetToken(token string) (string, error)
}

type TokenStorage struct {
	filePath string
}

func New() *TokenStorage {
	return &TokenStorage{
		filePath: FilePath,
	}
}

// SaveToken saves the token to a file.
func (t *TokenStorage) SaveToken(token string) error {
	file, err := os.OpenFile(t.filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	err = os.WriteFile(t.filePath, []byte(token), 0644)
	if err != nil {
		return fmt.Errorf("error when saving token to file: %v", err)
	}
	return nil
}

// LoadToken loads a token from a file.
func (t *TokenStorage) LoadToken() (string, error) {
	content, err := os.ReadFile(t.filePath)
	if err != nil {
		return "", fmt.Errorf("error when loading a token from a file: %v", err)
	}
	return string(content), nil
}

// GetToken returns the current token.
func (t *TokenStorage) GetToken() string {
	tokenContent, err := t.LoadToken()
	if err != nil {
		fmt.Println("Error when receiving token:", err)
	}
	return tokenContent
}
