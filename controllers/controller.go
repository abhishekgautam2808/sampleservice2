package controllers

import (
	"net/http"

	"github.com/abhishekgautam2808/sampleservice2/database"
	"github.com/gin-gonic/gin"
)

func FindAll(c *gin.Context) {
	data := database.Allsqldata()
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func GetSingleUser(c *gin.Context) {
	data := database.Sqldatabyid(c)
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})

}
