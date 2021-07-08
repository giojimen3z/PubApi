package beer

import (
	"fmt"
	"net/http"

	"github.com/PubApi/cmd/api/app/domain/model"
	"github.com/PubApi/cmd/api/app/domain/port"
	"github.com/PubApi/pkg/apierrors"
	"github.com/PubApi/pkg/logger"
)

const (
	errorIDExist          = "Beer id:%v already exists"
	logErrorInvalidBeerID  = "Beer id:%v already exists  [Class: CreateBeerService][Method:CreateBeer]"
)

type CreateBeerService interface {
	// CreateBeer Send to repository the Beer
	CreateBeer(beer model.Beer) apierrors.ApiError
}

type CreateBeer struct {
	BeerRepository port.BeerRepository
}

func (createBeer *CreateBeer) CreateBeer(beer model.Beer) apierrors.ApiError {

	errorRepository := createBeer.BeerRepository.Save(beer)

	if errorRepository != nil {
		logger.Error(fmt.Sprintf(logErrorInvalidBeerID, beer.BeerId), errorRepository)
		err := apierrors.NewApiError(fmt.Sprintf(errorIDExist, beer.BeerId), http.StatusText(http.StatusConflict), http.StatusConflict, nil)
		return err
	}

	return nil

}
