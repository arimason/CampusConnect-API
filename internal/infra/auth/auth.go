package authrep

import (
	"campusconnect-api/internal/domain/auth"
	"campusconnect-api/pkg/utils"
	"database/sql"
	"fmt"
)

// esturtura que garante que os métodos que os métodos da interface Repository estão sendo criados
type authRepositoryImpl struct {
	Tx *sql.Tx
}

// atribui os valores de retorno do select à minha entidade
func (r *authRepositoryImpl) scan(row *sql.Row) (*auth.Entity, error) {
	id := sql.NullString{}
	name := sql.NullString{}
	email := sql.NullString{}
	password := sql.NullString{}
	permission := sql.NullString{}
	err := row.Scan(
		&id,
		&name,
		&email,
		&password,
		&permission,
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
	if permission.Valid {
		ent.Permission = auth.Permission[permission.String]
	}
	return ent, nil
}

// prepara e realiza o insert no banco
func (r *authRepositoryImpl) create(e *auth.Entity) error {
	// prepara o sql e o valores para serem utilizados na instrução do banco
	sqlStmt := `
	insert into tb_user(
		id,
		name,
		email,
		password,
		permission
	) values ($1, $2, $3, $4, $5)
	`
	_, err := r.Tx.Exec(sqlStmt, e.ID, e.Name, e.Email, e.Password, e.Permission)
	if err != nil {
		return fmt.Errorf("falha ao executar SQL: %w", err)
	}
	// sucesso
	return nil
}

// prepara o select para consulta com o banco de dados e obtém os resultados
func (r *authRepositoryImpl) findByEmailOrName(emailOrName string) (*auth.Entity, error) {
	// sql para consulta no banco de dados
	sqlStmt := `
	select
		id,
		name,
		email,
		password,
		permission
	from tb_user
	where email = $1 OR name = $1
	`
	// realizando consulta
	row := r.Tx.QueryRow(sqlStmt, emailOrName)
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

// Realiza um INSERT no banco de dados, inserindo o dados de usuário
func (r *authRepositoryImpl) Create(e *auth.Entity) error {
	err := r.create(e)
	if err != nil {
		r.Tx.Rollback()
		return err
	}
	// commit da transação
	err = r.Tx.Commit()
	if err != nil {
		r.Tx.Rollback()
		return err
	}
	return nil
}

// Realiza uma consulta de acordo com o email ou name, retornandod dados do usuário
func (r *authRepositoryImpl) FindByEmailOrName(emailOrName string) (*auth.Entity, error) {
	ent, err := r.findByEmailOrName(emailOrName)
	if err == ErrFindByEmailNotFound {
		return nil, err
	}
	if err != nil {
		r.Tx.Rollback()
		return nil, err
	}

	return ent, nil
}

// função que realiza a possibilidade de exportar todos os métodos aqui, com tanto que estejam na interface Repository
func NewAuthRepository(tx *sql.Tx) auth.Repository {
	return &authRepositoryImpl{
		Tx: tx,
	}
}
