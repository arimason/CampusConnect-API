package auth

// métodos que iram realizar requisições diretamente banco de dados
type Repository interface {
	Create(e *Entity) error
	FindByEmail(email string) (*Entity, error)
}
