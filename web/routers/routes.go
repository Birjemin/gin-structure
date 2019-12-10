package routers

import (
	"github.com/birjemin/gin-structure/web/controllers"
	"github.com/gin-gonic/gin"
)

func SetRouters(g *gin.Engine) {
	book := g.Group("/api/v1/book")
	{
		// 添加 Get 请求路由
		book.GET("", controllers.Get)
		// 添加 GET 请求路由
		book.GET("/:id", controllers.GetBy)
		// 添加 POST 请求路由
		book.POST("", controllers.Post)
		// 添加 Put 请求路由
		book.PUT("/:id", controllers.PutBy)
		// 添加 Delete 请求路由
		book.DELETE("/:id", controllers.DeleteBy)
	}
}