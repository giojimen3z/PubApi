package mock

import (
	"github.com/PubApi/cmd/api/app/domain/model"
	"github.com/stretchr/testify/mock"
)

type CurrencyClientMock struct {
	mock.Mock
}

func (mock *CurrencyClientMock) GetCurrency(currency model.Currency) (model.CurrencyConversion, error) {
	args := mock.Called(currency)
	return args.Get(0).(model.CurrencyConversion), args.Error(1)
}
