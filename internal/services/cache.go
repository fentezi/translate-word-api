package services

import (
	"errors"
	"log/slog"

	"github.com/fentezi/translete-word/internal/repositories"
	customerrors "github.com/fentezi/translete-word/internal/utils/errors"
	"github.com/fentezi/translete-word/internal/utils/google"
)

type CacheService struct {
	Repository repositories.Repository
	logger     *slog.Logger
}

func NewCacheService(repository repositories.Repository, logger *slog.Logger) *CacheService {
	return &CacheService{
		Repository: repository,
		logger:     logger,
	}
}

func (s *CacheService) GetTranslateWord(word string) (string, error) {
	var value string
	s.logger.Info("Starting translation process", "word", word)

	// Попытка получить значение из кеша
	value, err := s.Repository.Get(word)
	if err == nil {
		s.logger.Info("Word found in cache", "word", word, "translation", value)
		return value, nil
	}

	if !errors.Is(err, customerrors.ErrKeyNotFound) {
		s.logger.Error("Failed to retrieve word from cache", "word", word, "error", err)
		return "", err
	}

	// Слово не найдено в кеше, попытка перевести
	s.logger.Warn("Word not found in cache, translating", "word", word)
	value, err = google.Translate(word)
	if err != nil {
		s.logger.Error("Translation failed", "word", word, "error", err)
		return "", err
	}

	// Сохранение результата перевода в кеш
	s.logger.Info("Translation successful, caching result", "word", word, "translation", value)
	if err = s.Repository.Set(word, value); err != nil {
		s.logger.Error("Failed to cache translation", "word", word, "error", err)
		return "", err
	}

	return value, nil
}
