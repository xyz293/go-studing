package admin

import (
	"Gin/model"

	"github.com/gin-gonic/gin"
)

func Userindex(c *gin.Context) {
	userlist := []model.User{}
	model.DB.Raw("select * from users where id =?",1).Find(&userlist)
	c.JSON(200, gin.H{
		"message": "userindex",
		"data":    userlist,
	})
}
