package helper

func ResponseSuccess(meta interface{}, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"meta": meta,
		"data": data,
	}
}

func ResponseError(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"error": data,
	}
}
