package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestString(t *testing.T) {
	requirement := Requirement{Id: 1, UserId: 1, Text: "whatever"}
	assert.Equal(t, requirement.String(), "{Id:1, UserId:1, Text:\"whatever\"}")
}
