package servise

import (
	"fmt"
	"github.com/abrorbeksoft/admin_api_gateway/config"
	auth "github.com/abrorbeksoft/admin_api_gateway/genproto/auth_service"
	"google.golang.org/grpc"
)

type ServiceManager interface {
	AuthService() auth.AuthServiceClient
}

type grpcClients struct {
	authService auth.AuthServiceClient
}

func NewGrpcClients(cfg *config.Config) (ServiceManager,error) {
	connAuthService,err:=grpc.Dial(fmt.Sprintf("%s:%s",cfg.AuthHost,cfg.AuthPort),grpc.WithInsecure())
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return &grpcClients{
		authService: auth.NewAuthServiceClient(connAuthService),
	}, nil
}

func (g *grpcClients) AuthService() auth.AuthServiceClient {
	return g.authService
}

