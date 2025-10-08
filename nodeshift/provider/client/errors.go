package client

import (
	"errors"
)

var (
	errNotImplemented = errors.New("update is not implemented")
	errStatusFailed   = errors.New("status failed")
	ErrPropertyEmpty  = errors.New("empty property")
	ErrPropertyType   = errors.New("incorrect property type")
	ErrPropertyCast   = errors.New("property cast")
	errRequestFailed  = errors.New("request failed")
)
