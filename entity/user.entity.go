package entity

import (
	"unicode"

	"github.com/irfhakeem/go-fiber-clean-starter/dto"
	"github.com/irfhakeem/go-fiber-clean-starter/helpers/constants"
	"github.com/irfhakeem/go-fiber-clean-starter/helpers/utils"
	"gorm.io/gorm"
)

type User struct {
	ID         int64              `json:"id"       gorm:"primaryKey;type:bigint;autoIncrement"`
	Email      string             `json:"email"    gorm:"varchar(255);unique;not null"`
	Password   string             `json:"password" gorm:"varchar(255);not null"`
	Name       string             `json:"name"     gorm:"varchar(255);not null"`
	Gender     constants.Gender   `json:"gender"   gorm:"gender;"`
	Role       constants.UserRole `json:"role"     gorm:"user_role;default:member"`
	Avatar     string             `json:"avatar"   gorm:"varchar(255);default:null"`
	IsVerified bool               `json:"is_verified" gorm:"default:false"`

	TimeStamps
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	var err error

	if u.Password != "" {
		n := len(u.Password)
		if n < 8 {
			return dto.ErrPasswordTooShort
		}

		if n > 20 {
			return dto.ErrPasswordTooLong
		}

		var hasNumber bool
		var hasUpperCase bool

		for _, r := range u.Password {
			if unicode.IsUpper(r) && unicode.IsLetter(r) {
				hasUpperCase = true
				continue
			}
			if unicode.IsNumber(r) {
				hasNumber = true
				continue
			}
		}

		if !hasNumber || !hasUpperCase {
			return dto.ErrPasswordWeak
		}

		u.Password, err = utils.HashPassword(u.Password)
		if err != nil {
			return err
		}
	}

	return nil
}
