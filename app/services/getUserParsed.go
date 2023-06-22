package services

import (
	"sufrimiento/app/db"
	"sufrimiento/app/models"
)

func GetUserParsed(id interface{}) (*models.UserDTO, int) {
	user := new(models.User)
	result := db.Ctx.First(user, id)

	if result.Error != nil {
		return nil, 404
	}

	userParsed := new(models.UserDTO)
	user.ParseToDTO(userParsed)

	return userParsed, 200
}
