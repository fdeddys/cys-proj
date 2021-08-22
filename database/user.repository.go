package database

import (
	"github.com/fdeddys/tes/model"
)

func SaveUser(user *model.User) error {

	db := GetDbCon()
	db.Debug()
	dbUser := db.Create(&user)
	if dbUser.Error != nil {
		return dbUser.Error
	}
	return nil
}

func FindUserByUsername(username string) (model.User, error) {

	db := GetDbCon()
	db.Debug()

	var user model.User
	r := db.Where("user_name = ? ", username).First(&user).Error
	if r != nil {
		return user, r
	}
	return user, nil
}
