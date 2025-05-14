package main

import (
	"errors"
)

var (
	ErrInvalidParameter = errors.New("invalid parameter")
)

type mutationResolver struct {
	server *Server
}
