package deepfinder

import (
	"reflect"
	"strconv"
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
		dict := obj.(map[string]interface{})
		return recHelper(dict[currentStep], remainingPath)
	}

	if reflect.TypeOf(obj).Kind().String() == "list" {
		list := obj.([]interface{})
		if currentStep == "*" {
			var res []interface{} // TODO: optimize this
			for _, subObj := range list {
				res = append(res, recHelper(subObj, remainingPath))
			}
			return res
		}

		if currentStep == "?*" || currentStep == "*?" {
			var res []interface{}
			for _, subObj := range list {
				subRes := recHelper(subObj, remainingPath)
				if subRes != nil {
					res = append(res, subRes)
				}
			}
			return res
		}

		if currentStep == "?" {
			var res []interface{}
			for _, subObj := range list {
				subRes := recHelper(subObj, remainingPath)
				if subRes != nil {
					return subRes
				}
			}
			return res
		}

		numericCurrentStep, err := strconv.Atoi(currentStep)
		if err != nil {
			return nil
		}

		if numericCurrentStep >= len(list) {
			return nil
		}

		return recHelper(list[numericCurrentStep], remainingPath)
	}
	return nil
}
