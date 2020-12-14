package application_test

import (
	"errors"

	"github.com/PubApi/cmd/api/app/application"
	"github.com/PubApi/cmd/api/app/domain/service"
	"github.com/PubApi/cmd/api/test/builder"
	"github.com/PubApi/cmd/api/test/mock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Handler", func() {
	Context("Create Beer", func() {
		var (
			repositoryMock    *mock.BeerRepositoryMock
			beerCreateUseCase application.CreateBeer
		)
		BeforeEach(func() {
			repositoryMock = new(mock.BeerRepositoryMock)
			beerCreateService := &service.CreateBeer{
				BeerRepository: repositoryMock,
			}
			beerCreateUseCase = application.CreateBeer{
				CreateBeerService: beerCreateService,
			}

		})

		When("a new valid beer request is received", func() {
			It("should return nil error", func() {

				beer := builder.NewBeerDataBuilder().Build()
				repositoryMock.On("Save", beer).Return(nil)

				err := beerCreateUseCase.Handler(beer)

				Expect(err).Should(BeNil())
				repositoryMock.AssertExpectations(GinkgoT())
			})
		})
		When("a new invalid beer request is received", func() {
			It("should return error", func() {

				errorMock := errors.New("Error 1062: Duplicate entry '1' for key 'beer.PRIMARY'")
				beer := builder.NewBeerDataBuilder().Build()
				repositoryMock.On("Save", beer).Return(errorMock)
				errorExpected := "Message: Beer id:1 already exists;Error Code: Conflict;Status: 409;Cause: []"

				err := beerCreateUseCase.Handler(beer)

				Expect(err).Should(Not(BeNil()))
				Expect(errorExpected).Should(Equal(err.Error()))
				repositoryMock.AssertExpectations(GinkgoT())
			})
		})
	})
})
