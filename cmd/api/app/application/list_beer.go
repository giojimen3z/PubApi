package application

import (
	"github.com/PubApi/cmd/api/app/domain/model"
	"github.com/PubApi/cmd/api/app/domain/service"
	"github.com/PubApi/pkg/apierrors"
)

// ListBeerApplication is the initial flow entry to list all beers
type ListBeerApplication interface {
	// Handler is the input for access to list all  beers
	Handler() ([]model.Beer, apierrors.ApiError)
}

type ListBeer struct {
	ListBeerService service.ListBeerService
}

func (listBeer *ListBeer) Handler() ([]model.Beer, apierrors.ApiError) {

	return listBeer.ListBeerService.ListBeer()

}
