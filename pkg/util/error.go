package util

/* Error Handling */
func HandleError(err error) bool {
	if err != nil {
		PrintRedString(err.Error() + "\n")
		return true
	}

	return false
}
