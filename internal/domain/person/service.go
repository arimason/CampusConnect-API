package person

type Service interface {
	// Criacao de pessoa
	Create(e *Entity) error
}
