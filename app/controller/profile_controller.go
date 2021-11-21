package controller

import (
	"net/http"

	req "github.com/garuda-hacks/go-api-hackathon/app/model/request"
	"github.com/garuda-hacks/go-api-hackathon/app/usecase"
	logger "github.com/garuda-hacks/go-api-hackathon/infrastructure/io"
	response "github.com/garuda-hacks/go-entities-hackathon/entities/response"
	"github.com/gin-gonic/gin"
)

type profilController struct {
	ProfileUsecase usecase.ProfilUsecase
}

type ProfilController interface {
	FindProfil(c *gin.Context)
}

func NewProfilController(
	pu usecase.ProfilUsecase,
) ProfilController {
	return &profilController{pu}
}

func (pc *profilController) FindProfil(c *gin.Context) {
	var input req.RegProfileUser
	requestid, _ := c.Get("RequestID")
	cl := logger.WithFields(logger.Fields{"UserController": "FindProfil"})
	cl.Infof("[INFO] Header values: %v", c.Request.Header)

	if err := c.ShouldBind(&input); err != nil {
		cl.Errorf("[ERROR] %v", err.Error())
		resp := response.Response{
			Meta: response.Meta{
				Message:   response.RespMeta.TelErrShouldBindError,
				RequestID: requestid.(string),
			},
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	if err := input.Validate(); err != nil {
		cl.Errorf("[ERROR] %v", err.Error())
		resp := response.Response{
			Meta: response.Meta{
				Message:   response.RespMeta.TelErrValidate,
				RequestID: requestid.(string),
			},
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	result, err := pc.ProfileUsecase.FindParam(input)
	if err != nil {
		cl.Errorf("[ERROR] Error on field: %v", response.RespMeta.TelErrNonActiveUser)
		resp := response.Response{
			Meta: response.Meta{
				Message:   response.RespMeta.TelErrNonActiveUser,
				RequestID: requestid.(string),
			},
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	resp := response.Response{
		Meta: response.Meta{
			RequestID: requestid.(string),
		},
		Data: result,
	}
	c.JSON(http.StatusOK, resp)
}
