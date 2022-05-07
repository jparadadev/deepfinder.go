package tests

import (
	"github.com/parada3desu/deepfinder.go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFirstValueOfList(t *testing.T) {
	data := []interface{}{"a", "b", "c"}
	result := deepfinder.DeepFind(data, "0")
	assert.EqualValues(t, "a", result)
}
