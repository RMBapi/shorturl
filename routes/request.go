package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.POST("/api/v1/shorturl", HandleRequest)
	server.GET("/api/v1/shorturl", HandleRequest)

}
