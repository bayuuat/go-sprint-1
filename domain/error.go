package domain

import "errors"

var ErrUserNotFound = errors.New("user not found")
var ErrInvalidCredential = errors.New("invalid credential")
var ErrInvalidActionItem = errors.New("action unknown")
var ErrEmailExists = errors.New("email already exists")
