package router

import "github.com/gin-gonic/gin"

func ClassRouter(r *gin.Engine) {
	class := r.Group("/class")
	{
		//r.Query   r.Param  可以获取url中的参数   r.PostForm 可以获取post表单中的参数
		class.GET("/classlist", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "ciohecosheichspoidcopsdcsdc",
			})
		})
	}
}
