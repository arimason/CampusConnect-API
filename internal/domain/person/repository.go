package person

type Repository interface {
	// Salvar no banco
	Store(e *Entity) error
}
