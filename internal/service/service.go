package service

import (
	"context"
	"translateapp/internal/client"
	"translateapp/internal/models"
)

type Service struct {
	client *client.Client
}

func NewService() *Service {
	ser := &Service{
		client.NewClient(),
	}
	return ser
}

func (s *Service) Languages(ctx context.Context) (*models.GetRes, *models.ClientErrResponse) {
	langList, err := s.client.GetLanguages(ctx)
	if err != nil {
		return nil, err
	}
	var response models.GetRes
	response.Data = *langList
	response.Code = 200
	response.Message = "success"
	return &response, nil
}

func (s *Service) Translate(ctx context.Context, input models.Input) (*models.PostRes, *models.ClientErrResponse) {
	translation, err := s.client.Translate(ctx, input)
	if err != nil {
		return nil, err
	}
	var response models.PostRes
	response.Data = *translation
	response.Code = 200
	response.Message = "success"
	return &response, nil
}
