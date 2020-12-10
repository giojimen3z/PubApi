package service_test

import (
	"errors"
	"os"

	"github.com/PubApi/cmd/api/app/domain/service"
	"github.com/PubApi/cmd/api/test/builder"
	"github.com/PubApi/cmd/api/test/mock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Service", func() {
	Context("Create Beer", func() {
		var (
			repositoryMock    *mock.BeerRepositoryMock
			beerCreateService service.CreateBeer
		)
		BeforeEach(func() {
			repositoryMock = new(mock.BeerRepositoryMock)
			beerCreateService = service.CreateBeer{
				BeerRepository: repositoryMock,
			}

		})

		AfterEach(func() {
			os.Clearenv()
		})
		When("a new valid beer request is received", func() {
			It("should return nil error", func() {

				beer := builder.NewBikeDataBuilder().Build()
				repositoryMock.On("Save", beer).Return(nil)

				err := beerCreateService.CreateBeer(beer)

				Expect(err).Should(BeNil())
				repositoryMock.AssertExpectations(GinkgoT())
			})
		})
		When("a new valid beer request is received with a invalid id", func() {
			It("should return error", func() {

				errorMock := errors.New("the id:1 is invalid")
				beer := builder.NewBikeDataBuilder().Build()
				repositoryMock.On("Save", beer).Return(errorMock)
				errorExpected := "Message: Beer id:1 already exists;Error Code: Conflict;Status: 409;Cause: []"

				err := beerCreateService.CreateBeer(beer)

				Expect(err).Should(Not(BeNil()))
				Expect(errorExpected).Should(Equal(err.Error()))
				repositoryMock.AssertExpectations(GinkgoT())
			})
		})
	})
})
