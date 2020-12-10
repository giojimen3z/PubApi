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
)

var _ = Describe("Service", func() {
	Context("List Beer", func() {
		var (
			repositoryMock  *mock.BeerRepositoryMock
			listBeerService service.ListBeer
		)
		BeforeEach(func() {
			repositoryMock = new(mock.BeerRepositoryMock)
			listBeerService = service.ListBeer{
				BeerRepository: repositoryMock,
			}

		})

		AfterEach(func() {
			os.Clearenv()
		})
		When("a new valid request is received", func() {
			It("should return beer list and nil error", func() {

				beerList := []model.Beer{builder.NewBikeDataBuilder().Build()}
				repositoryMock.On("ListBeer").Return(beerList, nil)

				beers, err := listBeerService.ListBeer()

				Expect(err).Should(BeNil())
				Expect(beerList).Should(Equal(beers))
				repositoryMock.AssertExpectations(GinkgoT())
			})
		})
		When("a new invalid request is received", func() {
			It("should return error", func() {
				beerList := []model.Beer{}
				errorRepository := errors.New("some type of parameters isn't correct")
				errorExpected := "Message: Error getting the beers from repository;Error Code: bad_request;Status: 400;Cause: []"
				repositoryMock.On("ListBeer").Return(beerList, errorRepository)

				beers, err := listBeerService.ListBeer()

				Expect(err).Should(Not(BeNil()))
				Expect(beerList).Should(Equal(beers))
				Expect(errorExpected).Should(Equal(err.Error()))
				repositoryMock.AssertExpectations(GinkgoT())
			})
		})
		When("a new valid request is received but the dba is empty", func() {
			It("should return error", func() {
				beerList := []model.Beer{}
				errorExpected := "Message: Error with the  information received, the Beers are empty;Error Code: internal_server_error;Status: 500;Cause: [Error with the  information received, the Beers are empty]"
				repositoryMock.On("ListBeer").Return(beerList, nil)

				beers, err := listBeerService.ListBeer()

				Expect(err).Should(Not(BeNil()))
				Expect(beerList).Should(Equal(beers))
				Expect(errorExpected).Should(Equal(err.Error()))
				repositoryMock.AssertExpectations(GinkgoT())
			})
		})
	})
})
