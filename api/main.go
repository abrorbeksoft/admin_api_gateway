package api

import (
	v1 "github.com/abrorbeksoft/admin_api_gateway/api/handlers/v1"
	"github.com/abrorbeksoft/admin_api_gateway/servise"
	"github.com/gin-gonic/gin"
)

type RouterOptions struct {
	Service servise.ServiceManager
}

func New(opt *RouterOptions) *gin.Engine {
	route:=gin.New()
	apiV1:=route.Group("/v1")
	handlerV1:=v1.New(&v1.HandlerV1Options{
		Service: opt.Service,
	});

	apiV1.Use()
	{
		apiV1.GET("/",handlerV1.Home)
	}

	return route
}