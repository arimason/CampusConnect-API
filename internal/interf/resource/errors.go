package resource

import (
	"encoding/json"
	"net/http"
)

const (
	ErrDecodeReqBody = "Erro ao decodificar o corpo da requisição"
	ErrCreateEnt     = "Erro ao criar entidade"
	ErrFind          = "Erro ao buscar entidade"
)

type errorResp struct {
	Message string `json:"message"` // referente a mensagem tratada do erro
	Error   string `json:"error"`   // refere diretamente ao erro
}

func responseError(w http.ResponseWriter, status int, message, errs string) {
	errResp := &errorResp{
		Message: message,
		Error:   errs,
	}
	// definindo content-type
	w.Header().Set("Content-Type", "application/json")
	// definindo status da requisição
	w.WriteHeader(status)
	// codificando para json
	err := json.NewEncoder(w).Encode(&errResp)
	if err != nil {
		http.Error(w, "Erro interno", http.StatusInternalServerError)
		return
	}
}
