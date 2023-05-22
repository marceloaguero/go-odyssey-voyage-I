package graph

//go:generate go run github.com/99designs/gqlgen

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"github.com/marceloaguero/go-odyssey-voyage-I/reviews/usecase"
)

type Resolver struct {
	usecase usecase.Usecase
}

func NewResolver(uc usecase.Usecase) *Resolver {
	return &Resolver{
		usecase: uc,
	}
}
