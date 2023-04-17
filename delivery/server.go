package delivery

import (
	"Merchant-Bank/config"
	"Merchant-Bank/controller"
	"Merchant-Bank/manager"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type AppServer struct {
	usecaseManager manager.UsecaseManager
	engine         *gin.Engine
	host           string
}

func (a *AppServer) v1() {
	v1Routes := a.engine.Group("/v1")
	a.UserController(v1Routes)
}

func (a *AppServer) UserController(rg *gin.RouterGroup) {
	controller.NewUserController(rg, a.usecaseManager.UserUsecase())
}

func (a *AppServer) Run() {
	a.v1()
	err := a.engine.Run(a.host)
	defer func() {
		if err := recover(); err != nil {
			log.Println("Application Failed to run", err)
		}
	}()
	if err != nil {
		panic(err)
	}
}

func Server() *AppServer {
	r := gin.Default()
	c := config.NewConfig()
	infraManager := manager.NewInfraManager(c)
	repoManager := manager.NewRepoManager(infraManager)
	usecaseManager := manager.NewUsecaseManager(repoManager)
	host := fmt.Sprintf(":%s", c.ApiPort)
	return &AppServer{
		usecaseManager: usecaseManager,
		engine:         r,
		host:           host,
	}
}
