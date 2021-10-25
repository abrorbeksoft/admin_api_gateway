package v1

import (
	"context"
	"fmt"
	auth "github.com/abrorbeksoft/admin_api_gateway/genproto/auth_service"
	"github.com/abrorbeksoft/admin_api_gateway/servise"
	"github.com/gin-gonic/gin"
	"net/http"
)

type handlerV1 struct {
	services servise.ServiceManager
}

type HandlerV1Options struct {
	Service servise.ServiceManager
}

func New(options *HandlerV1Options) *handlerV1 {
	return &handlerV1{
		services: options.Service,
	}
}

func (h *handlerV1) Home(c *gin.Context)  {
	var (
		user *auth.RegisterResponse
		err error
	)
	user, err=h.services.AuthService().Register(context.Background(),&auth.RegisterRequest{
		Name: "Abrorbek",
		Surname: "Ubaydullayev",
		Login: "Abrorbek",
		Password: "sdmdsmcsdmckdm",
	})

	if err!= nil {
		fmt.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}
