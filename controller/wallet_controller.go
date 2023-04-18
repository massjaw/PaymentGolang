package controller

import (
	"Merchant-Bank/middlewares"
	app_error "Merchant-Bank/model"
	"Merchant-Bank/model/dto/req"
	"Merchant-Bank/usecase"
	"Merchant-Bank/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WalletController struct {
	MainController
	router  *gin.RouterGroup
	usecase usecase.WalletUsecase
}

func NewWalletController(route *gin.RouterGroup, usecase usecase.WalletUsecase) *WalletController {
	controller := WalletController{
		router:  route,
		usecase: usecase,
	}
	WalletGroup := route.Group("/wallet")
	WalletGroup.Use(middlewares.JwtAuthMiddleware())
	WalletGroup.GET("/", controller.GetBalance)
	WalletGroup.POST("/topup", controller.TopUp)

	return &controller
}

func (c *WalletController) GetBalance(ctx *gin.Context) {
	userId, err := utils.ExtractTokenID(ctx)
	if err != nil {
		c.Failed(ctx, http.StatusBadRequest, "", app_error.UnknownError(""))
		return
	}
	res, err := c.usecase.GetWallet(userId)
	if err != nil {
		c.Failed(ctx, http.StatusInternalServerError, "", fmt.Errorf("can't found wallet"))
		return
	}
	c.Success(ctx, http.StatusCreated, "01", "wallet found", res)
}

func (c *WalletController) TopUp(ctx *gin.Context) {
	var topUp *req.TopUp
	userId, _ := utils.ExtractTokenID(ctx)

	if err := ctx.BindJSON(&topUp); err != nil {
		log.Println("binding")
		c.Failed(ctx, http.StatusBadRequest, "", app_error.UnknownError(""))
		return
	}

	res, err := c.usecase.TopUp(topUp, userId)
	if err != nil {
		log.Println("TopUp Procces")
		c.Failed(ctx, http.StatusInternalServerError, "", fmt.Errorf("top-up failed"))
		return
	}
	c.Success(ctx, http.StatusCreated, "01", "top-up success", res)
}
