package ws

import (
	"campusconnect-api/internal/interf/resource"
	"net/http"

	"github.com/gorilla/mux"
)

func Routes(r *mux.Router) {
	r.HandleFunc("/user", resource.CreateAuthHandler).Methods(http.MethodPost)
	r.HandleFunc("/user/{email}", resource.FindByEmailHandler).Methods(http.MethodGet)
}
