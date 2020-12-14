package container

import "github.com/PubApi/cmd/api/app/application"

func getCreateBeerApplication() application.CreateBeerApplication {
	return &application.CreateBeer{CreateBeerService: getCreateBeerService()}
}

func getListBeerApplication() application.ListBeerApplication {
	return &application.ListBeer{ListBeerService: getListBeerService()}
}

func getBeerApplication() application.GetBeerApplication {
	return &application.GetBeer{GetBeerService: getBeerService()}
}
func getBeerBoxPriceApplication() application.GetBeerBoxPriceApplication {
	return &application.GetBeerBoxPrice{
		GetBeerService:         getBeerService(),
		ConvertCurrencyService: getConvertCurrencyService(),
		GetBeerBoxPriceService: getBeerBoxPriceService(),
	}
}
