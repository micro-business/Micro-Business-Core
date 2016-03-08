// Contains function to simplify checking pre-conditions required before executing an operation
package diagnostics

import "strings"

// Makes sure provided value is not an empty string, otherwise it panics
// value: the value to be checked
// valueName: Optional. Contains the name of the value the Assert was called on
// message: Optional. Contains the message to be displated if the Assert failed.
func StringIsNotEmpty(value, valueName, message string) string {
	if len(value) == 0 {
		panic(getMessage(valueName, message))
	}

	return value
}

// Makes sure provided value is not an empty string or only contains whitespace, otherwise it panics
// value: the value to be checked
// valueName: Optional. Contains the name of the value the Assert was called on
// message: Optional. Contains the message to be displated if the Assert failed.
func StringIsNotEmptyOrWhitespace(value, valueName, message string) string {
	StringIsNotEmpty(value, valueName, message)

	if len(strings.TrimSpace(value)) == 0 {
		panic(getMessage(valueName, message))
	}

	return value
}

// Returns the message to be displayed when calling different Assert functions.
// valueName: Optional. Contains the name of value the Assert was called on.
// message: Optional. Contains the message to be displayed if the Assert failed.
//
// Remarks: Both valueName and message will be trimmed before running the checks
func getMessage(valueName, message string) string {
	valueName = strings.TrimSpace(valueName)
	message = strings.TrimSpace(message)

	if len(valueName) != 0 && len(message) != 0 {
		return message
	} else if len(valueName) == 0 && len(message) != 0 {
		return message
	} else if len(valueName) != 0 && len(message) == 0 {
		return valueName
	} else {
		return ""
	}
}
