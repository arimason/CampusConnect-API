package resource

import (
	authappl "campusconnect-api/internal/appl/auth"
	"campusconnect-api/internal/domain/auth"
	"campusconnect-api/internal/domain/person"
	"encoding/json"
	"fmt"
	"net/http"

	validator "github.com/go-playground/validator/v10"
)

// ================================================================================
//	Cria um novo usuário
// ================================================================================

// json esperado no corpo da requisição que a API irá receber
type createAuthReq struct {
	Name       string `json:"name" validate:"required"`           // nome realizado para fazer login
	Email      string `json:"email" validate:"required,email"`    // email realizado para fazer login
	Password   string `json:"password" validate:"required,min=8"` // senha deve conter pelo menos 8 caracteres
	Permission string `json:"permission" validate:"required"`     // permission deve ser um desses valores: 'student', 'teacher', 'admin', 'owner'
	CourseID   string `json:"courseID" validate:"required"`
	FirstName  string `json:"firstName" validate:"required"`
	LastName   string `json:"lastName" validate:"required"`
}

// json retornado no corpo da resposta
type createAuthResp struct {
}

// realiza o decode da requisição recebida pela API
func decodeCreateAuth(r *http.Request) (*createAuthReq, error) {
	var dto *createAuthReq
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		return nil, err
	}
	return dto, nil
}

// Create user godoc
// @Summary Create User
// @Description Create User
// @Tags auth
// @Accept json
// @Produce json
// @Param request body createAuthReq true "User request created"
// @Success 201
// @Response 201 {object} createAuthResp "User created successfully"
// @Failure 400 {object} errorResp "Bad Request"
// @Failure 401 {object} errorResp "Unauthorized"
// @Failure 403 {object} errorResp "Forbidden"
// @Failure 404 {object} errorResp "Not Found"
// @Failure 500 {object} errorResp "Internal Server Error"
// @Router /pub/user [post]
// utilizo as regras de negócio do appl e preparo o response de acordo
func CreateAuthHandler(w http.ResponseWriter, r *http.Request) {
	// caso tivesse um id no formato uuid no path: Param id path string true "product ID" Format(uuid)
	defer r.Body.Close()
	req, err := decodeCreateAuth(r)
	if err != nil {
		responseError(w, http.StatusBadRequest, ErrDecodeReqBody, err.Error())
		return
	}
	// Validando os campos usando o pacote validator
	validate := validator.New()
	err = validate.Struct(req)
	fmt.Println("nil", err)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// importandos o serviço
	appl := authappl.NewAuthApplication(r.Context())
	// criando entidade
	authEnt := &auth.Entity{
		Name:       req.Name,
		Email:      req.Email,
		Password:   req.Password,
		Permission: auth.Permission[req.Permission],
	}
	personEnt := person.Entity{
		CourseID:  req.CourseID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}
	// envio os dados de criação do usuário para o appl
	err = appl.Create(authEnt, &personEnt)
	if err != nil {
		responseError(w, http.StatusBadRequest, ErrCreateEnt, err.Error())
		return
	}
	// status 201
	w.WriteHeader(http.StatusCreated)
}

// ================================================================================
//	Obtem dados de um usuário - necessário token
// ================================================================================

// json retornado no corpo da resposta
type findByEmailResp struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Permission string `json:"permission"`
}

// realizo a consulta no appl para retornar os dados da requisição
func FindByEmailHandler(w http.ResponseWriter, r *http.Request) {
	appl := authappl.NewAuthApplication(r.Context())
	ent, err := appl.FindByEmail()
	if err != nil {
		responseError(w, http.StatusBadRequest, ErrFind, err.Error())
		return
	}
	resp := &findByEmailResp{
		ID:         string(ent.ID),
		Name:       ent.Name,
		Email:      ent.Email,
		Permission: string(ent.Permission),
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&resp)
	if err != nil {
		responseError(w, http.StatusInternalServerError, ErrFind, err.Error())
		return
	}
}

// ================================================================================
//	Realiza login
// ================================================================================

type loginReq struct {
	EmailOrName string `json:"emailOrName"`
	Password    string `json:"password"`
}

type loginResp struct {
	Token string `json:"token"`
}

// decodifica os dados do body da requisição
func decodeLoginReq(r *http.Request) (*loginReq, error) {
	dto := &loginReq{}
	err := json.NewDecoder(r.Body).Decode(dto)
	if err != nil {
		return nil, err
	}
	return dto, nil
}

// Create user godoc
// @Summary Login
// @Description Request for Login
// @Tags auth
// @Accept json
// @Produce json
// @Param request body loginReq true "User request login"
// @Success 200
// @Response 200 {object} loginResp "User successfully logged in"
// @Failure 400 {object} errorResp "Bad Request"
// @Failure 401 {object} errorResp "Unauthorized"
// @Failure 403 {object} errorResp "Forbidden"
// @Failure 404 {object} errorResp "Not Found"
// @Failure 500 {object} errorResp "Internal Server Error"
// @Router /pub/user/login [post]
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// no fim de tudo fecha o corpo do request para evitar vazamento de recurso
	defer r.Body.Close()
	// decodificação da requisição
	req, err := decodeLoginReq(r)
	if err != nil {
		responseError(w, http.StatusBadRequest, "erro ao decodificar o body da requisição", err.Error())
		return
	}
	// importo o appl e passo contexto
	appl := authappl.NewAuthApplication(r.Context())
	token, err := appl.Login(req.EmailOrName, req.Password)
	if err != nil {
		responseError(w, http.StatusBadRequest, "erro ao realizar login", err.Error())
		return
	}
	// retorno o token para o front
	resp := &loginResp{
		Token: token,
	}
	// definindo o tipo do conteúdo
	w.Header().Set("Content-Type", "application/json")
	// status code
	w.WriteHeader(http.StatusOK)
	// realizo o encode para o ResponseWriter
	err = json.NewEncoder(w).Encode(&resp)
	if err != nil {
		responseError(w, http.StatusInternalServerError, "erro ao realizar o encode para o ResponseWriter", err.Error())
	}
}
