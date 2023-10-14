package graph

import "github.com/greg0x46/fc2-graphql/database"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	AnimalDB *database.Animal
	ClassDB *database.Class
}
