package application

import (
	"github.com/PubApi/cmd/api/app/domain/model"
	"github.com/PubApi/cmd/api/app/domain/service"
	"github.com/PubApi/pkg/apierrors"
)

// GetBeerApplication is the initial flow entry to get one beer
type GetBeerApplication interface {
	// Handler is the input for access to get one beer
	Handler(id int64) (model.Beer, apierrors.ApiError)
}
type GetBeer struct {
	GetBeerService service.GetBeerService
}

func (getBeer *GetBeer) Handler(id int64) (model.Beer, apierrors.ApiError) {
	return getBeer.GetBeerService.GetBeer(id)
}
