package translateapp

import (
	"context"
	"go.uber.org/zap"
	"translateapp/internal/libretranslate"
	"translateapp/internal/translator"
)

type Service struct {
	Client     libretranslate.ClientInterface
	Logger     *zap.SugaredLogger
	Translator translator.TranslateInterface
}

type ServiceInterface interface {
	Languages(ctx context.Context) (*Response, error)
	Translate(ctx context.Context, input Input) (*Response, error)
}

func NewService(client libretranslate.ClientInterface, translator translator.TranslateInterface, logger *zap.SugaredLogger) *Service {
	return &Service{
		Client:     client,
		Translator: translator,
		Logger:     logger,
	}

}

func (s *Service) Languages(ctx context.Context) (*Response, error) {
	s.Logger.Debug("Service triggered a client")
	langList, err := s.Client.GetLanguages(ctx)

	if err != nil {
		s.Logger.Info("Client returned error")
		return nil, err
	}

	var response Response
	response.Data = *langList
	s.Logger.Info("Client was successful")
	return &response, nil
}

func (s *Service) Translate(ctx context.Context, input Input) (*Response, error) {
	s.Logger.Debug("Service triggered a client")
	//translation, err := s.Client.Translate(ctx, ConvertModel(input))
	translation, err := s.Translator.Translate(ctx, ConvertModel(input))
	if err != nil {
		s.Logger.Info("Translator couldn't translate the input due to third party service issue")
		return nil, err
	}

	var response Response
	response.Data = translation
	return &response, nil
}

func ConvertModel(input Input) libretranslate.Input {
	res := libretranslate.Input{
		Word:   input.Word,
		Source: input.Source,
		Target: input.Target,
	}
	return res
}
