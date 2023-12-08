package ws

import (
	"campusconnect-api/internal/interf/resource"
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type wsImpl struct {
	ctx context.Context
	db  *sql.DB
}

func (ws *wsImpl) addToContext(ctx context.Context, keyValues map[string]interface{}) context.Context {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	for key, value := range keyValues {
		ctx = context.WithValue(ctx, key, value)
	}
	return ctx
}

func (ws *wsImpl) prepareHttpWithContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// criando contexto e atribuindo valores
		ctx := ws.addToContext(ws.ctx, map[string]interface{}{
			"JWT":    r.Header.Get("Authorization"),
			"dbConn": ws.db,
		})
		// iniciando novo router agora com o contexto
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func Routes(r *mux.Router, db *sql.DB) {
	// middleware para adicionar a conexão do banco de dados e ouros valores ao contexto
	ws := &wsImpl{
		db:  db,
		ctx: context.Background(),
	}
	r.Use(ws.prepareHttpWithContext)
	// criando endpoints
	// criar usuário
	r.HandleFunc("/user", resource.CreateAuthHandler).Methods(http.MethodPost)
	// recupera dados de usuário
	r.HandleFunc("/user", resource.FindByEmailHandler).Methods(http.MethodGet)
	// realiza login
	r.HandleFunc("/user/login", resource.LoginHandler).Methods(http.MethodPost)
}
