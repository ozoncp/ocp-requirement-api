package utils

import (
	"bufio"
	"errors"
	"log"
	"math"
	"os"
)

func MakeSliceOfBatches(inSlice []int, batchLen int) ([][]int, error) {
	inSliceLen := len(inSlice)
	if batchLen < 1 {
		return nil, errors.New("batch len can't be less than 1")
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

func ReadConfigFile(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal("Failed to close the file.")
		}
	}(file)

	scanner := bufio.NewScanner(file)

	result := make([]string, 0)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
