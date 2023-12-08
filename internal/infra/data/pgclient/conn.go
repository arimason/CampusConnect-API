package pgclient

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"

	_ "github.com/lib/pq"
)

// estrutura com conexão com o banco
type Conn struct {
	DBConn *sql.DB
}

// configuraçẽs para realizar comunicações com o banco
type Config struct {
	Driver   string
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	Schema   string
}

// inicializo a comunicação com o banco de dados
func (c *Conn) InitConn(cfg *Config) {
	// // criando os parâmetros para conexão com o postgres
	pgurl := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(cfg.User, cfg.Password),
		Host:   fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Path:   fmt.Sprintf("/%s", cfg.Name),
	}
	// sslmode disable
	q := pgurl.Query()
	q.Add("sslmode", "disable")
	pgurl.RawQuery = q.Encode()
	// iniciando conexão com o postgres
	db, err := sql.Open(cfg.Driver, pgurl.String())
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close() fechar conexão apenas no final da requisição
	// realizando teste da conexão
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	c.DBConn = db
	// sucess
	log.Println("Conexão com o PostgreSQL estabelecidade com sucesso!")
}
