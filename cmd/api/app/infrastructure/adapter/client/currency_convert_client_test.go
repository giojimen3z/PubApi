package client_test

import (
	"encoding/json"
	"os"

	"github.com/PubApi/cmd/api/test/builder"
	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	. "github.com/onsi/ginkgo"
)

const (
	beerId = 1
)

var _ = Describe("Client", func() {
	Context("Rest client", func() {
		var (

			fakeUrl               = "https://api.cambio.today/v1/quotes/COP/USD/json?quantity=1500&key=6392|h_2OeBxS2ibfZ^D1cA1o_3cYBQNUD*Pm"
		)

		var _ = BeforeSuite(func() {
			client := resty.New()
			httpmock.ActivateNonDefault(client.GetClient())
		})

		BeforeEach(func() {
			httpmock.Reset()
			_ = os.Setenv("SCOPE", "local")



		})

		AfterEach(func() {
			os.Clearenv()
		})

		var _ = AfterSuite(func() {
			httpmock.DeactivateAndReset()
		})

		When("a new valid  request is received  and convert currency successfully", func() {
			It("should return CurrencyConversion struct and nil error", func() {
				client := resty.New()
				currency := builder.NewCurrencyDataBuilder().Build()
				currencyConversionExpected := builder.NewCurrencyConversionDataBuilder().Build()
				fixture, _ := json.Marshal(currency)
				responder := httpmock.NewStringResponder(200, string(fixture))
				httpmock.RegisterResponder("GET", fakeUrl, responder)
				client.R().SetResult(currencyConversionExpected).Get(fakeUrl)



				//Expect(err).Should(BeNil())

			})
		})
		When("a new valid  request is received  and failed convert currency", func() {
			It("should return CurrencyConversion struct and nil error", func() {



			})
		})

	})
})
