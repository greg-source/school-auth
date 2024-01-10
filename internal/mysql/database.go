package mysql

import (
	"context"
	"database/sql"

	"school-auth/internal"
	"school-auth/internal/mysql/sqlc"
)

var _ internal.Database = &Database{}

type Database struct {
	db *sql.DB
	*sqlc.Queries
}

func New(db *sql.DB) *Database {
	return &Database{db: db, Queries: sqlc.New(db)}
}

func (d *Database) Name() string {
	return "mysql"
}

func (d *Database) GetProfiles(ctx context.Context) ([]internal.Profile, error) {
	sqlcProfiles, err := d.Queries.GetProfiles(ctx)
	if err != nil {
		return nil, err
	}

	internalProfiles := make([]internal.Profile, len(sqlcProfiles))
	for i, profile := range sqlcProfiles {
		internalProfiles[i] = internal.Profile{
			ID:        int(profile.ID.Int64),
			Username:  profile.Username.String,
			FirstName: profile.FirstName,
			LastName:  profile.LastName,
			City:      profile.City,
			School:    profile.School,
		}
	}

	return internalProfiles, nil
}

func (d *Database) GetProfile(ctx context.Context, username string) (*internal.Profile, error) {
	sqlcProfile, err := d.Queries.GetProfile(ctx, username)
	if err != nil {
		return nil, err
	}

	internalProfile := &internal.Profile{
		ID:        int(sqlcProfile.ID.Int64),
		Username:  sqlcProfile.Username.String,
		FirstName: sqlcProfile.FirstName,
		LastName:  sqlcProfile.LastName,
		City:      sqlcProfile.City,
		School:    sqlcProfile.School,
	}

	return internalProfile, nil
}

func (d *Database) CountAPIKey(ctx context.Context, key string) (int, error) {
	count, err := d.Queries.CountApiKey(ctx, key)
	if err != nil {
		return 0, err
	}

	return int(count), nil
}
