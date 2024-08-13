package controller

import (
	"github.com/gin-gonic/gin"
	"go-ecommerce/model"
	"go-ecommerce/usecase"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type TaskController struct {
	TaskUsecase usecase.TaskUsecase
}

func (tc *TaskController) Create(c *gin.Context) {
	var task model.Task

	err := c.ShouldBindJSON(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	userID := c.GetString("x-user-id")
	task.ID = primitive.NewObjectID()

	task.UserID, err = primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	err = tc.TaskUsecase.Create(c.Request.Context(), &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (u *TaskController) Fetch(c *gin.Context) {
	userID := c.GetString("x-user-id")

	tasks, err := u.TaskUsecase.FetchByUserID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusOK, tasks)
}
