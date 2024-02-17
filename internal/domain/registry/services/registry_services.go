package services

import (
	"strings"
	"words-api/internal/domain"
	"words-api/internal/domain/registry/repository"
)

func Create(word *domain.Word) error {
	formatWord(word)

	if err := repository.Create(word); err != nil {
		return err
	}

	return nil
}

func Update(currentWord *domain.Word) error {
	formatWord(currentWord)

	if err := repository.Update(currentWord); err != nil {
		return err
	}

	return nil
}

func Delete(wordID int64) error {
	if err := repository.Delete(wordID); err != nil {
		return err
	}

	return nil
}

func formatWord(word *domain.Word) {
	word.Word = strings.TrimSpace(word.Word)
	word.Word = strings.ToLower(word.Word)
	word.Meaning = strings.TrimSpace(word.Meaning)
}
