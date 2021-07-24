package utils

import (
	"github.com/ozoncp/ocp-requirement-api/internal/models"
	test_utils "github.com/ozoncp/ocp-requirement-api/internal/test-utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeSliceOfBatchesSuccess(t *testing.T) {

	example := []int{1, 2, 3, 4}

	result, err := MakeSliceOfBatches(example, 2)
	assert.NoError(t, err)
	assert.Equal(t, [][]int{{1, 2}, {3, 4}}, result)

	result, err = MakeSliceOfBatches(example, 3)
	assert.NoError(t, err)
	assert.Equal(t, [][]int{{1, 2, 3}, {4}}, result)
}

func TestMakeSliceOfBatchesError(t *testing.T) {
	example := []int{1, 2, 3, 4}

	_, err := MakeSliceOfBatches(example, -1)
	assert.EqualError(t, err, "batch len can't be less than 1")

	_, err = MakeSliceOfBatches(example, 5)
	assert.EqualError(t, err, "len of the slice can't be less than batch len")
}

func TestReverseMap(t *testing.T) {
	example := map[string]int{"test": 1, "test2": 2}

	result := ReverseMap(example)
	assert.Equal(t, map[int]string{1: "test", 2: "test2"}, result)

	result = ReverseMap(map[string]int{})
	assert.Equal(t, map[int]string{}, result)
}

func TestFilterByHardcodedValues(t *testing.T) {
	example := []int{1, 2, 3, 4}
	result := FilterByHardcodedValues(example)

	assert.Equal(t, []int{4}, result)

	example = []int{9, 9, 9, 9, 9, 9, 9}
	result = FilterByHardcodedValues(example)

	assert.Equal(t, []int{9, 9, 9, 9, 9, 9, 9}, result)

	example = []int{}
	result = FilterByHardcodedValues(example)

	assert.Equal(t, []int{}, result)

	example = []int{9}
	result = FilterByHardcodedValues(example)

	assert.Equal(t, []int{9}, result)
}
func TestSplitToBulksSuccess(t *testing.T) {

	example := []models.Requirement{
		{1, 1, ""},
		{2, 2, ""},
		{3, 3, ""},
		{4, 4, ""},
	}

	result, err := SplitRequirementsToBulks(example, 2)
	assert.NoError(t, err)
	assert.Equal(
		t,
		[][]models.Requirement{
			{
				{1, 1, ""},
				{2, 2, ""},
			},
			{
				{3, 3, ""},
				{4, 4, ""},
			},
		},
		result,
	)

	result, err = SplitRequirementsToBulks(example, 3)
	assert.NoError(t, err)
	assert.Equal(
		t,
		[][]models.Requirement{
			{
				{1, 1, ""},
				{2, 2, ""},
				{3, 3, ""},
			},
			{
				{4, 4, ""},
			},
		},
		result,
	)
}

func TestSplitToBulksError(t *testing.T) {
	example := []models.Requirement{
		{1, 1, ""},
		{2, 2, ""},
		{3, 3, ""},
		{4, 4, ""},
	}

	_, err := SplitRequirementsToBulks(example, -1)
	assert.EqualError(t, err, "batch len can't be less than 1")

	_, err = SplitRequirementsToBulks(example, 5)
	assert.EqualError(t, err, "len of the slice can't be less than batch len")
}

func TestConvertRequirementsSliceToMap(t *testing.T) {
	example := []models.Requirement{
		{1, 1, ""},
		{2, 2, ""},
		{3, 3, ""},
		{4, 4, ""},
	}

	result := ConvertRequirementsSliceToMap(example)
	assert.Equal(
		t,
		map[uint64]models.Requirement{
			1: {1, 1, ""},
			2: {2, 2, ""},
			3: {3, 3, ""},
			4: {4, 4, ""},
		},
		result,
	)

	result = ConvertRequirementsSliceToMap([]models.Requirement{})
	assert.Equal(t, result, map[uint64]models.Requirement{})
}

func TestReadConfigFileSuccess(t *testing.T) {
	path, cleanUp := test_utils.CreateTmpFile("config.txt", []byte("line1\nline2"))
	result, err := ReadConfigFile(path)

	assert.NoError(t, err)
	assert.Equal(t, result, []string{"line1", "line2"})

	defer cleanUp()
}

func TestReadConfigFileFailed(t *testing.T) {
	_, err := ReadConfigFile("whatever")
	assert.EqualError(t, err, "open whatever: no such file or directory")
}
