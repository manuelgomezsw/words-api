package repository

import (
	"words-api/internal/domain"
	"words-api/internal/util/mysql"
)

func Create(newWord *domain.Word) error {
	newRecord, err := mysql.ClientDB.Exec(
		"INSERT INTO quotes.words (word, meaning) VALUES (?, ?)",
		newWord.Word,
		newWord.Meaning,
	)
	if err != nil {
		return err
	}

	newWord.WordID, err = newRecord.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func Update(currentWord *domain.Word) error {
	_, err := mysql.ClientDB.Exec(
		"UPDATE quotes.words SET word = ?, meaning = ?  WHERE word_id = ?",
		currentWord.Word,
		currentWord.Meaning,
		currentWord.WordID,
	)
	if err != nil {
		return err
	}

	return nil
}

func Delete(wordID int64) error {
	_, err := mysql.ClientDB.Exec(
		"DELETE FROM quotes.words WHERE word_id = ?",
		wordID,
	)
	if err != nil {
		return err
	}

	return nil
}
