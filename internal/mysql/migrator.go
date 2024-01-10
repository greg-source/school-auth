package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	migrate "github.com/rubenv/sql-migrate"
	"school-auth/config"
)

func (d *Database) MigrationsUp(migrationDir string) error {
	migrations := &migrate.FileMigrationSource{
		Dir: migrationDir,
	}

	_, err := migrate.Exec(d.db, d.Name(), migrations, migrate.Up)

	return err
}

func NewMySQLConnection(config config.Config, retry int) (*sql.DB, error) {
	for i := 0; i < retry; i++ {
		time.Sleep(time.Second)

		dbURLTemplate := "%s:%s@tcp(%s:%d)/%s?parseTime=true"
		dbURL := fmt.Sprintf(
			dbURLTemplate,
			config.MySQLUser,
			config.MySQLPassword,
			config.MySQLHost,
			config.MySQLPort,
			config.MySQLDatabase)

		connection, err := sql.Open("mysql", dbURL)
		if err != nil {
			continue
		}

		err = connection.Ping()
		if err != nil {
			continue
		}

		return connection, nil
	}

	return nil, errors.New("database timeout")
}
