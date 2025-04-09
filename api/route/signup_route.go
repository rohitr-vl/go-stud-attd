package route

import (
	"clean-architecture/bootstrap"
	"clean-architecture/psql"
	"clean-architecture/repository"
	"time"

	"github.com/gin-gonic/gin"
)

func NewSignupRouter(env *bootstrap.Environment, timeout time.Duration, db psql.PsqlRepository, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	sc := controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(ur, timeout),
		Env:           env,
	}
	group.POST("/signup", sc.Signup)
}
