/*
Package dao provides data access objects (DAOs) for interacting with the database.

DAOs are responsible for database operations, making it easier
to maintain and modify the code.
*/

package dao

import (
	"goweb/global"

	"gorm.io/gorm"
)

// BaseDao provides a base structure for DAO implementations.
type BaseDao struct {
	// Orm is the GORM database connection.
	Orm *gorm.DB
}

// NewBase creates a new instance of BaseDao with the global DB connection.
func NewBase() BaseDao {
	return BaseDao{
		Orm: global.DB, // Use the global DB connection from the global package.
	}
}
