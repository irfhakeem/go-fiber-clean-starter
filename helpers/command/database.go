package command

import (
	"log"
	"os"

	"github.com/irfhakeem/go-fiber-clean-starter/migrations"
	"gorm.io/gorm"
)

func DatabaseCommand(db *gorm.DB) bool {
	migrate := false
	seed := false
	fresh := false

	for _, arg := range os.Args {
		if arg == "--migrate" {
			migrate = true
		}
		if arg == "--seed" {
			seed = true
		}
		if arg == "--fresh" {
			fresh = true
		}
	}

	if migrate {
		if err := migrations.Migrate(db); err != nil {
			log.Println(err)
			return false
		}
	}
	if fresh {
		if err := migrations.Fresh(db); err != nil {
			log.Println(err)
			return false
		}
	}
	if seed {
		if err := migrations.Seeder(db); err != nil {
			log.Println(err)
			return false

		}
	}

	return true
}
