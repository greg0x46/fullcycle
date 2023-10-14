package database

import (
	"database/sql"
	"github.com/google/uuid"
)

type Animal struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	ClassID     string
}

func NewAnimal(db *sql.DB) *Animal {
	return &Animal{db: db}
}

func (a *Animal) Create(name, description, classId string) (Animal, error) {
	id := uuid.New().String()
	query := `INSERT INTO animals (id, name, description, class_id) VALUES ($1, $2, $3, $4)`
	_, err := a.db.Exec(query, id, name, description, classId)
	if err != nil {
		return Animal{}, err
	}

	return Animal{ID: id, Name: name, Description: description, ClassID: classId}, nil
}

func (a *Animal) FindAll() ([]Animal, error) {
	rows, err := a.db.Query("SELECT id, name, description, class_id FROM animals")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	animals := []Animal{}
	for rows.Next() {
		var id, name, description, classId string
		if rows.Scan(&id, &name, &description, &classId); err != nil {
			return nil, err
		}

		animals = append(animals, Animal{ID: id, Name: name, Description: description, ClassID: classId})
	}

	return animals, nil
}
