package application_test

import (
	"errors"

	"github.com/PubApi/cmd/api/app/application"
	"github.com/PubApi/cmd/api/app/domain/model"
	"github.com/PubApi/cmd/api/app/domain/service"
	"github.com/PubApi/cmd/api/test/builder"
	"github.com/PubApi/cmd/api/test/mock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	mockParameter "github.com/stretchr/testify/mock"
)

const (
	currency = "USD"
	quantity = 1
)

var _ = Describe("Handler", func() {
	Context("Get Beer Box Price", func() {
		var (
			repositoryMock         *mock.BeerRepositoryMock
			clientMock             *mock.CurrencyClientMock
			getBeerBoxPriceUseCase application.GetBeerBoxPrice
		)
		BeforeEach(func() {
			repositoryMock = new(mock.BeerRepositoryMock)
			clientMock = new(mock.CurrencyClientMock)
			getBeerService := &service.GetBeer{
				BeerRepository: repositoryMock,
			}
			convertCurrencyService := &service.ConvertCurrency{
				ConvertCurrencyClient: clientMock,
			}
			getBeerBoxPrice := &service.GetBeerBoxPrice{}
			getBeerBoxPriceUseCase = application.GetBeerBoxPrice{
				GetBeerService:         getBeerService,
				ConvertCurrencyService: convertCurrencyService,
				GetBeerBoxPriceService: getBeerBoxPrice,
			}
		})
		When("a new valid  request is received", func() {
			It("should return beer box price  and nil error", func() {
				currencyConversion := builder.NewCurrencyConversionDataBuilder().Build()
				beerExpected := builder.NewBeerDataBuilder().Build()
				beerBoxExpected := builder.NewBeerBoxDataBuilder().Build()
				repositoryMock.On("GetBeerByID", mockParameter.Anything).Return(beerExpected, nil)
				clientMock.On("GetCurrency", mockParameter.Anything).Return(currencyConversion, nil)

				beerBox, err := getBeerBoxPriceUseCase.Handler(beerId, currency, quantity)

				Expect(err).Should(BeNil())
				Expect(beerBoxExpected).Should(Equal(beerBox))
				repositoryMock.AssertExpectations(GinkgoT())
				clientMock.AssertExpectations(GinkgoT())
			})
		})
		When("a new invalid  request is received and getBeerService Failed", func() {
			It("should return error", func() {
				errorRepository := errors.New("some type of parameters isn't correct")
				beerExpected := model.Beer{}
				beerBoxExpected := model.BeerBox{}
				errorExpected := "Message: The beer id:1 isn´t exists;Error Code: not_found;Status: 404;Cause: []"
				repositoryMock.On("GetBeerByID", mockParameter.Anything).Return(beerExpected, errorRepository)

				beerBox, err := getBeerBoxPriceUseCase.Handler(beerId, currency, quantity)

				Expect(err).Should(Not(BeNil()))
				Expect(errorExpected).Should(Equal(err.Error()))
				Expect(beerBoxExpected).Should(Equal(beerBox))
				repositoryMock.AssertExpectations(GinkgoT())
			})
		})
		When("a new invalid  request is received and convertCurrencyService Failed", func() {
			It("should return error", func() {
				currencyConversion := model.CurrencyConversion{}
				errorRepository := errors.New("error converting the currency into repository")
				beerExpected := builder.NewBeerDataBuilder().Build()
				beerBoxExpected := model.BeerBox{}
				errorExpected := "Message: error converting the currency into repository;Error Code: not_found;Status: 404;Cause: []"
				repositoryMock.On("GetBeerByID", mockParameter.Anything).Return(beerExpected, nil)
				clientMock.On("GetCurrency", mockParameter.Anything).Return(currencyConversion, errorRepository)

				beerBox, err := getBeerBoxPriceUseCase.Handler(beerId, currency, quantity)

				Expect(err).Should(Not(BeNil()))
				Expect(errorExpected).Should(Equal(err.Error()))
				Expect(beerBoxExpected).Should(Equal(beerBox))
				repositoryMock.AssertExpectations(GinkgoT())
				clientMock.AssertExpectations(GinkgoT())
			})
		})
	})
})
