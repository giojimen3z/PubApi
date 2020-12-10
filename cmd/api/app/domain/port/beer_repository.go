package port

import "github.com/PubApi/cmd/api/app/domain/model"

// BeerRepository  use for all transactions about beer
type BeerRepository interface {
	//Save persist the beer data
	Save(beer model.Beer) (err error)
	//ListBeer get all beers from dba
	ListBeer() (beersList []model.Beer, err error)
}