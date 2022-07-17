package userusecase

import "errors"

var ErrUserNotFound = errors.New("User not found")
var ErrEmailAlreadyUsed = errors.New("Email already used")
var ErrInvalidPassword = errors.New("Invalid password")
