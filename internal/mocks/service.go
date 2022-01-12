package mocks

import (
	"context"
	"github.com/stretchr/testify/mock"
	"translateapp/internal/translateapp"
)

type Service struct {
	mock.Mock
}

func (t *Service) Languages(ctx context.Context) (*translateapp.Response, error) {
	args := t.Called(ctx)
	return args.Get(0).(*translateapp.Response), args.Error(1)
}
func (t *Service) Translate(ctx context.Context, input translateapp.Input) (*translateapp.Response, error) {
	args := t.Called(ctx, input)
	return args.Get(0).(*translateapp.Response), args.Error(1)
}
