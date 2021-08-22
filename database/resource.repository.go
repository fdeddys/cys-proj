package database

import (
	"github.com/fdeddys/tes/model"
)

func SaveResource(resource *model.Resource) (statusCode int, statusDesc string) {
	statusCode = 200
	statusDesc = "OK"

	db := GetDbCon()
	db.Debug()
	if r := db.Create(&resource); r.Error != nil {
		statusCode = 500
		statusDesc = r.Error.Error()
	}

	return
}

func FindResourceByUsername(username string) ([]model.Resource, error) {

	db := GetDbCon()
	db.Debug()

	var resources []model.Resource
	r := db.Where("user_name = ? ", username).Find(&resources).Error
	if r != nil {
		return resources, r
	}
	return resources, nil
}
