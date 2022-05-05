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
