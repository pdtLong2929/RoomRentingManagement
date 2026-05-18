package domain

import "errors"

// Owner Errors definition
var ErrOwnerNotFound = errors.New("Owner not found")

// 400 Bad Request scenarios
var ErrInvalidEmail = errors.New("provided email address format is invalid")
var ErrPhoneNumberTooShort = errors.New("phone number must be at least 10 digits")
var ErrNameRequired = errors.New("owner name cannot be left blank")

// 409 Conflict scenarios
var ErrEmailAlreadyExists = errors.New("an owner account with this email already exists")
