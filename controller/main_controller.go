package controller

import (
	"Merchant-Bank/model/dto/res"

	"github.com/gin-gonic/gin"
)

type MainController struct{}

func (m *MainController) Success(c *gin.Context, httpCode int, code string, msg string, data any) {
	res.NewSuccessJsonResponse(c, httpCode, code, msg, data).Send()
}

func (m *MainController) Failed(c *gin.Context, httpCode int, code string, err error) {
	res.NewErrorJsonResponse(c, httpCode, code, err).Send()
}
