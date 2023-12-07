package pgclient

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"

	_ "github.com/lib/pq"
)

type Conn struct {
	DBConn *sql.DB
}

type Config struct {
	Driver   string
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

func (c *Conn) InitConn(cfg *Config) {
	// criando os parâmetros para conexão com o postgres
	dbURL := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(cfg.User, cfg.Password),
		Host:   fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Path:   cfg.Name,
	}
	// sslmode disable
	q := dbURL.Query()
	q.Add("sslmode", "disable")
	dbURL.RawQuery = q.Encode()
	// iniciando conexão com o postgres
	conn := dbURL.String()
	db, err := sql.Open(cfg.Driver, conn)
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
