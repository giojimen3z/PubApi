package application

import (
	"github.com/PubApi/cmd/api/app/domain/model"
	"github.com/PubApi/cmd/api/app/domain/service"
	"github.com/PubApi/pkg/apierrors"
)

// CreateBeerApplication is the initial flow entry to create one beer
type CreateBeerApplication interface {
	// Handler is the input for access to create one beer
	Handler(beer model.Beer) apierrors.ApiError
}
type CreateBeer struct {
	CreateBeerService service.CreateBeerService
}

func (createBeer *CreateBeer) Handler(beer model.Beer) apierrors.ApiError {
	return createBeer.CreateBeerService.CreateBeer(beer)
}
