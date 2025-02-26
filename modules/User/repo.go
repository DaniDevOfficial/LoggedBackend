package User

import (
	"errors"
	"gorm.io/gorm"
)

func GetUserInformationByUsername(username string, db *gorm.DB) (DbUser, error) {
	query := db.Table("users").Where("username = ?", username)
	var userData DbUser

	result := query.First(&userData)

	if result.Error != nil {
		return userData, result.Error
	}
	if result.RowsAffected == 0 {
		return userData, NotFoundError
	}

	return userData, nil
}

var NotFoundError = errors.New("user not found")

func GetUserInformationById(userId string, db *gorm.DB) (DbUser, error) {
	var userData DbUser
	query := db.Table("users").Where("id = ?", userId)

	result := query.First(&userData)
	if result.Error != nil {
		return userData, result.Error
	}
	if result.RowsAffected == 0 {
		return userData, NotFoundError
	}

	return userData, nil
}

func MarkUserAsClaimed(userId string, claimUserData ClaimUser, db *gorm.DB) error {
	result := db.Table("users").
		Where("id = ?", userId).
		Updates(claimUserData)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return NotFoundError
	}

	return nil
}

func IsUserAdmin(userId string, db *gorm.DB) bool {
	var count int64

	result := db.Table("roles").Where("user_id = ?", userId).Where("role = 'admin'").Count(&count)

	if result.Error != nil {
		return false
	}
	return count > 0
}

func UsernameAlreadyInUse(username string, db *gorm.DB) bool {
	var count int64

	result := db.Table("users").Where("username = ?", username).Count(&count)

	if result.Error != nil {
		return false
	}
	return count > 0
}

func CreateNewUser(userData NewAccountRequest, db *gorm.DB) (string, error) {

	query := db.Table("users").Create(&userData)
	if query.Error != nil {
		return "", query.Error
	}

}
