package mocks

import (
	"context"
	"github.com/stretchr/testify/mock"
	"translateapp/internal/libretranslate"
)

type Translator struct {
	mock.Mock
}

func (t *Translator) Translate(ctx context.Context, input libretranslate.Input) (string, error) {
	args := t.Called(ctx, input)
	return args.Get(0).(string), args.Error(1)
}
