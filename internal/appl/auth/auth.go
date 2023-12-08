package authappl

import (
	"campusconnect-api/configs"
	"campusconnect-api/internal/domain/auth"
	authrep "campusconnect-api/internal/infra/auth"
	contxt "campusconnect-api/internal/infra/context"
	"campusconnect-api/pkg/utils"
	"context"
	"errors"
)

type authApplicationImpl struct {
	ctx context.Context
}

func (s *authApplicationImpl) Create(e *auth.Entity) (string, error) {
	// obtendo configuração
	cfg, err := configs.LoadConfigs("./configs/app.yaml")
	if err != nil {
		return "", errors.New("erro ao obter configs")
	}
	// iniciando transação
	tx, err := contxt.GetDbConn(s.ctx)
	if err != nil {
		return "", err
	}
	// importando métodos do repositório
	rep := authrep.NewAuthRepository(tx)
	// geração de hash a partir da senha
	hash, err := utils.GenerateHash(e.Password)
	if err != nil {
		return "", err
	}
	findEnt, err := rep.FindByEmail(e.Email)
	if err != nil && err != authrep.ErrFindByEmailNotFound {
		return "", err
	}
	if findEnt != nil {
		return "", errors.New("email já existe")
	}
	ent := &auth.Entity{
		ID:       utils.NewIdentity(),
		Name:     e.Name,
		Email:    e.Email,
		Password: hash,
	}
	// inserindo entidade no data base
	err = rep.Create(ent)
	if err != nil {
		return "", err
	}
	// gerando token
	token, err := createToken(string(ent.ID), e.Email, cfg.JWTSecret)
	if err != nil {
		return "", err
	}
	return token, nil
}

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
	ent, err := rep.FindByEmail(tk.Email)
	if err != nil {
		return nil, err
	}
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
