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

func TestExistingPathInDictInList(t *testing.T) {
	data := map[string]interface{}{"values": []interface{}{
		map[string]interface{}{"nya": "a"},
		map[string]interface{}{"value": "b"},
		map[string]interface{}{"nya": "c"},
	}}
	result := deepfinder.DeepFind(data, "values.?.value")
	assert.EqualValues(t, "b", result)
}

func TestAllValuesOfAListInList(t *testing.T) {
	data := map[string]interface{}{"all-values": []interface{}{
		map[string]interface{}{"values": []interface{}{"a"}},
		map[string]interface{}{"values": []interface{}{"b", "c", "d"}},
		map[string]interface{}{"values": []interface{}{"e"}},
	}}
	result := deepfinder.DeepFind(data, "all-values.*.values")
	assert.EqualValues(t, []interface{}{[]interface{}{"a"}, []interface{}{"b", "c", "d"}, []interface{}{"e"}}, result)
}

func TestCompleteListInList(t *testing.T) {
	data := map[string]interface{}{
		"all-values": []interface{}{[]interface{}{"a"}, []interface{}{"b", "c", "d"}, []interface{}{"e"}},
	}
	result := deepfinder.DeepFind(data, "all-values.*")
	assert.EqualValues(t, []interface{}{[]interface{}{"a"}, []interface{}{"b", "c", "d"}, []interface{}{"e"}}, result)
}

func TestCompleteListInListInTwoSteps(t *testing.T) {
	data := map[string]interface{}{
		"all-values": []interface{}{[]interface{}{"a"}, []interface{}{"b", "c", "d"}, []interface{}{"e"}},
	}
	result := deepfinder.DeepFind(data, "all-values.*.*")
	assert.EqualValues(t, []interface{}{[]interface{}{"a"}, []interface{}{"b", "c", "d"}, []interface{}{"e"}}, result)
}

func TestExistingPathOfAListInsideList(t *testing.T) {
	data := map[string]interface{}{
		"all-values": []interface{}{[]interface{}{"a"}, []interface{}{"b", "c", "d"}, []interface{}{"e"}},
	}
	result := deepfinder.DeepFind(data, "all-values.?.2")
	assert.EqualValues(t, "d", result)
}

func TestExistingPathInsideExistingPath(t *testing.T) {
	data := map[string]interface{}{
		"all-values": []interface{}{
			[]interface{}{"a"},
			[]interface{}{
				"b",
				map[string]interface{}{"correct": "correct"},
				"d",
			},
			[]interface{}{"e"},
		},
	}
	result := deepfinder.DeepFind(data, "all-values.?.?.correct")
	assert.EqualValues(t, "correct", result)
}
