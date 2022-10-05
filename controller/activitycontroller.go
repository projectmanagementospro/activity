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
	Insert(context *gin.Context)
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

func (c *activityController) All(context *gin.Context) {
	pactivity := c.activityService.All()
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: "",
		Data:   pactivity,
	}
	context.JSON(http.StatusOK, webResponse)
}

func (c *activityController) FindById(context *gin.Context) {
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	ok := helper.NotFoundError(context, err)
	if ok {
		return
	}
	pactivity, err := c.activityService.FindById(uint(id))
	ok = helper.NotFoundError(context, err)
	if ok {
		return
	}
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: "",
		Data:   pactivity,
	}
	context.JSON(http.StatusOK, webResponse)
}

func (c *activityController) Insert(context *gin.Context) {
	var u web.ActivityRequest
	err := context.BindJSON(&u)
	ok := helper.InternalServerError(context, err)
	if ok {
		return
	}
	pactivity, err := c.activityService.Create(u)
	println("ada")
	ok = helper.InternalServerError(context, err)
	if ok {
		return
	}
	println("tidak ada")
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: "",
		Data:   pactivity,
	}
	context.JSON(http.StatusOK, webResponse)
}

func (c *activityController) Update(context *gin.Context) {
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
	pactivity, err := c.activityService.Update(u)
	ok = helper.InternalServerError(context, err)
	if ok {
		return
	}
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: "",
		Data:   pactivity,
	}
	context.JSON(http.StatusOK, webResponse)
}

func (c *activityController) Delete(context *gin.Context) {
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	ok := helper.NotFoundError(context, err)
	if ok {
		return
	}
	err = c.activityService.Delete(uint(id))
	ok = helper.NotFoundError(context, err)
	if ok {
		return
	}
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: "",
		Data:   "Project charter has been removed",
	}
	context.JSON(http.StatusOK, webResponse)
}
