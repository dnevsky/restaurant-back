package models

import "errors"

var (
	ErrUnauthorized          = errors.New("unauthorized")
	ErrInvalidRequestParams  = errors.New("invalid request parameters. see invalidArgs")
	ErrBrokenPipe            = errors.New("write: broken pipe")
	ErrConnectionResetByPeer = errors.New("connection reset by peer")
	ErrBadRequest            = errors.New("bad request")
	ErrAlreadyExists         = errors.New("already exists")
	ErrInvalidAuthCreds      = errors.New("bad email or password")
	ErrEmptyAuthHeader       = errors.New("empty auth header")
	ErrInvalidAuthHeader     = errors.New("invalid auth header")
	ErrTokenIsEmpty          = errors.New("token is empty")
	ErrAccessDenied          = errors.New("access denied")
	ErrNotFound              = errors.New("not found")
	ErrTypeAssertionFailed   = errors.New("type assertion failed")
)

type ErrType int32

const (
	ErrTypeDefault          ErrType = 0
	ErrTypeValidation       ErrType = 1
	ErrTypeCustomValidation ErrType = 2
)
