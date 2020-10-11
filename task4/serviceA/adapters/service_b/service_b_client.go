package service_b

import (
	"context"
)

type phoneService struct {

}

func NewPhoneService() *phoneService {
	return &phoneService{

	}
}
func (ps *phoneService) CheckPhone(ctx context.Context, phone string) (bool, error) {
	return true, nil
}