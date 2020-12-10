package container

import (
	"database/sql"

	"github.com/PubApi/cmd/api/app/domain/port"
	"github.com/PubApi/cmd/api/app/infrastructure/adapter"
	"github.com/PubApi/cmd/api/app/infrastructure/config"
	"github.com/PubApi/cmd/api/app/infrastructure/controller/beer"
)

func GetCreateBeerController() *beer.CreateBeerController {
	return &beer.CreateBeerController{CreateBeerApplication: getCreateBeerApplication()}
}

func getCreateBeerRepository() port.BeerRepository {
	return &adapter.BeerMysqlRepository{
		WriteClient: getWriteConnectionClient(),
		ReadClient:  getReadConnectionClient(),
	}
}
func getWriteConnectionClient() *sql.DB {
	conn, _ := config.GetWriteConnection()
	return conn
}

func getReadConnectionClient() *sql.DB {
	conn, _ := config.GetReadConnection()
	return conn
}
