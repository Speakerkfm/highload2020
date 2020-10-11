package internal

import (
	"context"

	"serviceA/adapters/domain/errors"
)

type codeVerifier struct {
	phoneCodes map[string]string
}

func NewCodeVerifier() *codeVerifier {
	return &codeVerifier{
		phoneCodes: map[string]string{
			"+79125947072": "1234",
		},
	}
}

func (cc *codeVerifier) VerifyCode(ctx context.Context, phone string, code string) (bool, error) {
	validCode, ok := cc.phoneCodes[phone]
	if !ok {
		return false, errors.New("code not exists")
	}
	if code != validCode {
		return false, errors.New("invalid code")
	}

	return true, nil
}
