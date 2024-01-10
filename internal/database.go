package internal

import "context"

type Database interface {
	GetProfiles(ctx context.Context) ([]Profile, error)
	GetProfile(ctx context.Context, username string) (*Profile, error)
	CountAPIKey(ctx context.Context, key string) (int, error)
	Migrator
}

type Migrator interface {
	MigrationsUp(migrationDir string) error
	Name() string
}
