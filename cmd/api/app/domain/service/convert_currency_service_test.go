package service_test

import (
	"errors"
	"os"

	"github.com/PubApi/cmd/api/app/domain/model"
	"github.com/PubApi/cmd/api/app/domain/service"
	"github.com/PubApi/cmd/api/test/builder"
	"github.com/PubApi/cmd/api/test/mock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	mockParameter "github.com/stretchr/testify/mock"
)

var _ = Describe("Service", func() {
	Context("Convert Currency", func() {
		var (
			clientMock                *mock.CurrencyClientMock
			getConvertCurrencyService service.ConvertCurrency
		)
		BeforeEach(func() {
			clientMock = new(mock.CurrencyClientMock)
			getConvertCurrencyService = service.ConvertCurrency{
				ConvertCurrencyClient: clientMock,
			}

		})

		AfterEach(func() {
			os.Clearenv()
		})
		When("a new valid request is received", func() {
			It("should return currency conversion and nil error", func() {

				currency := builder.NewCurrencyDataBuilder().Build()
				currencyConversionExpected := builder.NewCurrencyConversionDataBuilder().Build()
				clientMock.On("GetCurrency", mockParameter.Anything).Return(currencyConversionExpected, nil)

				currencyConversion, err := getConvertCurrencyService.ConvertCurrency(currency)

				Expect(err).Should(BeNil())
				Expect(currencyConversionExpected).Should(Equal(currencyConversion))
				clientMock.AssertExpectations(GinkgoT())
			})
		})
		When("a new valid request is received", func() {
			It("should return  error", func() {
				errorRepository := errors.New("error converting the currency into repository")
				errorExpected := "Message: error converting the currency into repository;Error Code: not_found;Status: 404;Cause: []"
				currency := builder.NewCurrencyDataBuilder().Build()
				currencyConversionExpected := model.CurrencyConversion{}
				clientMock.On("GetCurrency", mockParameter.Anything).Return(currencyConversionExpected, errorRepository)

				currencyConversion, err := getConvertCurrencyService.ConvertCurrency(currency)

				Expect(err).Should(Not(BeNil()))
				Expect(errorExpected).Should(Equal(err.Error()))
				Expect(currencyConversionExpected).Should(Equal(currencyConversion))
				clientMock.AssertExpectations(GinkgoT())
			})
		})
	})
})
