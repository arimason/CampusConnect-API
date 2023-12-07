package main

import (
	"campusconnect-api/configs"
	"campusconnect-api/internal/infra/data/pgclient"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
)

func errControl() {
	if r := recover(); r != nil {
		e := r.(error)
		log.Printf("Application shutdown caused by error: '%s'\n", e.Error())
	}
}

func cmdName(order int) string {
	switch order {
	case 0:
		return "up"
	case 1:
		return "down"
	case 2:
		return "reverse"
	case 3:
		return "force"
	default:
		return "unknown"
	}
}

func main() {
	// capture panic error
	defer errControl()
	// vars
	var (
		migrateCmd     int
		migrateVersion int
		yamlFile       string
	)
	// flags
	flag.StringVar(&yamlFile, "c", "configs/app.yaml", "Load application settings from path file name")
	flag.IntVar(&migrateCmd, "cmd", 0, "Comands: up=0, down=1, reverse=2, force=3")
	flag.IntVar(&migrateVersion, "force", 1, "For force version. Default 1")
	flag.Parse()
	// lendo as configurações
	cfg, err := configs.LoadConfigs("configs/app.yaml")
	if err != nil {
		panic(err)
	}
	log.Printf("Load configurations in %s", yamlFile)
	// criando configuração para conexão com o postgres
	log.Printf("Configuring database %s in %s:%d", cfg.DBName, cfg.DBHost, cfg.DBPort)
	pgConfig := &pgclient.Config{
		Host:     cfg.DBHost,
		Port:     cfg.DBPort,
		User:     cfg.DBUser,
		Password: cfg.DBPassword,
		Name:     cfg.DBName,
		Driver:   cfg.DBDriver,
	}
	pgConn := &pgclient.Conn{}
	pgConn.InitConn(pgConfig)
	// configure database
	log.Printf("Connection database")
	// configure pg drive
	driver, err := postgres.WithInstance(pgConn.DBConn, &postgres.Config{})
	if err != nil {
		panic(err)
	}
	// set directory for migrations
	path, _ := os.Getwd()
	m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s/migrations", path), "postgres", driver)
	if err != nil {
		panic(err)
	}
	//
	v, d, err := m.Version()
	if err != nil && !errors.Is(err, migrate.ErrNilVersion) {
		panic(err)
	}
	log.Printf("Current version: %d", v)
	log.Printf("Dirty: %t", d)
	// migration of commands
	log.Printf("Executing command: %s", cmdName(migrateCmd))
	err = nil
	if migrateCmd == 0 {
		log.Printf("Up to all")
		err = m.Up()
	} else if migrateCmd == 1 {
		log.Printf("Down to all")
		err = m.Down()
	} else if migrateCmd == 2 {
		log.Printf("Reversing to %d", v-1)
		err = m.Steps(-1)
	} else if migrateCmd == 3 {
		log.Printf("Forcing version: %d", migrateVersion)
		err = m.Force(migrateVersion)
	} else {
		log.Printf("Unknown command")
	}
	// fail
	if err != nil {
		panic(err)
	}
}
