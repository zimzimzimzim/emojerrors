package emojerrors

import (
	"errors"
	"fmt"
)

type eError struct {
	err   error
	emoji string
}

func (e *eError) Error() string {
	return fmt.Sprintf("%s %s", e.err.Error(), e.emoji)
}

func (e *eError) Unwrap() error {
	return e.err
}

func New(text string, emoji string) error {
	return &eError{
		err:   errors.New(text),
		emoji: emoji,
	}
}

func Wrap(err error, emoji string) error {
	return &eError{
		err:   err,
		emoji: emoji,
	}
}
