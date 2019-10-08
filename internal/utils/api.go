package utils

func ParseError(err error) int {
	switch err.Error() {
	case "NOT_FOUND":
		return 404
	default:
		return 500
	}
}