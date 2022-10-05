package routes

import (
	"activity/middleware"

	"activity/injector"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewActivityRoutes(db *gorm.DB, route *gin.Engine) {
	pactivityController := injector.InitActivity(db)
	pactivityRoute := route.Group("/api/v1/pdata")
	pactivityRoute.Use(middleware.ErrorHandler())
	pactivityRoute.Use(cors.Default())
	pactivityRoute.GET("/", pactivityController.All)
	pactivityRoute.GET("/:id", pactivityController.FindById)
	pactivityRoute.POST("/", pactivityController.Insert)
	pactivityRoute.PUT("/:id", pactivityController.Update)
	pactivityRoute.DELETE("/:id", pactivityController.Delete)
}
