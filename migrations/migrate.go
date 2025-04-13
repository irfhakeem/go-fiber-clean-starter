package migrations

import (
	"log"
	"strings"

	"github.com/irfhakeem/go-fiber-clean-starter/dto"
	"github.com/irfhakeem/go-fiber-clean-starter/entity"
	"gorm.io/gorm"
)

var (
	// Enums
	enums = map[string][]string{
		"user_role": {
			"Admin",
			"Member",
			"Guest",
		},
		"gender": {
			"Male",
			"Female",
			"Prefer Not To Say",
		},
	}

	// Models
	models = []any{
		&entity.User{},
	}
)

func Migrate(db *gorm.DB) error {
	for name, values := range enums {
		quotedValues := make([]string, len(values))
		for i, value := range values {
			quotedValues[i] = "'" + value + "'"
		}
		if err := db.Exec("CREATE TYPE " + name + " AS ENUM (" + strings.Join(quotedValues, ", ") + ")").Error; err != nil {
			log.Print(err)
			return dto.ErrCreateEnum
		}
	}

	if err := db.AutoMigrate(
		models...,
	); err != nil {
		return err
	}

	return nil
}

func Fresh(db *gorm.DB) error {
	for name := range enums {
		if err := db.Exec("DROP TYPE IF EXISTS " + name).Error; err != nil {
			log.Print(err)
			return dto.ErrDropEnum
		}
	}

	if err := db.Migrator().DropTable(
		models...,
	); err != nil {
		return err
	}

	if err := Migrate(db); err != nil {
		return err
	}

	return nil
}
