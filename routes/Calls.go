package routes

import (
	"github.com/gin-gonic/gin"
	"pager-service/controllers/controllers"
)

func Calls(r *gin.Engine) {
	r.POST("/calls/create", controllers.Call)
	r.POST("/calls/verify", controllers.Verify)
}
