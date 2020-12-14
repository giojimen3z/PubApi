package beer

import (
	"net/http"

	"github.com/PubApi/cmd/api/app/application"
	"github.com/gin-gonic/gin"
)

// ListCreateBeerController  used for inject the use case
type ListBeerController struct {
	ListBeerApplication application.ListBeerApplication
}

func (listBeerController *ListBeerController) MakeListBeer(context *gin.Context) {

	BeerList, err := listBeerController.ListBeerApplication.Handler()

	if err != nil {
		context.JSON(err.Status(), err)
	}

	context.JSON(http.StatusOK, BeerList)

}
