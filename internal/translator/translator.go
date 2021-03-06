package translator

import (
	"context"
	"go.uber.org/zap"
	"translateapp/internal/cache"
	"translateapp/internal/libretranslate"
)

type Translator struct {
	Cache  cache.CacheInterface
	Client libretranslate.ClientInterface
	Logger *zap.SugaredLogger
}

type TranslateInterface interface {
	Translate(ctx context.Context, input libretranslate.Input) (string, error)
}

func NewTranslator(c cache.CacheInterface, client libretranslate.ClientInterface, logger *zap.SugaredLogger) *Translator {
	return &Translator{
		Cache:  c,
		Client: client,
		Logger: logger,
	}
}

// Tries to get value from cache, implements read-through approach
func (t *Translator) Translate(ctx context.Context, input libretranslate.Input) (string, error) {
	t.Logger.Info("Translate function triggered in module Translator")
	res := t.Cache.Get(input.Word)
	if len(res) > 0 {
		return res, nil
	} else {
		fromClient, err := t.Client.Translate(ctx, input)
		if err != nil {
			return "", err
		}
		translation := *fromClient
		t.Cache.Set(input.Word, translation.Text)
		return t.Cache.Get(input.Word), nil
	}
}
