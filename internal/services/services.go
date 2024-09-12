package services

import (
	"log/slog"

	"github.com/fentezi/translete-word/internal/repositories"
)

type Service struct {
	CacheService *CacheService
}

func NewService(repository repositories.Repository, logger *slog.Logger) *Service {
	return &Service{
		CacheService: NewCacheService(repository, logger),
	}
}
