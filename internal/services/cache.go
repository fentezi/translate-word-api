package services

import (
	"errors"
	"log/slog"

	"github.com/fentezi/translete-word/internal/models"
	"github.com/fentezi/translete-word/internal/repositories"
	customerrors "github.com/fentezi/translete-word/internal/utils/errors"
	"github.com/fentezi/translete-word/internal/utils/google"
)

type wordCacheImpl struct {
	wordRepository repositories.WordRepository
	logger         *slog.Logger
}

func NewCacheService(repository repositories.WordRepository, logger *slog.Logger) WordService {
	return &wordCacheImpl{
		wordRepository: repository,
		logger:         logger,
	}
}

func (w *wordCacheImpl) WordTranslate(in *models.AddWord) (string, error) {
	var value string
	w.logger.Info("Starting translation process", "word", in.Word)

	value, err := w.wordRepository.Get(in.Word)
	if err == nil {
		w.logger.Info("Word found in cache", "word", in.Word, "translation", value)
		return value, nil
	}

	if !errors.Is(err, customerrors.ErrKeyNotFound) {
		w.logger.Error("Failed to retrieve word from cache", "word", in.Word, "error", err)
		return "", err
	}

	w.logger.Warn("Word not found in cache, translating", "word", in.Word)
	value, err = google.Translate(in.Word)
	if err != nil {
		w.logger.Error("Translation failed", "word", in.Word, "error", err)
		return "", err
	}

	w.logger.Info("Translation successful, caching result", "word", in.Word, "translation", value)
	if err = w.wordRepository.Set(in.Word, value); err != nil {
		w.logger.Error("Failed to cache translation", "word", in.Word, "error", err)
		return "", err
	}

	return value, nil
}
