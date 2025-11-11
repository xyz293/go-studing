package admin

import (
	"github.com/gin-gonic/gin"
)

func Userindex(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "userindex",
	})
}
