package beer

import (
	"fmt"
	"net/http"

	"github.com/PubApi/cmd/api/app/application"
	"github.com/PubApi/cmd/api/app/domain/model"
	"github.com/PubApi/pkg/apierrors"
	"github.com/PubApi/pkg/logger"
	"github.com/gin-gonic/gin"
)

var (
	invalidBodyErr = apierrors.NewBadRequestApiError("invalid request")
	successMassage = "the beer %s was created successfully"
)

// CreateBeerController  used for inject the use case
type CreateBeerController struct {
	CreateBeerApplication application.CreateBeerApplication
}

func (createBeerController *CreateBeerController) MakeCreateBeer(context *gin.Context) {

	beer := model.Beer{}

	if err := context.ShouldBindJSON(&beer); err != nil {
		context.JSON(invalidBodyErr.Status(), invalidBodyErr)
		return
	}

	err := createBeerController.CreateBeerApplication.Handler(beer)

	if err != nil {
		logger.Error(err.Message(), err)
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusCreated, fmt.Sprintf(successMassage, beer.Name))

}
