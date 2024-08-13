package controller

import (
	"github.com/gin-gonic/gin"
	"go-ecommerce/bootstrap"
	"go-ecommerce/payload"
	"go-ecommerce/usecase"
	"go-ecommerce/utils/passwordutil"
	"net/http"
)

type LoginController struct {
	LoginUsecase usecase.LoginUsecase
	Env          *bootstrap.Env
}

func (lc *LoginController) Login(c *gin.Context) {
	var request payload.LoginRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	user, err := lc.LoginUsecase.GetUserByEmail(c.Request.Context(), request.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	if !passwordutil.CompareHashedPassword(request.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{})
		return
	}

	accessToken, err := lc.LoginUsecase.CreateAccessToken(&user, lc.Env.AccessTokenSecret, lc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	refreshToken, err := lc.LoginUsecase.CreateRefreshToken(&user, lc.Env.RefreshTokenSecret, lc.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	loginResponse := payload.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, loginResponse)
	return
}
