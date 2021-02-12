package service_test

import (
	"os"

	"github.com/PubApi/cmd/api/app/domain/service"
	"github.com/PubApi/cmd/api/test/builder"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	zeroQuantity = 0
)

var _ = Describe("Service", func() {
	Context("Get Beer Box Price", func() {
		var (
			getBeerBoxPriceService service.GetBeerBoxPrice
		)
		BeforeEach(func() {
			getBeerBoxPriceService = service.GetBeerBoxPrice{}

		})

		AfterEach(func() {
			os.Clearenv()
		})
		When("a new valid request is received", func() {
			It("should return beer box price", func() {

				beerBoxExpected := builder.NewBeerBoxDataBuilder().WithPrice(2.4966).Build()
				currencyConversion := builder.NewCurrencyConversionDataBuilder().Build()

				beerBox := getBeerBoxPriceService.GetBeerBoxPrice(zeroQuantity, currencyConversion)

				Expect(beerBoxExpected).Should(Equal(beerBox))
			})
		})

	})
})