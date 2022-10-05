//go:build wireinject
// +build wireinject

package injector

import (
	"activity/controller"
	"activity/repository"
	"activity/service"

	"github.com/google/wire"
	"gorm.io/gorm"
)

var activitySet = wire.NewSet(
	repository.NewActivityRepository,
	service.NewActivityService,
	controller.NewActivityController,
)

func InitActivity(db *gorm.DB) controller.ActivityController {
	wire.Build(
		activitySet,
	)
	return nil
}
