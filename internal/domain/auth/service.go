package auth

// método utilizado no endpoint do lado cliente
type Service interface {
	// Cria usuário
	Create(e *Entity) (string, error)
	// Realiza login
	Login(emailOrName, password string) (string, error)
	// Obtém-se os dados de usuário de acordo com o email
	FindByEmail() (*Entity, error)
}
