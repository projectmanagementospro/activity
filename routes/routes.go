package routes

import (
	"activity/middleware"

	"activity/injector"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewActivityRoutes(db *gorm.DB, route *gin.Engine) {
	activityController := injector.InitActivity(db)
	activityRoute := route.Group("/api/v1/activity")
	activityRoute.Use(middleware.ErrorHandler())
	activityRoute.Use(cors.Default())
	activityRoute.GET("/", activityController.All)
	activityRoute.GET("/:id", activityController.FindById)
	activityRoute.POST("/", activityController.Insert)
	activityRoute.PUT("/:id", activityController.Update)
	activityRoute.DELETE("/:id", activityController.Delete)
}
