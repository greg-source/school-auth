package service

import (
	"context"

	"school-auth/internal"
)

var _ internal.Service = &Service{}

type Service struct {
	internal.Database
}

func New(database internal.Database) *Service {
	return &Service{Database: database}
}

func (s *Service) Get(ctx context.Context) ([]internal.Profile, error) {
	return s.GetProfiles(ctx)
}

func (s *Service) GetByUsername(ctx context.Context, username string) (*internal.Profile, error) {
	return s.GetProfile(ctx, username)
}

func (s *Service) ValidateAPIKey(ctx context.Context, key string) (bool, error) {
	apiKey, err := s.CountAPIKey(ctx, key)
	if err != nil {
		return false, err
	}

	if apiKey == 0 {
		return false, nil
	}

	return true, nil
}
