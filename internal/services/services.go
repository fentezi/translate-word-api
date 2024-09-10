package services

import (
	"github.com/fentezi/translete-word/internal/repositories"
	"github.com/fentezi/translete-word/internal/utils/google"
	"github.com/redis/go-redis/v9"
)

type Service struct {
	repositories.Repository
}

func NewService(repository repositories.Repository) *Service {
	return &Service{
		Repository: repository,
	}
}

func (s *Service) GetTranslateWord(word string) (string, error) {
	var value string

	value, err := s.Repository.Get(word)
	if err != nil {
		if err == redis.Nil {
			value, err = google.Translate(word)
			if err != nil {
				return "", err
			}

			err = s.Repository.Set(word, value)
			if err != nil {
				return "", err
			}
		} else {
			return "", err
		}

	}

	return value, nil
}
