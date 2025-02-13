package User

import "gorm.io/gorm"

func GetUserInformationWithUsername(username string, db *gorm.DB) (string, string) {
	query := db.Table("users")

}
