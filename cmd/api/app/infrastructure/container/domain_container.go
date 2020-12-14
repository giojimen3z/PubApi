package container

import "github.com/PubApi/cmd/api/app/domain/service"

func getCreateBeerService() service.CreateBeerService {
	return &service.CreateBeer{
		BeerRepository: getCreateBeerRepository(),
	}
}

func getListBeerService() service.ListBeerService {
	return &service.ListBeer{
		BeerRepository: getCreateBeerRepository(),
	}
}

func getBeerService() service.GetBeerService {
	return &service.GetBeer{
		BeerRepository: getCreateBeerRepository(),
	}
}
func getConvertCurrencyService() service.ConvertCurrencyService {
	return &service.ConvertCurrency{
		ConvertCurrencyClient: getConvertCurrencyClient(),
	}
}

func getBeerBoxPriceService() service.GetBeerBoxPriceService {
	return &service.GetBeerBoxPrice{}
}
