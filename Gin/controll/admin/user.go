package admin
import (
	"github.com/gin-gonic/gin"
	"Gin/model"
)

func Userindex(c *gin.Context) {
	userlist :=[]model.User{}
	model.DB.Find(&userlist)
	c.JSON(200, gin.H{
		"message": "userindex",
		"data":userlist,
	})
}
