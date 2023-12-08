package authappl

import (
	"campusconnect-api/internal/domain/auth"
	authrep "campusconnect-api/internal/infra/auth"
	contxt "campusconnect-api/internal/infra/context"
	"campusconnect-api/pkg/utils"
	"context"
	"errors"
)

// estrutura me garante que eu terei simetria com os métodos necessário e expressado na interface Service referente a entidade auth
type authApplicationImpl struct {
	ctx context.Context
}

// Cria um novo usuário e retorna o token gerado após a criação desse novo usuário, não sendo necessário o cliente passsar por um login
func (s *authApplicationImpl) Create(e *auth.Entity) (string, error) {
	// iniciando transação
	tx, err := contxt.GetDbConn(s.ctx)
	if err != nil {
		return "", err
	}
	// importando métodos do repositório
	rep := authrep.NewAuthRepository(tx)
	// verificando se o email já existe, de acordo com o dados do banco
	findEnt, err := rep.FindByEmailOrName(e.Email)
	if err != nil && err != authrep.ErrFindByEmailNotFound {
		return "", err
	}
	if findEnt != nil {
		return "", errors.New("email já existe")
	}
	// geração de hash a partir da senha
	hash, err := utils.GenerateHash(e.Password)
	if err != nil {
		return "", err
	}
	// atribuindo valor a entidade
	ent := &auth.Entity{
		ID:         utils.NewIdentity(),
		Name:       e.Name,
		Email:      e.Email,
		Password:   hash,
		Permission: e.Permission,
	}
	// inserindo entidade no data base
	err = rep.Create(ent)
	if err != nil {
		return "", err
	}
	// gerando token
	token, err := createToken(string(ent.ID), e.Email, string(e.Permission))
	if err != nil {
		return "", err
	}
	// commit da transação
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return "", err
	}
	// sucesso
	return token, nil
}

// Realizo o login, recebendo email, ou username e uma senha, dados esses que serão validados
func (s *authApplicationImpl) Login(emailOrName, password string) (string, error) {
	// obtendo transação
	tx, err := contxt.GetDbConn(s.ctx)
	if err != nil {
		return "", err
	}
	// importando repositorio
	rep := authrep.NewAuthRepository(tx)
	// realizo consulta no banco para comparar se os dados fornecidos para login são válidos
	ent, err := rep.FindByEmailOrName(emailOrName)
	// caso a consulta retorne vazio, ou seja, não possui esse usuário no banco, retorna um erro
	if err != nil {
		return "", err
	}
	// verifico se a senha é válida
	err = utils.ValidateHash(ent.Password, password)
	if err != nil {
		return "", errors.New("senha inválida")
	}
	// geração do token após a confrimação dos dados
	token, err := createToken(string(ent.ID), ent.Email, string(ent.Permission))
	if err != nil {
		return "", err
	}
	// sucesso
	return token, nil
}

// Busca os dados do usuário de acordo com o email dentro do token
func (s *authApplicationImpl) FindByEmail() (*auth.Entity, error) {
	// obtendo transação
	tx, err := contxt.GetDbConn(s.ctx)
	if err != nil {
		return nil, err
	}
	// validando token
	tk, err := tokenValues(s.ctx)
	if err != nil {
		return nil, err
	}
	// importando repositorio
	rep := authrep.NewAuthRepository(tx)
	// realizando consulta no banco de dados
	ent, err := rep.FindByEmailOrName(tk.Email)
	if err != nil {
		return nil, err
	}
	// commit da transação
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	// sucesso
	return ent, nil
}

func (s *authApplicationImpl) UpdatePassword(password string) (string, error) {
	// validando token
	_, err := tokenValues(s.ctx)
	if err != nil {
		return "", err
	}
	return "", nil
}

func NewAuthApplication(ctx context.Context) auth.Service {
	return &authApplicationImpl{
		ctx: ctx,
	}
}
