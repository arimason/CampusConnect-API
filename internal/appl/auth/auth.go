package authappl

import (
	"campusconnect-api/configs"
	"campusconnect-api/internal/domain/auth"
	authrep "campusconnect-api/internal/infra/auth"
	contxt "campusconnect-api/internal/infra/context"
	"campusconnect-api/pkg/utils"
	"context"
	"errors"
	"fmt"
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
	_, err = s.FindByEmail(e.Email)
	if err == authrep.ErrFindByEmailNotFound {
		return "", errors.New("email já existe")
	}
	if err != nil {
		return "", err
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
	token, err := createToken(string(ent.ID), cfg.JWTSecret)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *authApplicationImpl) FindByEmail(email string) (*auth.Entity, error) {
	// obtendo transação
	tx, err := contxt.GetDbConn(s.ctx)
	if err != nil {
		return nil, err
	}
	rep := authrep.NewAuthRepository(tx)
	ent, err := rep.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	return ent, nil
}

func (s *authApplicationImpl) UpdatePassword(password string) (string, error) {
	cfg, err := configs.LoadConfigs("./configs/app.yaml")
	if err != nil {
		return "", errors.New("erro ao obter configs")
	}
	tk, err := verifierToken(s.ctx, cfg.JWTSecret)
	if err != nil {
		return "", err
	}
	fmt.Println(">>", tk)
	return "", nil
}

func NewAuthApplication(ctx context.Context) auth.Service {
	return &authApplicationImpl{
		ctx: ctx,
	}
}
