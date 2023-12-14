package auth

import "campusconnect-api/internal/domain/person"

// método utilizado no endpoint do lado cliente
type Service interface {
	// Cria usuário
	Create(e *Entity, p *person.Entity) error
	// Realiza login
	Login(emailOrName, password string) (string, error)
	// Obtém-se os dados de usuário de acordo com o email
	FindByEmail() (*Entity, error)
}
