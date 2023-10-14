package database

import (
	"database/sql"
	"github.com/google/uuid"
)

type Class struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewClass(db *sql.DB) *Class {
	return &Class{db: db}
}

func (t *Class) Create(name string, description string) (Class, error) {
	id := uuid.New().String()
	_, err := t.db.Exec("INSERT INTO classes (id, name, description) VALUES ($1, $2, $3)", id, name, description)

	if err != nil {
		return Class{}, err
	}

	return Class{ID: id, Name: name, Description: description}, nil
}

func (c *Class) FindAll() ([]Class, error) {
	rows, err := c.db.Query("SELECT id, name, description FROM classes")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	classes := []Class{}
	for rows.Next() {
		var id, name, description string
		if rows.Scan(&id, &name, &description); err != nil {
			return nil, err
		}

		classes = append(classes, Class{ID: id, Name: name, Description: description})
	}

	return classes, nil
}

func (c *Class) FindByAnimalId(animalID string) (Class, error) {
	var id, name, description string
	err := c.db.QueryRow("SELECT c.id, c.name, c.description FROM classes c JOIN animals a ON c.id = a.class_id WHERE a.id = $1", animalID).
	Scan(&id, &name, &description)

	if err != nil {
		return Class{}, err
	}

	return Class{ID: id, Name: name, Description: description}, nil
}
