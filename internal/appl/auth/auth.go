package authappl

import (
	"campusconnect-api/internal/domain/auth"
	authrep "campusconnect-api/internal/infra/auth"
	contxt "campusconnect-api/internal/infra/context"
	"campusconnect-api/pkg/utils"
	"context"

	"golang.org/x/crypto/bcrypt"
)

type authApplicationImpl struct {
	ctx context.Context
}

func (s *authApplicationImpl) validatePassword(hashPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}

func (s *authApplicationImpl) Create(e *auth.Entity) (string, error) {
	// iniciando transação
	tx, err := contxt.GetDbConn(s.ctx)
	if err != nil {
		return "", err
	}
	// importando métodos do repositório
	rep := authrep.NewAuthRepository(tx)
	// geração de hash a partir da senha
	hash, err := bcrypt.GenerateFromPassword([]byte(e.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	ent := &auth.Entity{
		ID:       utils.NewIdentity(),
		Name:     e.Name,
		Email:    e.Email,
		Password: string(hash),
	}
	// inserindo entidade no data base
	err = rep.Create(ent)
	if err != nil {
		return "", err
	}
	return string(ent.ID), nil
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

func NewAuthApplication(ctx context.Context) auth.Service {
	return &authApplicationImpl{
		ctx: ctx,
	}
}
