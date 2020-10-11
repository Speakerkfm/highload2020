package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"serviceA/graph/generated"
	"serviceA/graph/model"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *mutationResolver) VerifyPhone(ctx context.Context, input model.PhoneVerificationRequest) (*model.PhoneVerificationResponse, error) {
	if input.Phone == "" {
		return nil, gqlerror.Errorf("phone is empty")
	}

	if input.Code == "" {
		return nil, gqlerror.Errorf("code is empty")
	}

	verified, err := r.phoneVerifier.VerifyPhone(ctx, input.Phone, input.Code)
	if err != nil {
		return nil, gqlerror.Errorf("fail to verify phone, cause: %s", err.Error())
	}

	return &model.PhoneVerificationResponse{Verified: verified}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
