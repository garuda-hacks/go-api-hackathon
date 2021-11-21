package controller

import (
	"net/http"
	"strconv"

	"github.com/garuda-hacks/go-api-hackathon/app/usecase"
	logger "github.com/garuda-hacks/go-api-hackathon/infrastructure/io"
	response "github.com/garuda-hacks/go-entities-hackathon/entities/response"

	"github.com/gin-gonic/gin"
)

type ugcController struct {
	UgcUsecase usecase.UgcUsecase
}

type UgcController interface {
	ListUgc(c *gin.Context)
}

func NewUgcController(
	uu usecase.UgcUsecase,
) UgcController {
	return &ugcController{uu}
}

func (d ugcController) ListUgc(c *gin.Context) {
	params, _ := strconv.Atoi(c.Query("idinterest"))
	logger.Infof("Input Data Search By Interest", params)
	requestid, _ := c.Get("RequestID")
	result, err := d.UgcUsecase.List(params)
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
