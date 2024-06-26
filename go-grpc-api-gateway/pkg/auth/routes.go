package auth

import (
	"github.com/ARunni/go-grpc-api-gateway/pkg/auth/routes"
	"github.com/ARunni/go-grpc-api-gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)

}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}

func (svc *ServiceClient) AdminLogin(ctx *gin.Context) {
	routes.AdminLogin(ctx, svc.Client)
}
func (svc *ServiceClient) AdminRegister(ctx *gin.Context) {
	routes.AdminRegister(ctx, svc.Client)
}

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/auth")
	routes.POST("/register", svc.Register)
	routes.POST("/login", svc.Login)
	routes.POST("/adminlogin", svc.AdminLogin)
	routes.POST("/adminregister", svc.AdminRegister)

	return svc
}
