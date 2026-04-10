package auth

import "errors"

var ErrUserAlreayExists error = errors.New("user already exists!")
var ErrInternalWhileRegisterUser error = errors.New("internal error while register user!")
var ErrIncorrectUserData error = errors.New("incorrect user data")
