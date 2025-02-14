package User

import (
	"gorm.io/gorm"
)

func GetUserInformationByUsername(username string, db *gorm.DB) (DbUser, error) {
	query := db.Table("users").Where("username = ?", username)
	var userData DbUser

	err := query.Find(&userData).Error
	return userData, err
}
