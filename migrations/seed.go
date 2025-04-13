package migrations

import (
	"github.com/irfhakeem/go-fiber-clean-starter/migrations/seeds"
	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) error {
	if err := seeds.UserSeeder(db); err != nil {
		return err
	}

	return nil
}
