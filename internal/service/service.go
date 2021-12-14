package service

import (
	"context"
	"log"
	"translateapp/internal/client"
	m "translateapp/internal/models"
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

func (ser *Service) Languages(ctx context.Context) (*m.ClientSuccessResponse, *m.ClientErrResponse) {
	response, err := ser.client.GetLanguages(ctx)
	log.Printf("in service %v,%v", response, err)
	if err != nil {
		return nil, err
	}
	return response, nil
}
