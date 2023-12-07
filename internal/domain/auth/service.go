package auth

// método utilizado no endpoint do lado cliente
type Service interface {
	Create(e *Entity) (string, error)
	FindByEmail(email string) (*Entity, error)
}
