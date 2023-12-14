package person

import "campusconnect-api/pkg/utils"

// Struct referente a pessoa do usuario
type Entity struct {
	ID        utils.Identity
	UserID    string
	CourseID  string
	FirstName string
	LastName  string
}
