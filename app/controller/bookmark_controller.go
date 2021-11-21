package controller

import (
	"net/http"

	"github.com/garuda-hacks/go-api-hackathon/app/usecase"
	logger "github.com/garuda-hacks/go-api-hackathon/infrastructure/io"
	response "github.com/garuda-hacks/go-entities-hackathon/entities/response"

	"github.com/gin-gonic/gin"
)

type bookmarkController struct {
	BookmarkUsecase usecase.BookmarkUsecase
}

type BookmarkController interface {
	Listbookmark(c *gin.Context)
}

func NewBookmarkController(
	ubm usecase.BookmarkUsecase,
) BookmarkController {
	return &bookmarkController{ubm}
}

func (d bookmarkController) Listbookmark(c *gin.Context) {
	requestid, _ := c.Get("RequestID")
	cl := logger.WithFields(logger.Fields{"UserController": "Listbookmark"})
	cl.Infof("[INFO] Header values: %v", c.Request.Header)
	result, err := d.BookmarkUsecase.List()
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
