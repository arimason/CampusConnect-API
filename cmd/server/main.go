package main

import (
	"campusconnect-api/configs"
	"campusconnect-api/internal/infra/data/pgclient"
	ws "campusconnect-api/internal/interf/webservice"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// lendo as configurações
	cfg, err := configs.LoadConfigs("configs/app.yaml")
	if err != nil {
		panic(err)
	}
	// criando configuração para conexão com o postgres
	pgConfig := &pgclient.Config{
		Host:     cfg.DBHost,
		Port:     cfg.DBPort,
		User:     cfg.DBUser,
		Password: cfg.DBPassword,
		Name:     cfg.DBName,
		Driver:   cfg.DBDriver,
	}
	db := &pgclient.Conn{}
	db.InitConn(pgConfig)
	// criando config do servidor
	router := mux.NewRouter()
	// adicionando meus endpoints e dados no contexto do router
	ws.Routes(router, db.DBConn)
	// criando servidor
	port := "18080"
	// path := "localhost"
	path := "0.0.0.0"
	url := fmt.Sprintf("%s:%s", path, port)
	// iniciando servidor
	log.Printf("Servidor inicializado  http://%s...\n", url)
	err = http.ListenAndServe(url, router)
	if err != nil {
		log.Fatal(err)
	}
}
