package auth

import "campusconnect-api/pkg/utils"

// criação do tipo Kind
type Kind string

// criação constantes do tipo Kind atribuindo valores de permissão
const (
	Admin   Kind = "admin"
	Teacher Kind = "professor"
	Student Kind = "aluno"
)

// criação mapa Permission para obter o dado do tipo Kind de acordo com uma string
var Permission map[string]Kind = map[string]Kind{
	"admin":     Admin,
	"professor": Teacher,
	"aluno":     Student,
}

type Entity struct {
	ID       utils.Identity
	Name     string
	Email    string
	Password string
	// ao utilizar Kind para definir o tipo de Permission, eu crio o controle dos valores aceitos em Permission
	Permission Kind // pode ser admin, professor ou aluno, isso irá definir as permissões de usuário
}
