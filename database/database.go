package database

import (
	"strconv"

	"github.com/abhishekgautam2808/sampleservice2/models"
	"github.com/gin-gonic/gin"
)

func Allsqldata() []models.User {
	var users []models.User
	models.DB.Find(&users)
	return users
}

func Sqldatabyid(c *gin.Context) models.User {
	var user models.User
	id, _ := strconv.ParseUint(c.Param("id"), 0, 0)
	user.ID = uint(id)
	_ = models.DB.Where("id = ?", user.ID).Find(&user).Error
	return user
}
