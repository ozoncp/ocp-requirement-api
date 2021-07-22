package utils

import (
	"errors"
	"math"
)

func MakeSliceOfBatches(inSlice []int, batchLen int) ([][]int, error) {
	inSliceLen := len(inSlice)
	if batchLen < 1 {
		return nil, errors.New("batch len can't be less than 1")
	}

	if inSliceLen < batchLen {
		return nil, errors.New("len of the slice can't be less than batch len")
	}

	regularBatchesNumber := inSliceLen / batchLen

	batchNumber := int(math.Ceil(float64(inSliceLen) / float64(batchLen)))
	outSlice := make([][]int, 0, batchNumber)

	for i := 0; i < regularBatchesNumber; i++ {
		offset := i * batchLen
		outSlice = append(outSlice, inSlice[offset:offset+batchLen])
	}

	if regularBatchesNumber != batchNumber {
		outSlice = append(outSlice, inSlice[regularBatchesNumber*batchLen:])
	}

	return outSlice, nil
}

func ReverseMap(inMap map[string]int) map[int]string {
	outMap := make(map[int]string)

	for k, v := range inMap {
		outMap[v] = k
	}

	return outMap
}

func FilterByHardcodedValues(inSlice []int) []int {
	hardcodedValues := map[int]struct{}{1: {}, 2: {}, 3: {}}
	outSlice := make([]int, 0)

	lenInSlice := len(inSlice)
	if lenInSlice == 0 {
		return outSlice
	}

	for _, v := range inSlice {
		if _, ok := hardcodedValues[v]; !ok {
			outSlice = append(outSlice, v)
		}
	}

	return outSlice
}
