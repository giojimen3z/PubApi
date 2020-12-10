package adapter

import (
	"database/sql"
	"fmt"

	"github.com/PubApi/cmd/api/app/domain/exception"
	"github.com/PubApi/cmd/api/app/domain/model"
	"github.com/PubApi/cmd/api/app/infrastructure/config"
	"github.com/PubApi/pkg/logger"
)

const (
	loggerErrorTransactionBeer = "an error occurred  in the  transaction to save the beer: %s   [Class: BeerMysqlRepository][Method:Save]"
	loggerErrorSavingBeer      = "an error occurred  in the  transaction to save the beer, Error:%s   [Class: BeerMysqlRepository][Method:Save]"
	queryToSaveBeer            = `INSERT INTO beer (id, name, brewery, country, price, currency)VALUES(?,?,?,?,?,?)`
)

// BeerMysqlRepository represent the mysql repository
type BeerMysqlRepository struct {
	WriteClient *sql.DB
	ReadClient  *sql.DB
}

func (beerMysqlRepository *BeerMysqlRepository) Save(beer model.Beer) (err error) {
	var tx *sql.Tx

	defer func() {
		config.CloseConnections(err, tx, nil, nil)
	}()
	tx, err = beerMysqlRepository.WriteClient.Begin()
	if err != nil {

		logger.Error(fmt.Sprintf(loggerErrorTransactionBeer, beer.Name), err)
		return exception.InternalServerError{ErrMessage: err.Error()}
	}
	_, err = beerMysqlRepository.WriteClient.Exec(queryToSaveBeer,
		beer.BeerId,
		beer.Name,
		beer.Brewery,
		beer.Country,
		beer.Price,
		beer.Currency)

	if err != nil {
		logger.Error(fmt.Sprintf(loggerErrorSavingBeer, err.Error()), err)
		return exception.InternalServerError{ErrMessage: err.Error()}
	}

	return err
}

func (beerMysqlRepository *BeerMysqlRepository) ListBeer() (beersList []model.Beer, err error) {

	return beersList, nil
}
