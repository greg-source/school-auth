package internal

import "context"

type Profile struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	City      string `json:"city"`
	School    string `json:"school"`
}

type Service interface {
	Get(ctx context.Context) ([]Profile, error)
	GetByUsername(ctx context.Context, username string) (*Profile, error)
	ValidateAPIKey(ctx context.Context, key string) (bool, error)
}
