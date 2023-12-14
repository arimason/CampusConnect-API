package personappl

import (
	"campusconnect-api/internal/domain/person"
	contxt "campusconnect-api/internal/infra/context"
	"campusconnect-api/internal/infra/personrep"
	"campusconnect-api/pkg/utils"
	"context"
)

// estrutura que possuira os metodos utilizados no interf
type personApplicationImpl struct {
	ctx context.Context
}

// utilizado apenas na criacao de usuario
func (s *personApplicationImpl) Create(e *person.Entity) error {
	// iniciando transação
	tx, err := contxt.GetDbConn(s.ctx)
	if err != nil {
		return err
	}
	// dto
	ent := &person.Entity{
		ID:        utils.NewIdentity(),
		UserID:    e.UserID,
		CourseID:  e.CourseID,
		FirstName: e.FirstName,
		LastName:  e.LastName,
	}
	// importando repository
	rep := personrep.NewPersonRepository(s.ctx, tx)
	err = rep.Store(ent)
	if err != nil {
		return err
	}
	// success
	return nil
}

// funcao que exporta os metodos do meu appl
func NewPersonApplication(ctx context.Context) person.Service {
	return &personApplicationImpl{
		ctx: ctx,
	}
}
