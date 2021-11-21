package middleware

import (
	"encoding/json"
	"net/http"

	entity "github.com/garuda-hacks/go-api-hackathon/app/model/entity"
	"github.com/garuda-hacks/go-api-hackathon/pkg/helper"
	"github.com/gin-gonic/gin"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/garuda-hacks/go-api-hackathon/config"
	"github.com/garuda-hacks/go-entities-hackathon/entities/response"
	"github.com/garuda-hacks/go-library-hackathon/security/jwt"
)

// PaAuth patient app auth.
func PaAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		requestID, _ := context.Get("RequestID")

		maker, err := jwt.NewJWTMaker(config.C.AuthKey.PaSecret, "", 0)
		if err != nil {
			context.JSON(http.StatusUnauthorized, response.Response{
				Meta: response.Meta{
					Message:   response.RespMeta.TelErrUserNotFound,
					RequestID: requestID.(string),
				},
			})
			context.Abort()
			return
		}

		validate, err := maker.VerifyToken(context.Request.Header.Get("Authorization"))
		if validate == nil || err != nil {
			context.JSON(http.StatusUnauthorized, response.Response{
				Meta: response.Meta{
					Message:   response.RespMeta.TelErrUserNotFound,
					RequestID: requestID.(string),
				},
			})
			context.Abort()
			return
		}

		var pubUser entity.User
		jsonData, _ := json.Marshal(validate.Claims.(jwtgo.MapClaims))

		// check signature
		json.Unmarshal(jsonData, &pubUser)
		data := helper.GinStorerData{
			Context: context,
			Payload: pubUser,
			Key:     "UserAuth",
		}

		var ginStorer = helper.NewStore(context, pubUser, "UserAuth")
		ginStorer.Setter(data)
		context.Next()
	}
}
