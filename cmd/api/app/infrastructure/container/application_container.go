package container

import "github.com/PubApi/cmd/api/app/application"

func getCreateBeerApplication() application.CreateBeerApplication {
	return &application.CreateBeer{CreateBeerService: getCreateBeerService()}
}
