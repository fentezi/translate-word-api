package services

import "github.com/fentezi/translete-word/internal/models"

type WordService interface {
	WordTranslate(in *models.AddWord) (string, error)
}
