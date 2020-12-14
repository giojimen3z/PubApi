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
	beerId = 1
)

var _ = Describe("Handler", func() {
	Context("Get Beer", func() {
		var (
			repositoryMock *mock.BeerRepositoryMock
			getBeerUseCase application.GetBeer
		)
		BeforeEach(func() {
			repositoryMock = new(mock.BeerRepositoryMock)
			getBeerService := &service.GetBeer{
				BeerRepository: repositoryMock,
			}
			getBeerUseCase = application.GetBeer{
				GetBeerService: getBeerService,
			}
		})
		When("a new valid  request is received", func() {
			It("should return beer  and nil error", func() {
				beerExpected := builder.NewBeerDataBuilder().Build()
				repositoryMock.On("GetBeerByID", mockParameter.Anything).Return(beerExpected, nil)

				beer, err := getBeerUseCase.Handler(beerId)

				Expect(err).Should(BeNil())
				Expect(beerExpected).Should(Equal(beer))
				repositoryMock.AssertExpectations(GinkgoT())
			})
		})
		When("a new invalid  request is received", func() {
			It("should return error", func() {
				errorRepository := errors.New("some type of parameters isn't correct")
				beerExpected := model.Beer{}
				errorExpected := "Message: The beer id:1 isn´t exists;Error Code: not_found;Status: 404;Cause: []"
				repositoryMock.On("GetBeerByID", mockParameter.Anything).Return(beerExpected, errorRepository)

				beer, err := getBeerUseCase.Handler(beerId)

				Expect(err).Should(Not(BeNil()))
				Expect(errorExpected).Should(Equal(err.Error()))
				Expect(beerExpected).Should(Equal(beer))
				repositoryMock.AssertExpectations(GinkgoT())
			})
		})
	})
})
