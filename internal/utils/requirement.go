package utils

import (
	"errors"
	"github.com/ozoncp/ocp-requirement-api/internal/models"
	"math"
)

func SplitRequirementsToBulks(inSlice []models.Requirement, batchLen int) ([][]models.Requirement, error) {
	inSliceLen := len(inSlice)
	if batchLen < 1 {
		return nil, errors.New("batch len can't be less than 1")
	}

	if inSliceLen < batchLen {
		return nil, errors.New("len of the slice can't be less than batch len")
	}

	regularBatchesNumber := inSliceLen / batchLen

	batchNumber := int(math.Ceil(float64(inSliceLen) / float64(batchLen)))
	outSlice := make([][]models.Requirement, 0, batchNumber)

	for i := 0; i < regularBatchesNumber; i++ {
		offset := i * batchLen
		outSlice = append(outSlice, inSlice[offset:offset+batchLen])
	}

	if regularBatchesNumber != batchNumber {
		outSlice = append(outSlice, inSlice[regularBatchesNumber*batchLen:])
	}

	return outSlice, nil
}

func ConvertRequirementsSliceToMap(inSlice []models.Requirement) map[uint64]models.Requirement {
	resultMap := make(map[uint64]models.Requirement, len(inSlice))
	for _, v := range inSlice {
		resultMap[v.Id] = v
	}

	return resultMap
}
