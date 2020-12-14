package application

import (
	"fmt"
	"strings"

	"github.com/PubApi/cmd/api/app/domain/model"
	"github.com/PubApi/cmd/api/app/domain/service"
	"github.com/PubApi/pkg/apierrors"
	"github.com/PubApi/pkg/logger"
)

const (
	loggerErrorGetBeer         = "Error Getting the beer with id:%v from service  [Class: GetBeerBoxPrice][Method:Handler]"
	loggerErrorConvertCurrency = "Error Converting Currency from service [Class: GetBeerBoxPrice][Method:Handler]"
)

// GetBeerBoxPriceApplication is the initial flow entry to get the beer box price
type GetBeerBoxPriceApplication interface {
	// Handler is the input for et the beer box price
	Handler(id int64, currency string, quantity int64) (model.BeerBox, apierrors.ApiError)
}

type GetBeerBoxPrice struct {
	GetBeerService         service.GetBeerService
	ConvertCurrencyService service.ConvertCurrencyService
	GetBeerBoxPriceService service.GetBeerBoxPriceService
}

func (getBeerBoxPrice *GetBeerBoxPrice) Handler(id int64, currency string, quantity int64) (model.BeerBox, apierrors.ApiError) {

	beer, err := getBeerBoxPrice.GetBeerService.GetBeer(id)

	if err != nil {
		logger.Error(fmt.Sprintf(loggerErrorGetBeer, id), err)
		return model.BeerBox{}, err
	}

	currencyFill := getBeerBoxPrice.fillCurrency(beer, currency)

	currencyConversion, err := getBeerBoxPrice.ConvertCurrencyService.ConvertCurrency(currencyFill)

	if err != nil {
		logger.Error(loggerErrorConvertCurrency, err)
		return model.BeerBox{}, err
	}

	beerBox := getBeerBoxPrice.GetBeerBoxPriceService.GetBeerBoxPrice(quantity, currencyConversion)

	return beerBox, nil
}

func (getBeerBoxPrice *GetBeerBoxPrice) fillCurrency(beer model.Beer, currency string) model.Currency {

	beerCurrency := strings.ToUpper(beer.Currency)
	currencyRequest := strings.ToUpper(currency)
	currencyFill := model.Currency{

		Source:   beerCurrency,
		Target:   currencyRequest,
		Quantity: beer.Price,
	}

	return currencyFill
}
