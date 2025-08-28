package auth

import (
	"net/http"
	"testing"
)

func TestAPIKey(t *testing.T) {
	// Тест 1: Нет заголовка Authorization
	t.Run("No Authorization Header", func(t *testing.T) {
		headers := make(http.Header)
		apiKey, err := GetAPIKey(headers)

		if err != ErrNoAuthHeaderIncluded {
			t.Errorf("Expected error %v, got %v", ErrNoAuthHeaderIncluded, err)
		}
		if apiKey != "" {
			t.Errorf("Expected empty API key, got %v", apiKey)
		}
	})

	// Тест 2: Неправильный формат заголовка Authorization
	t.Run("Malformed Authorization Header", func(t *testing.T) {
		headers := make(http.Header)
		headers.Set("Authorization", "ApiKey123456")

		apiKey, err := GetAPIKey(headers)

		if err == nil || err.Error() != "malformed authorization header" {
			t.Errorf("Expected 'malformed authorization header' error, got %v", err)
		}
		if apiKey != "" {
			t.Errorf("Expected empty API key, got %v", apiKey)
		}
	})

	// Тест 3: Правильный заголовок Authorization
	t.Run("Valid Authorization Header", func(t *testing.T) {
		headers := make(http.Header)
		headers.Set("Authorization", "ApiKey abc123")

		apiKey, err := GetAPIKey(headers)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if apiKey != "abc123" {
			t.Errorf("Expected API key 'abc123', got %v", apiKey)
		}
	})
}
