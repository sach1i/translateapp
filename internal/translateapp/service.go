package translateapp

import (
	"context"
	"go.uber.org/zap"
	"translateapp/internal/libretranslate"
)

type Service struct {
	Client libretranslate.ClientInterface
	Logger *zap.SugaredLogger
}

type Servicer interface {
	Languages(ctx context.Context) (*Response, error)
	Translate(ctx context.Context, input interface{}) (*Response, error)
}

func NewService(client libretranslate.ClientInterface, logger *zap.SugaredLogger) Servicer {
	return &Service{
		Client: client,
		Logger: logger,
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

func (s *Service) Translate(ctx context.Context, input interface{}) (*Response, error) {
	s.Logger.Debug("Service triggered a client")

	translation, err := s.Client.Translate(ctx, input)

	if err != nil {
		s.Logger.Info("Client returned error")
		return nil, err
	}

	var response Response

	response.Data = *translation

	s.Logger.Info("Client was successful")

	return &response, nil

}
