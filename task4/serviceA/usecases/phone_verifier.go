package usecases

import (
	"context"

	"serviceA/adapters/domain/errors"
)

type PhoneService interface {
	CheckPhone(ctx context.Context, phone string) (bool, error)
}

type CodeVerifier interface {
	VerifyCode(ctx context.Context, phone string, code string) (bool, error)
}

type phoneVerifier struct {
	phoneService PhoneService
	codeVerifier CodeVerifier
}

func NewPhoneVerifier(phoneService PhoneService, codeVerifier CodeVerifier) *phoneVerifier {
	return &phoneVerifier{
		phoneService: phoneService,
		codeVerifier: codeVerifier,
	}
}

func (pv *phoneVerifier) VerifyPhone(ctx context.Context, phone string, code string) (bool, error) {
	phoneExists, err := pv.phoneService.CheckPhone(ctx, phone)
	if err != nil {
		return false, errors.Wrap(err, "fail to check phone")
	}

	if !phoneExists {
		return false, errors.New("phone doesn't exist")
	}

	codeVerified, err := pv.codeVerifier.VerifyCode(ctx, phone, code)
	if err != nil {
		return false, errors.Wrap(err, "fail to verify code")
	}

	return codeVerified, nil
}
