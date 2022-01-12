package mocks

import (
	"context"
	"github.com/stretchr/testify/mock"
	"translateapp/internal/libretranslate"
)

type Client struct {
	mock.Mock
}

func (t *Client) GetLanguages(ctx context.Context) (*libretranslate.LanguageList, error) {
	args := t.Called(ctx)
	return args.Get(0).(*libretranslate.LanguageList), args.Error(1)
}

func (t *Client) Translate(ctx context.Context, input libretranslate.Input) (*libretranslate.Translation, error) {
	args := t.Called(ctx, input)
	return args.Get(0).(*libretranslate.Translation), args.Error(1)
}
