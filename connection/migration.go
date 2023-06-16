package connection

import (
	"fly/user"

	"gorm.io/gorm"
)

func MigrationUser(db *gorm.DB) error {
	var users user.Users

	if err := db.Limit(1).Find(&users).Error; err != nil {
		return err
	}

	if users.ID == 0 {
		allUser := user.SeederUser()
		for _, singleUser := range *allUser {
			if err := db.Create(singleUser).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
