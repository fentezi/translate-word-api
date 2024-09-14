package services

import (
	"log/slog"
	"testing"

	"errors"

	"github.com/fentezi/translete-word/internal/models"
	"github.com/fentezi/translete-word/internal/repositories/mocks"
	customerrors "github.com/fentezi/translete-word/internal/utils/errors"
	"github.com/stretchr/testify/assert"
)

func TestWordTranslate(t *testing.T) {
	mockRepo := new(mocks.WordRepository)
	logger := slog.Default()
	service := NewCacheService(mockRepo, logger)

	t.Run("Word found in cache", func(t *testing.T) {
		mockRepo.On("Get", "hello").Return("привет", nil)

		result, err := service.WordTranslate(&models.AddWord{Word: "hello"})

		assert.NoError(t, err)
		assert.Equal(t, "привет", result)
		mockRepo.AssertCalled(t, "Get", "hello")
	})

	t.Run("Word not found in cache, translate successfully", func(t *testing.T) {
		mockRepo.On("Get", "world").Return("", customerrors.ErrKeyNotFound)
		mockRepo.On("Set", "world", "мир").Return(nil)

		result, err := service.WordTranslate(&models.AddWord{Word: "world"})

		assert.NoError(t, err)
		assert.Equal(t, "мир", result)
		mockRepo.AssertCalled(t, "Get", "world")
		mockRepo.AssertCalled(t, "Set", "world", "мир")
	})

	t.Run("Error retrieving word from cache", func(t *testing.T) {
		mockRepo.On("Get", "error_word").Return("", errors.New("cache error"))

		result, err := service.WordTranslate(&models.AddWord{Word: "error_word"})

		assert.Error(t, err)
		assert.Empty(t, result)
		mockRepo.AssertCalled(t, "Get", "error_word")
	})
}
