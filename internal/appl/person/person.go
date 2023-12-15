package personappl

import (
	"campusconnect-api/internal/domain/person"
	"campusconnect-api/internal/infra/personrep"
	"campusconnect-api/pkg/utils"
	"context"
	"database/sql"
)

// estrutura que possuira os metodos utilizados no interf
type personApplicationImpl struct {
	ctx context.Context
	tx  *sql.Tx
}

// utilizado apenas na criacao de usuario
func (s *personApplicationImpl) Create(e *person.Entity) error {
	// dto
	ent := &person.Entity{
		ID:        utils.NewIdentity(),
		UserID:    e.UserID,
		CourseID:  e.CourseID,
		FirstName: e.FirstName,
		LastName:  e.LastName,
	}
	// importando repository
	rep := personrep.NewPersonRepository(s.ctx, s.tx)
	err := rep.Store(ent)
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
