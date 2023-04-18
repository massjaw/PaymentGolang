package controller

import (
	"Merchant-Bank/middlewares"
	app_error "Merchant-Bank/model"
	"Merchant-Bank/model/dto/req"
	"Merchant-Bank/usecase"
	"Merchant-Bank/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PaymentController struct {
	MainController
	router  *gin.RouterGroup
	usecase usecase.PaymentUsecase
}

func NewPaymentController(route *gin.RouterGroup, usecase usecase.PaymentUsecase) *PaymentController {
	controller := PaymentController{
		router:  route,
		usecase: usecase,
	}
	PaymentGroup := route.Group("/Payment")
	PaymentGroup.Use(middlewares.JwtAuthMiddleware())
	PaymentGroup.POST("/", controller.AddTransfer)

	return &controller
}

func (c *PaymentController) AddTransfer(ctx *gin.Context) {
	var newTransfer *req.Transfer
	senderId, _ := utils.ExtractTokenID(ctx)

	if err := ctx.BindJSON(&newTransfer); err != nil {
		c.Failed(ctx, http.StatusBadRequest, "", app_error.UnknownError(""))
		return
	}

	if newTransfer.ReceiptUsername == "" {
		c.Failed(ctx, http.StatusBadRequest, "X01", app_error.InvalidError("one or more required fields are missing"))
		return
	}
	if newTransfer.TransferAmount < 1 {
		c.Failed(ctx, http.StatusBadRequest, "X01", app_error.InvalidError("invalid amount"))
		return
	}

	res, err := c.usecase.PaymentProcces(newTransfer, senderId)
	if err != nil {
		c.Failed(ctx, http.StatusInternalServerError, "", fmt.Errorf("failed to transfer fund"))
		return
	}

	c.Success(ctx, http.StatusCreated, "01", "Payment success", res)
}
