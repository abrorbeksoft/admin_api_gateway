package main

import (
	"fmt"
	"github.com/abrorbeksoft/admin_api_gateway/api"
	"github.com/abrorbeksoft/admin_api_gateway/config"
	"github.com/abrorbeksoft/admin_api_gateway/servise"
)

func main(){
	cfg:=config.Load();
	grpcServices,err:=servise.NewGrpcClients(cfg)
	if err!= nil {
		fmt.Println(err.Error())
	}

	server:=api.New(&api.RouterOptions{
		Service: grpcServices,
	})

	server.Run(":8080")
}