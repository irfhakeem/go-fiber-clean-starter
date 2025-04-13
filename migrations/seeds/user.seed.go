package seeds

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/irfhakeem/go-fiber-clean-starter/dto"
	"github.com/irfhakeem/go-fiber-clean-starter/entity"
	"gorm.io/gorm"
)

func UserSeeder(db *gorm.DB) error {
	jsonFile, err := os.Open("migrations/json/user.json")
	if err != nil {
		return dto.ErrOpenFile
	}
	defer jsonFile.Close()

	jsonData, _ := io.ReadAll(jsonFile)

	var users []entity.User
	if err := json.Unmarshal(jsonData, &users); err != nil {
		return dto.ErrUnmarshalJSON
	}

	hasTable := db.Migrator().HasTable(&entity.User{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&entity.User{}); err != nil {
			return err
		}
	}

	for _, user := range users {
		isData := db.Find(&user, "email = ?", user.Email).RowsAffected
		if isData == 0 {
			if err := db.Create(&user).Error; err != nil {
				log.Printf("error seeding user: %v", err)
				return dto.ErrSeed
			}
		}
	}

	return nil
}
