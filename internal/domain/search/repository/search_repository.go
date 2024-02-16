package repository

import (
	"words-api/internal/domain"
	"words-api/internal/util/mysql"
)

func GetByID(wordID int64) (domain.Word, error) {
	resultWord, err := mysql.ClientDB.Query(
		"SELECT word_id, word, meaning, date_created FROM quotes.words WHERE word_id = ?", wordID)
	if err != nil {
		return domain.Word{}, err
	}

	var quote domain.Word
	for resultWord.Next() {
		err = resultWord.Scan(&quote.WordID, &quote.Word, &quote.Meaning, &quote.DateCreated)
		if err != nil {
			return domain.Word{}, err
		}
	}

	return quote, nil
}

func GetByWord(word string) ([]domain.Word, error) {
	resultWord, err := mysql.ClientDB.Query(
		"SELECT word_id, word, meaning, date_created FROM quotes.words WHERE word LIKE ?", "%"+word+"%")
	if err != nil {
		return nil, err
	}

	var quotes []domain.Word
	for resultWord.Next() {
		var quote domain.Word

		err = resultWord.Scan(&quote.WordID, &quote.Word, &quote.Meaning, &quote.DateCreated)
		if err != nil {
			return nil, err
		}

		quotes = append(quotes, quote)
	}

	return quotes, nil
}
