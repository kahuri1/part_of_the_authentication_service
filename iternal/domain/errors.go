package domain

import "errors"

var (
	ErrUserAlreadyExists       = errors.New("user with such email already exists")
	ErrVerificationCodeInvalid = errors.New("verification code is invalid")
	ErrSessionOpen             = errors.New("session is open")
)
