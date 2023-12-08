package auth

import "campusconnect-api/pkg/utils"

type Kind string

const (
	Admin   Kind = "admin"
	Teacher Kind = "professor"
	Student Kind = "aluno"
)

var Permission map[string]Kind = map[string]Kind{
	"admin":     Admin,
	"professor": Teacher,
	"aluno":     Student,
}

type Entity struct {
	ID         utils.Identity
	Name       string
	Email      string
	Password   string
	Permission Kind
}
