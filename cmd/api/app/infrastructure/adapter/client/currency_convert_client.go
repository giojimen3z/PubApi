package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/PubApi/cmd/api/app/domain/model"
	"github.com/PubApi/cmd/api/app/infrastructure/config"
	"github.com/PubApi/pkg/logger"
)

const (
	convertCurrencyURL                   = "https://api.cambio.today/v1/quotes/%s/%s/json?quantity=%v&key=%v"
	loggerErrorSendingRequest            = "Error sending request to api.cambio.today [Class: CurrencyConvertClient][Method:GetCurrency]"
	loggerErrorGettingRequestInformation = "Error getting information from api.cambio.today [Class: CurrencyConvertClient][Method:GetCurrency]"
)

type CurrencyConvertClient struct {
}

func (currencyConvertClient *CurrencyConvertClient) GetCurrency(currency model.Currency) (model.CurrencyConversion, error) {

	var currencyConversion model.CurrencyConversion
	url := fmt.Sprintf(convertCurrencyURL, currency.Source, currency.Target, currency.Quantity, config.GetCurrencyApiKey())

	response, err := http.Get(url)

	if err != nil {
		logger.Error(loggerErrorSendingRequest, err)
		return model.CurrencyConversion{}, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logger.Error(loggerErrorGettingRequestInformation, err)
		return model.CurrencyConversion{}, err
	}

	json.Unmarshal(responseData, &currencyConversion)

	return currencyConversion, nil
}
