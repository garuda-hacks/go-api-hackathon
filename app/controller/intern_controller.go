package controller

import (
	"net/http"
	"strconv"

	"github.com/garuda-hacks/go-api-hackathon/app/usecase"
	logger "github.com/garuda-hacks/go-api-hackathon/infrastructure/io"
	response "github.com/garuda-hacks/go-entities-hackathon/entities/response"

	"github.com/gin-gonic/gin"
)

type internController struct {
	InternUsecase usecase.InternUsecase
}

type InternController interface {
	FindByID(c *gin.Context)
	ListIntern(c *gin.Context)
}

func NewInternController(
	iu usecase.InternUsecase,
) InternController {
	return &internController{iu}
}

func (d internController) ListIntern(c *gin.Context) {
	requestid, _ := c.Get("RequestID")
	cl := logger.WithFields(logger.Fields{"UserController": "ListScholarship"})
	cl.Infof("[INFO] Header values: %v", c.Request.Header)
	result, err := d.InternUsecase.List()
	logger.Infof("Result ", result)
	if err != nil {
		rsp := response.Response{
			Meta: response.Meta{
				Message:   response.RespMeta.TelErrUserNotFound,
				RequestID: requestid.(string),
			},
		}
		c.JSON(http.StatusBadRequest, rsp)
		return
	}
	rsp := response.Response{
		Meta: response.Meta{
			RequestID: requestid.(string),
		},
		Data: result,
	}
	c.JSON(
		http.StatusOK,
		rsp,
	)
}

func (d internController) FindByID(c *gin.Context) {
	params, _ := strconv.Atoi(c.Param("id"))
	// logger.Infof("Input Data Search Name", paramsearch)
	// paramsortcity := c.Query("city")
	// logger.Infof("Input Data Sort City", paramsortcity)
	// paramsortspeciality := c.Query("speciality")
	// logger.Infof("Input Data Sort Speciality", paramsortspeciality)
	// paramfilteridcity, _ := strconv.Atoi(c.Query("idcity"))
	// logger.Infof("Input Data Search City", paramfilteridcity)
	requestid, _ := c.Get("RequestID")
	cl := logger.WithFields(logger.Fields{"UserController": "ListScholarship"})
	cl.Infof("[INFO] Header values: %v", c.Request.Header)
	result, err := d.InternUsecase.Find(params)
	logger.Infof("Result ", result)
	if err != nil {
		rsp := response.Response{
			Meta: response.Meta{
				Message:   response.RespMeta.TelErrUserNotFound,
				RequestID: requestid.(string),
			},
		}
		c.JSON(http.StatusBadRequest, rsp)
		return
	}
	rsp := response.Response{
		Meta: response.Meta{
			RequestID: requestid.(string),
		},
		Data: result,
	}
	c.JSON(
		http.StatusOK,
		rsp,
	)
}
