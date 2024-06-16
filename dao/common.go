package dao

import (
	"goweb/dto"

	"gorm.io/gorm"
)

// Paginate is a function that creates a GORM scope for pagination.
// It takes a Paginate struct as input and returns a function that
// can be used to apply pagination to a GORM query.
func Paginate(p dto.Paginate) func(orm *gorm.DB) *gorm.DB {
	return func(orm *gorm.DB) *gorm.DB {
		// Calculate the offset based on the current page and limit.
		offset := (p.GetPage() - 1) * p.GetLimit()

		// Apply the offset and limit to the GORM query.
		return orm.Offset(offset).Limit(p.GetLimit())
	}
}
