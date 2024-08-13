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

func NewRefreshTokenRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, model.CollectionUser)
	rtu := usecase.NewRefreshTokenUsecase(ur, timeout)

	rtc := &controller.RefreshTokenController{
		RefreshTokenUsecase: rtu,
		Env:                 env,
	}

	group.POST("/refresh", rtc.RefreshToken)
}
