package validator

import "errors"

var (
	ErrDownloadFailed = errors.New("database download failed")
	ErrDatabaseWrite  = errors.New("database write failed")
	ErrDatabaseRead = errors.New("database read failed")
	ErrCardNotFound = errors.New("card not found")
	ErrInvalidDeck = errors.New("deck invalid")
)