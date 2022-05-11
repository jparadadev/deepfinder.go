package tests

import (
	"github.com/parada3desu/deepfinder.go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBasicNumber(t *testing.T) {
	var data interface{} = 39
	result := deepfinder.DeepFind(data, "")
	assert.EqualValues(t, 39, result)
}

func TestNumberInDict(t *testing.T) {
	var data interface{} = map[string]interface{}{
		"value": 39,
	}
	result := deepfinder.DeepFind(data, "value")
	assert.EqualValues(t, 39, result)
}

func TestBasicString(t *testing.T) {
	var data interface{} = "test str"
	result := deepfinder.DeepFind(data, "")
	assert.EqualValues(t, "test str", result)
}

func TestStringInDict(t *testing.T) {
	var data interface{} = map[string]interface{}{
		"value": "test str",
	}
	result := deepfinder.DeepFind(data, "value")
	assert.EqualValues(t, "test str", result)
}
