package tests

import (
	"github.com/parada3desu/deepfinder.go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDictWith0Lvl(t *testing.T) {
	var data interface{} = map[string]interface{}{
		"value": 39,
	}
	result := deepfinder.DeepFind(data, "")
	assert.EqualValues(t, data, result)
}

func TestDictWith1Lvl(t *testing.T) {
	var data interface{} = map[string]interface{}{
		"value": 39,
	}
	result := deepfinder.DeepFind(data, "value")
	assert.EqualValues(t, result, 39)
}

func TestDictWith2Lvl(t *testing.T) {
	var data interface{} = map[string]interface{}{
		"value": map[string]interface{}{
			"subdata": 39,
		},
	}
	result := deepfinder.DeepFind(data, "value.subdata")
	assert.EqualValues(t, result, 39)
}
