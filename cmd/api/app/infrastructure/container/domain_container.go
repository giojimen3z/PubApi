package container

import "github.com/PubApi/cmd/api/app/domain/service"

func getCreateBeerService() service.CreateBeerService {
	return &service.CreateBeer{
		BeerRepository: getCreateBeerRepository(),
	}
}

