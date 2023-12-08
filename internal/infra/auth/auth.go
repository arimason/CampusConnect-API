package authrep

import (
	"campusconnect-api/internal/domain/auth"
	"campusconnect-api/pkg/utils"
	"database/sql"
	"fmt"
)

type authRepositoryImpl struct {
	Tx *sql.Tx
}

func (r *authRepositoryImpl) scan(row *sql.Row) (*auth.Entity, error) {
	id := sql.NullString{}
	name := sql.NullString{}
	email := sql.NullString{}
	password := sql.NullString{}
	err := row.Scan(
		&id,
		&name,
		&email,
		&password,
	)
	if err != nil {
		return nil, err

	}
	ent := new(auth.Entity)
	if id.Valid {
		ent.ID = utils.Identity(id.String)
	}
	if name.Valid {
		ent.Name = name.String
	}
	if email.Valid {
		ent.Email = email.String
	}
	if password.Valid {
		ent.Password = password.String
	}
	return ent, nil
}

func (r *authRepositoryImpl) create(e *auth.Entity) error {
	// prepara o sql e o valores para serem utilizados na instrução do banco
	sqlStmt := `
	insert into tb_user(
		id,
		name,
		email,
		password
	) values ($1, $2, $3, $4)
	`
	_, err := r.Tx.Exec(sqlStmt, e.ID, e.Name, e.Email, e.Password)
	if err != nil {
		return fmt.Errorf("falha ao executar SQL: %w", err)
	}
	// sucesso
	return nil
}

func (r *authRepositoryImpl) findByEmail(email string) (*auth.Entity, error) {
	// sql para consulta no banco de dados
	sqlStmt := `
	select
		id,
		name,
		email,
		password
	from tb_user
	where email = $1
	`
	// realizando consulta
	row := r.Tx.QueryRow(sqlStmt, email)
	// atribuindo os valores obtidos do banco de dados para a minha entidade
	ent, err := r.scan(row)
	// entidade não encontrada
	if err == sql.ErrNoRows {
		return nil, ErrFindByEmailNotFound
	}
	if err != nil {
		return nil, err
	}
	// sucesso
	return ent, nil
}

func (r *authRepositoryImpl) Create(e *auth.Entity) error {
	err := r.create(e)
	if err != nil {
		r.Tx.Rollback()
		return err
	}
	return nil
}

func (r *authRepositoryImpl) FindByEmail(email string) (*auth.Entity, error) {
	ent, err := r.findByEmail(email)
	if err != nil && err != ErrFindByEmailNotFound {
		r.Tx.Rollback()
		return nil, err
	}
	return ent, nil
}

func NewAuthRepository(tx *sql.Tx) auth.Repository {
	return &authRepositoryImpl{
		Tx: tx,
	}
}
