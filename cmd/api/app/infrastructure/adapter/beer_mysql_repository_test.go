package adapter_test

import (
	"database/sql"
	"os"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/PubApi/cmd/api/app/domain/exception"
	"github.com/PubApi/cmd/api/app/infrastructure/adapter"
	"github.com/PubApi/cmd/api/test/builder"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	insertQueryBeer = "INSERT INTO beer"
)

var _ = Describe("Repository", func() {
	Context("Beer Mysql Repository", func() {
		var (
			db                  *sql.DB
			dbMock              sqlmock.Sqlmock
			beerMysqlRepository adapter.BeerMysqlRepository
		)
		BeforeEach(func() {
			db, dbMock, _ = sqlmock.New()
			beerMysqlRepository = adapter.BeerMysqlRepository{
				WriteClient: db,
			}

		})

		AfterEach(func() {
			os.Clearenv()
		})
		When("a new valid beer request is received  and save in dba", func() {
			It("should return nil error", func() {

				beer := builder.NewBikeDataBuilder().Build()

				dbMock.ExpectBegin()
				dbMock.ExpectExec(insertQueryBeer).WillReturnResult(sqlmock.NewResult(1, 1))
				dbMock.ExpectCommit()
				err := beerMysqlRepository.Save(beer)

				Expect(err).Should(BeNil())
			})
		})
		When("a new valid beer request and failed the transaction", func() {
			It("should return  error", func() {
				transactionErrorMessage := "an error happened when execute the transaction"
				beer := builder.NewBikeDataBuilder().Build()
				errorOnTransaction := exception.InternalServerError{ErrMessage: transactionErrorMessage}

				dbMock.ExpectBegin().WillReturnError(errorOnTransaction)

				err := beerMysqlRepository.Save(beer)

				Expect(err).Should(Not(BeNil()))
				Expect(errorOnTransaction).Should(Equal(err))
			})
		})

		When("a new valid beer request and failed insert into dba", func() {
			It("should return  error", func() {
				transactionErrorMessage := "an error happened when try insert into dba"
				beer := builder.NewBikeDataBuilder().Build()
				errorOnInsert := exception.InternalServerError{ErrMessage: transactionErrorMessage}
				dbMock.ExpectBegin()
				dbMock.ExpectExec(insertQueryBeer).WillReturnError(errorOnInsert)

				err := beerMysqlRepository.Save(beer)

				Expect(err).Should(Not(BeNil()))
				Expect(errorOnInsert).Should(Equal(err))
			})
		})
	})
})
