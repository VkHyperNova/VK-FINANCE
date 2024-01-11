package util

/* Error Handling */
func HandleError(err error) bool {
	if err != nil {
		PrintRed(err.Error() + "\n")
		return true
	}

	return false
}