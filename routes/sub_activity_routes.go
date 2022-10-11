package routes

import (
	"activity/middleware"

	"activity/injector"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewSubActivityRoutes(db *gorm.DB, route *gin.Engine) {
	subactivityController := injector.InitSubActivity(db)
	subactivityroute := route.Group("/api/v1/subactivity")
	subactivityroute.Use(middleware.ErrorHandler())
	subactivityroute.Use(cors.Default())
	subactivityroute.GET("/", subactivityController.All)
	subactivityroute.GET("/:id", subactivityController.FindById)
	subactivityroute.POST("/", subactivityController.Insert)
	subactivityroute.PUT("/:id", subactivityController.Update)
	subactivityroute.DELETE("/:id", subactivityController.Delete)
}
