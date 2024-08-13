package controller

import (
	"github.com/gin-gonic/gin"
	"go-ecommerce/bootstrap"
	"go-ecommerce/usecase"
)

type LoginController struct {
	loginUsecase usecase.LoginUsecase
	env          *bootstrap.Env
}

func (lc *LoginController) Login(c *gin.Context) {
	var request
}
