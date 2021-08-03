package flusher_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"github.com/ozoncp/ocp-requirement-api/internal/flusher"
	"github.com/ozoncp/ocp-requirement-api/internal/mocks"
	"github.com/ozoncp/ocp-requirement-api/internal/models"
)

var _ = Describe("Flusher", func() {

	var (
		mockCtrl     *gomock.Controller
		repository   *mocks.MockRepo
		myFlusher    flusher.Flusher
		requirements []models.Requirement
		expectedArgs [][]models.Requirement
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		repository = mocks.NewMockRepo(mockCtrl)
		requirements = []models.Requirement{
			{1, 1, ""},
			{2, 2, ""},
			{3, 3, ""}}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("Storage is fine", func() {
		When("bulk size is less than 1", func() {
			BeforeEach(func() {
				myFlusher = flusher.NewFlusher(-1, repository)
			})

			It("should not add any record and return empty slice.", func() {

				repository.EXPECT().AddEntities(gomock.Any()).Times(0)

				result := myFlusher.Flush(requirements)

				gomega.Expect(result).Should(gomega.Equal([]models.Requirement{}))

			})
		})
		When("bulk size greater than 1", func() {
			BeforeEach(func() {
				myFlusher = flusher.NewFlusher(2, repository)
				expectedArgs = [][]models.Requirement{
					{
						{1, 1, ""},
						{2, 2, ""},
					},
					{
						{3, 3, ""},
					},
				}
			})

			It("should add and return whole slice", func() {

				for _, args := range expectedArgs {
					repository.EXPECT().AddEntities(args)
				}

				result := myFlusher.Flush(requirements)

				gomega.Expect(result).Should(gomega.Equal(requirements))

			})

		})

		When("bulk size greater than slice", func() {
			BeforeEach(func() {
				myFlusher = flusher.NewFlusher(99, repository)
				expectedArgs = [][]models.Requirement{
					{
						{1, 1, ""},
						{2, 2, ""},
						{3, 3, ""},
					},
				}
			})

			It("should add whole slice by one repo method call", func() {

				repository.EXPECT().AddEntities(expectedArgs[0]).Times(1)

				result := myFlusher.Flush(requirements)

				gomega.Expect(result).Should(gomega.Equal(requirements))

			})

		})
	})

	Context("Storage have just crushed during the flush", func() {

		BeforeEach(func() {
			myFlusher = flusher.NewFlusher(2, repository)
			expectedArgs = [][]models.Requirement{
				{
					{1, 1, ""},
					{2, 2, ""},
				},
			}
			repository.EXPECT().AddEntities(expectedArgs[0])
			repository.EXPECT().AddEntities(gomock.Any()).Return(errors.New("error"))
		})

		It("should add and return one batch", func() {

			result := myFlusher.Flush(requirements)

			gomega.Expect(result).Should(gomega.Equal(expectedArgs[0]))

		})

	})

	Context("Storage is down", func() {

		BeforeEach(func() {
			myFlusher = flusher.NewFlusher(2, repository)
			expectedArgs = [][]models.Requirement{{}}
		})

		It("shouldn't add anything and should return empty slice", func() {
			repository.EXPECT().AddEntities(gomock.Any()).Return(errors.New("error")).Times(2)

			result := myFlusher.Flush(requirements)

			gomega.Expect(result).Should(gomega.Equal(expectedArgs[0]))

		})

	})
})
