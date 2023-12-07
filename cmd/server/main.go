package main

import (
	"campusconnect-api/configs"
	"campusconnect-api/internal/infra/data/pgclient"
	ws "campusconnect-api/internal/interf/webservice"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// lendo as configurações
	cfg, err := configs.LoadConfigs("./cmd/server/app.yaml")
	if err != nil {
		panic(err)
	}
	// criando contexto
	ctx := context.Background()
	// criando configuração para conexão com o postgres
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
	// criando config do servidor
	router := mux.NewRouter()
	// Middleware para adicionar a conexão do banco de dados ao contexto
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// adicionando conexão do banco ao contexto
			ctx = context.WithValue(r.Context(), "dbConn", pgConn.DBConn)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	})
	// adicionando meus endpoints
	ws.Routes(router)
	// criando servidor
	port := "18080"
	path := "localhost"
	url := fmt.Sprintf("%s:%s", path, port)
	// iniciando servidor
	log.Printf("Servidor inicializado  http://%s...\n", url)
	err = http.ListenAndServe(url, router)
	if err != nil {
		log.Fatal(err)
	}
}
