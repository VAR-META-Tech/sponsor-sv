package models

import "errors"

var ErrInvalidLenMsg error = errors.New("invalid sponsor message length")
var ErrInvalidCallerMsg error = errors.New("caller is calling to a bad realm")
var ErrOnExecuteMsg error = errors.New("error on executing message")
