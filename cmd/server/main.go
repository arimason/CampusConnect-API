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
			// criando contexto e adicionando conexão do banco ao contexto
			ctx := context.WithValue(r.Context(), "dbConn", pgConn.DBConn)
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

// func addToContext(ctx context.Context, keyValues map[string]interface{}) context.Context {
//     for key, value := range keyValues {
//         ctx = context.WithValue(ctx, key, value)
//     }
//     return ctx
// }

// // Uso da função addToContext no seu middleware
// router.Use(func(next http.Handler) http.Handler {
//     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//         // Extrai o token do cabeçalho da solicitação
//         token := r.Header.Get("Authorization")

//         // Cria um contexto com cancelamento e timeout, adiciona token e conexão do banco
//         ctx := context.WithTimeout(r.Context(), 10*time.Second)
//         ctx = addToContext(ctx, map[string]interface{}{
//             "token":  token,
//             "dbConn": pgConn.DBConn,
//             // Adicione outros valores conforme necessário
//         })

//         // Defer para garantir que a conexão seja fechada no final do manipulador
//         defer func() {
//             if err := pgConn.DBConn.Close(); err != nil {
//                 // Lidar com o erro de fechamento, se necessário
//                 log.Println("Erro ao fechar a conexão com o banco de dados:", err)
//             }
//         }()

//         // Chama o próximo manipulador com o novo contexto
//         next.ServeHTTP(w, r.WithContext(ctx))
//     })
// })
