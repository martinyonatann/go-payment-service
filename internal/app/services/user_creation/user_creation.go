package user_creation

import "context"

type (
	UserCreationSVC interface {
		UserCreation(ctx context.Context, req UserCreationRequest) (res UserCreationResponse, err error)
	}

	UserCreationRequest struct {
	}

	UserCreationResponse struct {
	}

	service struct {
	}
)

func New(impl service) UserCreationSVC {
	return &impl
}

func (x *service) UserCreation(ctx context.Context, req UserCreationRequest) (res UserCreationResponse, err error) {

	return res, err
}
