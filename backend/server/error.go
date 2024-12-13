package server

func ErrorResponse(err error, message string) map[string]interface{} {
	return map[string]interface{}{
		"error":   err.Error(),
		"message": message,
	}
}
