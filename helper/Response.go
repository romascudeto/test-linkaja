package helper

func ResponseSuccess(meta interface{}, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status": "success",
		"data":   data,
	}
}

func ResponseError(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status": "failed",
		"error":  data,
	}
}
