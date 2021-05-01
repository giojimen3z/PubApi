package user

import (
	"fmt"
	"net/http"

	"github.com/PubApi/cmd/api/app/application/user"
	"github.com/PubApi/cmd/api/app/domain/model"
	"github.com/PubApi/pkg/apierrors"
	"github.com/PubApi/pkg/logger"
	"github.com/gin-gonic/gin"
)

var (
	invalidBodyErr = apierrors.NewBadRequestApiError("invalid request")
	successMassage = "the user %s was created successfully"
)

// RegisterUserController  used for inject the use case
type RegisterUserController struct {
	RegisterUserApplication user.RegisterUserApplication
}

func (registerUserController *RegisterUserController) MakeRegisterUser(context *gin.Context) {

	user := model.User{}

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(invalidBodyErr.Status(), invalidBodyErr)
		return
	}

	err := registerUserController.RegisterUserApplication.Handler(user)

	if err != nil {
		logger.Error(err.Message(), err)
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusCreated, fmt.Sprintf(successMassage, user.Name))

}
