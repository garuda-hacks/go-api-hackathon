package controller

import (
	"net/http"

	req "github.com/garuda-hacks/go-api-hackathon/app/model/request"
	"github.com/garuda-hacks/go-api-hackathon/app/usecase"
	logger "github.com/garuda-hacks/go-api-hackathon/infrastructure/io"
	response "github.com/garuda-hacks/go-entities-hackathon/entities/response"

	"github.com/gin-gonic/gin"
)

type userController struct {
	UserUsecase usecase.UserUsecase
}

type UserController interface {
	LoginUsername(c *gin.Context)
}

func NewUserController(
	uu usecase.UserUsecase,
) UserController {
	return &userController{uu}
}

func (uc *userController) LoginUsername(c *gin.Context) {
	var input req.RegLoginEmail
	requestid, _ := c.Get("RequestID")
	cl := logger.WithFields(logger.Fields{"UserController": "LoginUsername"})
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
		cl.Errorf("[ERROR] Error on field: %v", err.Error())
		resp := response.Response{
			Meta: response.Meta{
				Message:   response.RespMeta.TelErrValidate,
				RequestID: requestid.(string),
			},
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	if uc.UserUsecase.IsDuplicateEmail(input.Username) {
		cl.Errorf("[ERROR] Error on field: %v", response.RespMeta.TelErrEmailAlreadyUsed)
		resp := response.Response{
			Meta: response.Meta{
				Message:   response.RespMeta.TelErrEmailAlreadyUsed,
				RequestID: requestid.(string),
			},
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	result := uc.UserUsecase.VerifyEmail(input.Username, input.Password)
	if result == false {
		cl.Errorf("[ERROR] Error on field: %v", response.RespMeta.TelErrVerifyEmail)
		resp := response.Response{
			Meta: response.Meta{
				Message:   response.RespMeta.TelErrVerifyEmail,
				RequestID: requestid.(string),
			},
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	resp := response.Response{
		Meta: response.Meta{
			Message:   response.RespMeta.TelSuccessLogin,
			RequestID: requestid.(string),
		},
		Data: result,
	}
	c.JSON(http.StatusOK, resp)
}
