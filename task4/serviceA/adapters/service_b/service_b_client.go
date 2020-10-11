package service_b

import (
	"context"
	"time"

	"google.golang.org/grpc"

	"serviceA/adapters/domain/errors"
)

type phoneService struct {
	addr string
}

func NewPhoneService(addr string) *phoneService {
	return &phoneService{
		addr: addr,
	}
}
func (ps *phoneService) CheckPhone(ctx context.Context, phone string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	conn, err := grpc.DialContext(ctx, ps.addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return false, errors.Wrap(err, "did not connect to serviceB")
	}
	defer conn.Close()

	c := NewPhoneServiceClient(conn)

	request := CheckPhoneRequest{
		Phone: phone,
	}

	// Contact the server and print out its response.
	ctx, cancel = context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	resp, err := c.CheckPhone(ctx, &request)
	if err != nil {
		return false, errors.Wrap(err, "could not send CheckPhone to serviceB")
	}

	return resp.Exists, nil
}
