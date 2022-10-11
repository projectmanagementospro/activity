package controller

import (
	"activity/helper"
	"activity/models/web"
	"activity/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SubActivityController interface {
	All(context *gin.Context)
	FindById(context *gin.Context)
	Insert(context *gin.Context) //
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type subactivityController struct {
	subactivityService service.SubActivityService
}

func NewSubActivityController(subactivityService service.SubActivityService) SubActivityController {
	return &subactivityController{
		subactivityService: subactivityService,
	}
}

func (ac *subactivityController) All(context *gin.Context) {
	subactivity := ac.subactivityService.All()
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: "",
		Data:   subactivity,
	}
	context.JSON(http.StatusOK, webResponse)
}

func (ac *subactivityController) FindById(context *gin.Context) {
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	ok := helper.NotFoundError(context, err)
	if ok {
		return
	}
	subactivity, err := ac.subactivityService.FindById(uint(id))
	ok = helper.NotFoundError(context, err)
	if ok {
		return
	}
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: "",
		Data:   subactivity,
	}
	context.JSON(http.StatusOK, webResponse)
}

func (ac *subactivityController) Insert(context *gin.Context) {
	var request web.SubActivityRequest
	err := context.BindJSON(&request)
	ok := helper.InternalServerError(context, err)
	if ok {
		return
	}
	subactivity, err := ac.subactivityService.Create(request)

	ok = helper.InternalServerError(context, err)
	if ok {
		return
	}

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: "",
		Data:   subactivity,
	}
	context.JSON(http.StatusOK, webResponse)
}

func (ac *subactivityController) Update(context *gin.Context) {
	var request web.SubActivityUpdateRequest
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	ok := helper.NotFoundError(context, err)
	if ok {
		return
	}
	request.ID = uint(id)
	err = context.BindJSON(&request)
	ok = helper.ValidationError(context, err)
	if ok {
		return
	}
	subactivity, err := ac.subactivityService.Update(request)
	ok = helper.InternalServerError(context, err)
	if ok {
		return
	}
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: "",
		Data:   subactivity,
	}
	context.JSON(http.StatusOK, webResponse)
}

func (ac *subactivityController) Delete(context *gin.Context) {
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	ok := helper.NotFoundError(context, err)
	if ok {
		return
	}
	err = ac.subactivityService.Delete(uint(id))
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
