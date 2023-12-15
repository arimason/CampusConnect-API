package personrep

import (
	"campusconnect-api/internal/domain/person"
	"campusconnect-api/pkg/utils"
	"fmt"

	"context"
	"database/sql"
)

type personRepositoryImpl struct {
	Tx *sql.Tx
}

func (r *personRepositoryImpl) scan(rows *sql.Rows) (*person.Entity, error) {
	id := sql.NullString{}
	userID := sql.NullString{}
	courseID := sql.NullString{}
	firstName := sql.NullString{}
	lastName := sql.NullString{}
	//scan
	err := rows.Scan(
		&id,
		&userID,
		&courseID,
		&firstName,
		&lastName,
	)
	if err != nil {
		return nil, err
	}
	ent := &person.Entity{}
	if id.Valid {
		ent.ID = utils.Identity(id.String)
	}
	if userID.Valid {
		ent.UserID = userID.String
	}
	if courseID.Valid {
		ent.CourseID = courseID.String
	}
	if firstName.Valid {
		ent.FirstName = firstName.String
	}
	if lastName.Valid {
		ent.LastName = lastName.String
	}
	return ent, nil
}

func (r *personRepositoryImpl) Store(e *person.Entity) error {
	sqlStmt := `
		INSERT INTO tb_person (
			id,
			user_id,
			first_name,
			last_name,
			course_id
		) VALUES ($1, $2, $3, $4, $5)
	`
	_, err := r.Tx.Exec(sqlStmt, e.ID, e.UserID, e.FirstName, e.LastName, e.CourseID)
	fmt.Println(e.UserID, err)
	if err != nil {
		return fmt.Errorf("falha ao executar SQL: %w", err)
	}
	// sucesso
	return nil
}

func NewPersonRepository(ctx context.Context, tx *sql.Tx) person.Repository {
	return &personRepositoryImpl{
		Tx: tx,
	}
}
