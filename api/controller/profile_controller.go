package controller

import (
	"github.com/gin-gonic/gin"
	"go-ecommerce/usecase"
	"net/http"
)

type ProfileController struct {
	ProfileUsecase usecase.ProfileUsecase
}

func (pc *ProfileController) Fetch(c *gin.Context) {
	userID := c.GetString("x-user-id")

	profile, err := pc.ProfileUsecase.GetProfileByID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusOK, profile)
}
