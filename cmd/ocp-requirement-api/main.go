package main

import (
	"fmt"
	"github.com/ozoncp/ocp-requirement-api/internal/utils"
)

func main() {
	fmt.Println("project name: ocp-requirement-api")

	fmt.Println("Utils usage example:")

	example1 := []int{1, 2, 3, 4}
	// [] batch len can't be less than 1
	fmt.Println(utils.MakeSliceOfBatches(example1, -1))
	// [[1 2] [3 4]] <nil>
	fmt.Println(utils.MakeSliceOfBatches(example1, 2))
	// [] len of the slice can't be less than batch len
	fmt.Println(utils.MakeSliceOfBatches(example1, 5))
	// [[1 2 3] [4]] <nil>
	fmt.Println(utils.MakeSliceOfBatches(example1, 3))

	example2 := map[string]int{"test": 1, "test2": 2}
	// map[1:test 2:test2]
	fmt.Println(utils.ReverseMap(example2))
	example3 := map[string]int{}
	// map[]
	fmt.Println(utils.ReverseMap(example3))

	example5 := []int{1, 2, 3, 4}
	// [4]
	fmt.Println(utils.FilterByHardcodedValues(example5))
	example6 := []int{9, 9, 9, 9, 9, 9, 9}
	// [9 9 9 9 9 9 9]
	fmt.Println(utils.FilterByHardcodedValues(example6))
	var example7 []int
	// []
	fmt.Println(utils.FilterByHardcodedValues(example7))
	example8 := []int{9}
	// [9]
	fmt.Println(utils.FilterByHardcodedValues(example8))
}
