package emojerrors

import (
	"errors"
	"fmt"
)

type Error struct {
	err   error
	emoji string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s %s", e.err.Error(), e.emoji)
}

func (e *Error) Unwrap() error {
	return e.err
}

func New(text string, emoji string) error {
	return &Error{
		err:   errors.New(text),
		emoji: emoji,
	}
}

func Wrap(err error, emoji string) error {
	return &Error{
		err:   err,
		emoji: emoji,
	}
}
