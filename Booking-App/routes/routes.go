package routes

import (
	"booking-app/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/status", controllers.GetStatus)

	r.POST("/book", controllers.BookTicket)
}
