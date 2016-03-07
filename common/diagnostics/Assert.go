package diagnostics

import "strings"

func StringIsNotEmpty(value, valueName, message string) string {
	if len(value) == 0 {
		panic(getMessage(valueName, message))
	}

	return value
}

// Returns the message to be displayed when calling different Assert functions.
// valueName contains the name of value the Assert was called on.
// message contains the message to be displayed if the Assert failed.
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
