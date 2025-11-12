package router

import (
	"Gin/controll/admin"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.GET("/new", func(c *gin.Context) { //这里可以使用中间件，里面存在next和abort方法一个是继续执行，一个是中断执行
			c.JSON(200, gin.H{
				"message": "aaaaaa",
			})
		})
		user.GET("/name", admin.Userindex)
	}
}
