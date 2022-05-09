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

func TestSecondValueOfList(t *testing.T) {
	data := []interface{}{"a", "b", "c"}
	result := deepfinder.DeepFind(data, "1")
	assert.EqualValues(t, "b", result)
}

func TestLastValueOfList(t *testing.T) {
	data := []interface{}{"a", "b", "c"}
	result := deepfinder.DeepFind(data, "2")
	assert.EqualValues(t, "c", result)
}

func TestNonNumericAccessOfList(t *testing.T) {
	data := []interface{}{"a", "b", "c"}
	result := deepfinder.DeepFind(data, "b")
	assert.EqualValues(t, nil, result)
}

func TestNonHasheableKeyAccessOfList(t *testing.T) {
	data := []interface{}{"a", "b", "c"}
	result := deepfinder.DeepFind(data, "[]")
	assert.EqualValues(t, nil, result)
}

func TestEmbeddedList(t *testing.T) {
	data := map[string]interface{}{"values": []interface{}{"a", "b", "c"}}
	result := deepfinder.DeepFind(data, "values.1")
	assert.EqualValues(t, "b", result)
}

func TestFirstValueInDictInList(t *testing.T) {
	data := map[string]interface{}{"values": []interface{}{
		map[string]interface{}{"value": "a"},
		map[string]interface{}{"value": "b"},
		map[string]interface{}{"value": "c"},
	}}
	result := deepfinder.DeepFind(data, "values.0.value")
	assert.EqualValues(t, "a", result)
}

func TestAllValuesInDictInList(t *testing.T) {
	data := map[string]interface{}{"values": []interface{}{
		map[string]interface{}{"value": "a"},
		map[string]interface{}{"value": "b"},
		map[string]interface{}{"value": "c"},
	}}
	result := deepfinder.DeepFind(data, "values.*.value")
	assert.EqualValues(t, []interface{}{"a", "b", "c"}, result)
}
