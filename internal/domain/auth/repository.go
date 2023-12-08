package auth

// métodos que iram realizar requisições diretamente banco de dados
type Repository interface {
	// insere um novo usuário no banco
	Create(e *Entity) error
	// realiza uma select em tb_user de acordo com o campo email
	FindByEmailOrName(emailOrName string) (*Entity, error)
}
