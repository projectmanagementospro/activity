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

var subactivitySet = wire.NewSet(
	repository.NewSubActivityRepository,
	service.NewSubActivityService,
	controller.NewSubActivityController,
)

func InitSubActivity(db *gorm.DB) controller.SubActivityController {
	wire.Build(
		subactivitySet,
	)
	return nil
}

func InitActivity(db *gorm.DB) controller.ActivityController {
	wire.Build(
		activitySet,
	)
	return nil
}
