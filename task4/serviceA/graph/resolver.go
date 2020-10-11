package graph

import (
	"context"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	phoneVerifier phoneVerifier
}

func NewResolver(phoneVerifier phoneVerifier) *Resolver {
	return &Resolver{
		phoneVerifier: phoneVerifier,
	}
}

type phoneVerifier interface {
	VerifyPhone(ctx context.Context, phone string, code string) (bool, error)
}