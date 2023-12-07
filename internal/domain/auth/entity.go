package auth

import "campusconnect-api/pkg/utils"

type Entity struct {
	ID       utils.Identity
	Name     string
	Email    string
	Password string
}
