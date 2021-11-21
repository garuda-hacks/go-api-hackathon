package controller

import (
	"net/http"
	"strconv"

	"github.com/garuda-hacks/go-api-hackathon/app/usecase"
	logger "github.com/garuda-hacks/go-api-hackathon/infrastructure/io"
	response "github.com/garuda-hacks/go-entities-hackathon/entities/response"

	"github.com/gin-gonic/gin"
)

type scholarController struct {
	ScholarUsecase usecase.ScholarUsecase
}

type ScholarController interface {
	FindByID(c *gin.Context)
	ListScholar(c *gin.Context)
}

func NewScholarController(
	us usecase.ScholarUsecase,
) ScholarController {
	return &scholarController{us}
}

func (d scholarController) ListScholar(c *gin.Context) {
	requestid, _ := c.Get("RequestID")
	cl := logger.WithFields(logger.Fields{"UserController": "ListScholarship"})
	cl.Infof("[INFO] Header values: %v", c.Request.Header)
	result, err := d.ScholarUsecase.List()
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

func (d scholarController) FindByID(c *gin.Context) {
	params, _ := strconv.Atoi(c.Param("id"))
	requestid, _ := c.Get("RequestID")
	cl := logger.WithFields(logger.Fields{"UserController": "ListScholarship"})
	cl.Infof("[INFO] Header values: %v", c.Request.Header)
	result, err := d.ScholarUsecase.Find(params)
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
