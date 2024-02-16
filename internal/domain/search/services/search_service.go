package services

import (
	"words-api/internal/domain"
	"words-api/internal/domain/search/repository"
)

func GetByID(quoteID int64) (domain.Word, error) {
	quote, err := repository.GetByID(quoteID)
	if err != nil {
		return domain.Word{}, err
	}

	return quote, nil
}

func GetByWord(keyword string) ([]domain.Word, error) {
	quote, err := repository.GetByWord(keyword)
	if err != nil {
		return nil, err
	}

	return quote, nil
}
