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

func NewSignupRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, model.CollectionUser)
	su := usecase.NewSignupUsecase(ur, timeout)
	sc := &controller.SignupController{
		SignupUsecase: su,
		Env:           env,
	}
	group.POST("/signup", sc.Signup)
}
