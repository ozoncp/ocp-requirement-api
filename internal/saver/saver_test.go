package saver_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"github.com/ozoncp/ocp-requirement-api/internal/mocks"
	"github.com/ozoncp/ocp-requirement-api/internal/models"
	"github.com/ozoncp/ocp-requirement-api/internal/saver"
	"sort"
	"sync"
	"time"
)

var _ = Describe("Saver", func() {
	var (
		mockCtrl     *gomock.Controller
		flusher      *mocks.MockFlusher
		mySaver      saver.Saver
		requirements []models.Requirement
		expectedArgs []models.Requirement
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		flusher = mocks.NewMockFlusher(mockCtrl)
		mySaver = saver.NewSaver(5, flusher)
		requirements = []models.Requirement{
			{1, 1, ""},
			{2, 2, ""},
			{3, 3, ""}}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("Simple usage.", func() {

		It("flush is triggered by 'ticker' and 'close' from 2 to 3 times within 2 seconds", func() {
			flusher.EXPECT().Flush([]models.Requirement{}).MinTimes(2).MaxTimes(3).Return([]models.Requirement{})
			mySaver.Init()
			time.Sleep(2 * time.Second)
			mySaver.Close()
		})

		It("same as before, but one of calls flushes saved requirement", func() {
			flusher.EXPECT().Flush([]models.Requirement{}).MinTimes(1).MaxTimes(2).Return([]models.Requirement{})
			flusher.EXPECT().Flush(requirements[:1]).Return([]models.Requirement{}).Times(1)

			mySaver.Init()
			mySaver.Save(requirements[0])
			time.Sleep(2 * time.Second)
			mySaver.Close()
		})

		It("flusher fails to save batch for the first time, but succeed at the second time", func() {
			flusher.EXPECT().Flush([]models.Requirement{}).MinTimes(0).MaxTimes(1).Return([]models.Requirement{})
			flusher.EXPECT().Flush(requirements[:1]).Return(requirements[:1]).Times(1)
			flusher.EXPECT().Flush(requirements[:1]).Return([]models.Requirement{}).Times(1)

			mySaver.Init()
			mySaver.Save(requirements[0])
			time.Sleep(2 * time.Second)
			mySaver.Close()
		})

		It("still works with calling Init after Close", func() {
			flusher.EXPECT().Flush([]models.Requirement{}).MinTimes(2).MaxTimes(3).Return([]models.Requirement{})

			mySaver.Init()
			time.Sleep(2 * time.Second)
			mySaver.Close()

			flusher.EXPECT().Flush([]models.Requirement{}).MinTimes(2).MaxTimes(3).Return([]models.Requirement{})

			mySaver.Init()
			time.Sleep(2 * time.Second)
			mySaver.Close()
		})
	})

	Context("Concurrent usage.", func() {
		BeforeEach(func() {
			mySaver = saver.NewSaver(5, flusher)
			expectedArgs = []models.Requirement{
				requirements[0],
				requirements[0],
				requirements[0],
				requirements[1],
				requirements[1],
				requirements[1],
			}
		})

		It("data doesnt disappear", func() {
			totalSaved := make([]models.Requirement, 0)
			flusher.EXPECT().Flush(gomock.Any()).DoAndReturn(
				func(requirements []models.Requirement) []models.Requirement {
					totalSaved = append(totalSaved, requirements...)
					return make([]models.Requirement, 0)
				}).AnyTimes()

			mySaver.Init()
			wg := sync.WaitGroup{}
			wg.Add(2)
			go func() {
				for i := 0; i < 3; i++ {
					mySaver.Save(requirements[0])
				}
				defer wg.Done()
			}()
			go func() {
				for i := 0; i < 3; i++ {
					mySaver.Save(requirements[1])
				}
				defer wg.Done()
			}()
			wg.Wait()
			time.Sleep(2 * time.Second)

			mySaver.Close()

			sort.Slice(totalSaved, func(i int, j int) bool {
				return totalSaved[i].Id < totalSaved[j].Id
			})
			gomega.Expect(totalSaved).Should(gomega.Equal(expectedArgs))
		})

	})
})
