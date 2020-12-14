package client_test

import (
	"os"

	"github.com/PubApi/cmd/api/app/domain/model"
	"github.com/PubApi/cmd/api/app/infrastructure/adapter/client"
	"github.com/PubApi/cmd/api/test/builder"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	insertQueryBeer = "INSERT INTO beer"
	beerListQuery   = "SELECT (.+) FROM beer"
)
const (
	beerId = 1
)

var _ = Describe("Client", func() {
	Context("Beer Mysql Repository", func() {
		var (
			currencyConvertClient client.CurrencyConvertClient
		)
		BeforeEach(func() {
			_ = os.Setenv("SCOPE", "local")
			currencyConvertClient = client.CurrencyConvertClient{}

		})

		AfterEach(func() {
			os.Clearenv()
		})
		When("a new valid  request is received  and convert currency successfully", func() {
			It("should return CurrencyConversion struct and nil error", func() {

				currency := builder.NewCurrencyDataBuilder().Build()
				currencyConversionExpected := builder.NewCurrencyConversionDataBuilder().Build()

				currencyConversion, err := currencyConvertClient.GetCurrency(currency)

				Expect(err).Should(BeNil())
				Expect(currencyConversionExpected).Should(Equal(currencyConversion))
			})
		})
		When("a new valid  request is received  and failed convert currency", func() {
			It("should return CurrencyConversion struct and nil error", func() {

				currency := model.Currency{}
				currencyConversionExpected := model.CurrencyConversion{}

				currencyConversion, err := currencyConvertClient.GetCurrency(currency)

				Expect(err).Should(BeNil())
				Expect(currencyConversionExpected).Should(Equal(currencyConversion))
			})
		})

	})
})
