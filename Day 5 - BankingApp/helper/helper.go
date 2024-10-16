package helper

import "strings"

//used to check output of all Validation Functions and return the first Error Found. returns Nil if All Validations are asserted
func ValidateAll(validationFunctionOutputs ...error) error {

	for _, error := range validationFunctionOutputs {
		if error != nil {
			return error
		}
	}
	return nil
}

func RemoveAllLeadingAndTrailingSpaces(s string) string {
	return strings.TrimSpace(s)
}
