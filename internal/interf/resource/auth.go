package resource

import (
	authappl "campusconnect-api/internal/appl/auth"
	"campusconnect-api/internal/domain/auth"
	"encoding/json"
	"net/http"
)

// ================================================================================
//	Cria um novo usuário
// ================================================================================

type createAuthReq struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Permission string `json:"permision"`
}

type createAuthResp struct {
	JWT string `json:"token"`
}

func decodeCreateAuth(r *http.Request) (*createAuthReq, error) {
	var dto *createAuthReq
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		return nil, err
	}
	return dto, nil
}

func CreateAuthHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	req, err := decodeCreateAuth(r)
	if err != nil {
		responseError(w, http.StatusBadRequest, ErrDecodeReqBody, err.Error())
		return
	}
	// importandos o serviço
	appl := authappl.NewAuthApplication(r.Context())
	// criando entidade
	ent := &auth.Entity{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	token, err := appl.Create(ent)
	if err != nil {
		responseError(w, http.StatusBadRequest, ErrCreateEnt, err.Error())
		return
	}
	resp := &createAuthResp{
		JWT: token,
	}
	// definindo tipo do conteúdo
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&resp)
	if err != nil {
		responseError(w, http.StatusInternalServerError, "Erro ao criar json para resposta", err.Error())
		return
	}
}

// ================================================================================
//	Obtem dados de um usuário - necessário token
// ================================================================================

type findByEmailResp struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func FindByEmailHandler(w http.ResponseWriter, r *http.Request) {
	appl := authappl.NewAuthApplication(r.Context())
	ent, err := appl.FindByEmail()
	if err != nil {
		responseError(w, http.StatusBadRequest, ErrFind, err.Error())
		return
	}
	resp := &findByEmailResp{
		ID:    string(ent.ID),
		Name:  ent.Name,
		Email: ent.Email,
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&resp)
	if err != nil {
		responseError(w, http.StatusInternalServerError, ErrFind, err.Error())
		return
	}
}
