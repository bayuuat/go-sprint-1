package domain

import "errors"

var ErrUserNotFound = errors.New("user not found")
var ErrIdentityNumberNotFound = errors.New("identity number not found")
var ErrDepartmentNotFound = errors.New("department not found")
var ErrInvalidCredential = errors.New("invalid credential")
var ErrInvalidActionItem = errors.New("action unknown")
var ErrEmailExists = errors.New("email already exists")
var ErrEmployeeExists = errors.New("employee already exists")
