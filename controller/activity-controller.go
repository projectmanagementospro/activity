package controller

import (
	"activity/helper"
	"activity/models/web"
	"activity/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ActivityController interface {
	All(context *gin.Context)
	FindById(context *gin.Context)
	Insert(context *gin.Context) //
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type activityController struct {
	activityService service.ActivityService
}

func NewActivityController(activityService service.ActivityService) ActivityController {
	return &activityController{
		activityService: activityService,
	}
}

func (ac *activityController) All(context *gin.Context) {
	activity := ac.activityService.All()
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: "",
		Data:   activity,
	}
	context.JSON(http.StatusOK, webResponse)
}

func (ac *activityController) FindById(context *gin.Context) {
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	ok := helper.NotFoundError(context, err)
	if ok {
		return
	}
	activity, err := ac.activityService.FindById(uint(id))
	ok = helper.NotFoundError(context, err)
	if ok {
		return
	}
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: "",
		Data:   activity,
	}
	context.JSON(http.StatusOK, webResponse)
}

func (ac *activityController) Insert(context *gin.Context) {
	var u web.ActivityRequest
	err := context.BindJSON(&u)
	ok := helper.InternalServerError(context, err)
	if ok {
		return
	}
	activity, err := ac.activityService.Create(u)

	ok = helper.InternalServerError(context, err)
	if ok {
		return
	}

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: "",
		Data:   activity,
	}
	context.JSON(http.StatusOK, webResponse)
}

func (ac *activityController) Update(context *gin.Context) {
	var u web.ActivityUpdateRequest
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	ok := helper.NotFoundError(context, err)
	if ok {
		return
	}
	u.ID = uint(id)
	err = context.BindJSON(&u)
	ok = helper.ValidationError(context, err)
	if ok {
		return
	}
	activity, err := ac.activityService.Update(u)
	ok = helper.InternalServerError(context, err)
	if ok {
		return
	}
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: "",
		Data:   activity,
	}
	context.JSON(http.StatusOK, webResponse)
}

func (ac *activityController) Delete(context *gin.Context) {
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	ok := helper.NotFoundError(context, err)
	if ok {
		return
	}
	err = ac.activityService.Delete(uint(id))
	ok = helper.NotFoundError(context, err)
	if ok {
		return
	}
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: "",
		Data:   "Activity has been removed",
	}
	context.JSON(http.StatusOK, webResponse)
}
