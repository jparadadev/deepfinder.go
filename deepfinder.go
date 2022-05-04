package deepfinder

import (
	"reflect"
	"strings"
)

func DeepFind(obj interface{}, path string) interface{} {
	steps := strings.Split(path, ".")
	if len(steps) == 1 && steps[0] == "" {
		steps = []string{}
	}

	res := recHelper(obj, steps)
	return res
}

func recHelper(obj interface{}, path []string) interface{} {

	if obj == nil {
		return nil
	}

	if len(path) == 0 {
		return obj
	}

	currentStep, remainingPath := path[0], path[1:]

	if reflect.TypeOf(obj).Kind().String() == "map" {
		parsedObj := obj.(map[string]interface{})
		return recHelper(parsedObj[currentStep], remainingPath)
	}
	return nil
}
