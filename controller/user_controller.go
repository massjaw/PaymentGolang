package controller

import (
	app_error "Merchant-Bank/model"
	"Merchant-Bank/model/dto/req"
	"Merchant-Bank/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	MainController
	router  *gin.RouterGroup
	usecase usecase.UserUsecase
}

func NewUserController(route *gin.RouterGroup, usecase usecase.UserUsecase) *UserController {
	controller := UserController{
		router:  route,
		usecase: usecase,
	}
	route.POST("/register", controller.Register)
	route.POST("/login", controller.Login)

	return &controller
}

func (c *UserController) Register(ctx *gin.Context) {
	var newUser req.UserRegist
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if newUser.Password != newUser.PasswordConfirm {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "confirmation password invalid"})
		return
	}

	user, err := c.usecase.Register(&newUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

func (c *UserController) Login(ctx *gin.Context) {
	var input req.UserLogin

	if err := ctx.BindJSON(&input); err != nil {
		c.Failed(ctx, http.StatusBadRequest, "", app_error.InvalidError("invalid request body"))
		return
	}

	token, err := c.usecase.Login(&input)

	if err != nil {
		c.Failed(ctx, http.StatusBadRequest, "", app_error.InvalidError("Username or Password is Incorrect"))
		return
	}

	c.Success(ctx, http.StatusOK, "", token, nil)
}
