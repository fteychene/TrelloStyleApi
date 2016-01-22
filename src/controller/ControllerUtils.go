package controller

import (
	"fmt"
)

func checkParameters(parameters map[string][]string, parametersName []string) (bool, string) {
	parametersInError := []string{}
	for index := range parametersName {
		if parameters[parametersName[index]] == nil || len(parameters[parametersName[index]]) == 0 {
			parametersInError = append(parametersInError, parametersName[index])
		}
	}
	if len(parametersInError) > 0 {
		return false, fmt.Sprintf("Illegal argument. Need %v", parametersName)
	}
	return true, "Ok"
}