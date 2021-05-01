package beer

import (
	"net/http"
	"strconv"

	"github.com/PubApi/cmd/api/app/application/beer"
	"github.com/gin-gonic/gin"
)

// GetBeerController  used for inject the use case
type GetBeerController struct {
	GetBeerApplication beer.GetBeerApplication
}

func (getBeerController *GetBeerController) MakeGetBeer(context *gin.Context) {
	id, _ := strconv.ParseUint(context.Param("id"), 10, 16)
	beerID := int64(id)

	beer, err := getBeerController.GetBeerApplication.Handler(beerID)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, beer)
}
