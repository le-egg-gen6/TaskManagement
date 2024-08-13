package route

import (
	"github.com/gin-gonic/gin"
	"go-ecommerce/api/controller"
	"go-ecommerce/bootstrap"
	"go-ecommerce/model"
	"go-ecommerce/mongo"
	"go-ecommerce/repository"
	usecase "go-ecommerce/usecase/impl"
	"time"
)

func NewTaskRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	tr := repository.NewTaskRepository(db, model.CollectionTask)
	tc := &controller.TaskController{
		TaskUsecase: usecase.NewTaskUsecase(tr, timeout),
	}
	group.GET("/task", tc.Fetch)
	group.POST("/task", tc.Create)
}
