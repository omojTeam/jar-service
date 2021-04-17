package domainerrors

import "errors"

var (
	ErrInternalServerError = errors.New("Internal Server Error")
	ErrNotFound            = errors.New("Your requested Item is not found")
	ErrRecordNotFound      = errors.New("Record not found")
	ErrConflict            = errors.New("Conflict detected")
	ErrBadParamInput       = errors.New("Given parameter is not valid")
	ErrUnauthorized        = errors.New("Unauthorized")
	ErrUnprocessableEntity = errors.New("Unprocessable Entity")
	ErrForbidden           = errors.New("Forbidden")
	ErrNoCardsLeftToday    = errors.New("There are no cards left for today!")
	ErrNoCardsLeft         = errors.New("There are no cards left, you can reset your jar!")
	ErrEmptyTitle          = errors.New("Title can not be empty!")
	ErrEmptyCardsPerDay    = errors.New("CardsPerDay can not be empty!")
	ErrEmptyRecipientEmail = errors.New("RecipientEmail can not be empty!")
	ErrEmptyCardArray      = errors.New("Card array can not be empty!")
	ErrCardsPerDayTooLarge = errors.New("CardsPerDay can not be larger than total number of cards!")
)
