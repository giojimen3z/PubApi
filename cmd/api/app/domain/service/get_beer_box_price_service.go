package service

import (
	"github.com/PubApi/cmd/api/app/domain/model"
)

const (
	Zero            = 0
	defaultQuantity = 6
)

type GetBeerBoxPriceService interface {
	// ConvertCurrency Send to repository the currency
	GetBeerBoxPrice(quantity int64, currencyConversion model.CurrencyConversion) model.BeerBox
}

type GetBeerBoxPrice struct{}

func (getBeerBoxPrice *GetBeerBoxPrice) GetBeerBoxPrice(quantity int64, currencyConversion model.CurrencyConversion) model.BeerBox {

	if quantity == Zero {
		quantity = defaultQuantity
	}

	price := float64(quantity) * currencyConversion.Currency.Amount

	beerBox := model.BeerBox{
		Price: price,
	}

	return beerBox
}
