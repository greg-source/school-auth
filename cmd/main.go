package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"school-auth/config"
	"school-auth/internal"
	"school-auth/internal/mysql"
	"school-auth/internal/server"
	"school-auth/internal/service"
)

const migrationsDir = "migrations"

func main() {
	appConfig, err := config.Load("config")
	if err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	conn, err := mysql.NewMySQLConnection(appConfig, 5)
	if err != nil {
		logrus.Fatalf("error initializing mysql connection: %s", err.Error())
	}
	defer conn.Close()

	db := mysql.New(conn)
	err = db.MigrationsUp(migrationsDir)
	if err != nil {
		logrus.Fatalf("error applying migrations: %s", err.Error())
		return
	}

	var app internal.Server

	app = server.New(service.New(mysql.New(conn)))

	err = app.Start(appConfig.ApplicationPort)
	if err != nil {
		logrus.Fatalf("error initializing new server: %s", err.Error())
	}

	systemQuit := make(chan os.Signal, 1)
	signal.Notify(systemQuit, os.Interrupt)

	<-systemQuit

	logrus.Info("Graceful signint successful")
}
