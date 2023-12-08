package contxt

import (
	"context"
	"database/sql"
	"errors"
)

// a partir do contexto eu obtenho a conexão com o banco e gero uma transação
func GetDbConn(ctx context.Context) (*sql.Tx, error) {
	dbConn, ok := ctx.Value("dbConn").(*sql.DB)
	if !ok {
		return nil, errors.New("erro ao pegar conexão do banco de dados no contexto")
	}
	tx, err := dbConn.Begin()
	if err != nil {
		return nil, errors.New("erro ao iniciar transação")
	}
	return tx, nil
}
