package main

import (
	"context"
	"fmt"
	auth "github.com/abrorbeksoft/admin_api_gateway/genproto/auth_service"
	"google.golang.org/grpc"
)

func main(){
	var (
		res *auth.LoginResponse
	)
	conAuthService,err:=grpc.Dial("localhost:8000",grpc.WithInsecure())
	if err != nil {
		fmt.Println(err.Error())
	}
	authService:=auth.NewAuthServiceClient(conAuthService)
	res,err=authService.LoginUser(context.Background(),&auth.LoginRequest{
		Login: "sdfldfmlds",
		Password: "sdmfldmflsd",
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(res)
	fmt.Println("Hello world")
}