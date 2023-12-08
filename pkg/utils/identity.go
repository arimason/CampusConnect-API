package utils

import "github.com/google/uuid"

// Identity é o tipo de todo identificador de entidades
type Identity string

// Gero um novo valor apara Identity
func NewIdentity() Identity {
	return Identity(uuid.New().String())
}
