package dto

func ResponseFailed(msg string, code int) map[string]interface{} {
	return map[string]interface{}{
		"status":      "error",
		"status_code": code,
		"message":     msg,
	}
}

func ResponseSuccesNoData(msg string) map[string]interface{} {
	return map[string]interface{}{
		"status":  "success",
		"message": msg,
	}
}

func ResponseSuccesWithData(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":  "success",
		"message": msg,
		"data":    data,
	}
}

var FailedMessage = map[string]interface{}{
	"binding":   "Failed to insert(binding) data, check your input",
	"parseInt":  "Invalid ID, make sure you enter the correct ID",
	"transform": "Failed to transform data into response payload",
	"mail":      "Error sending email",
}

type PayloadID struct {
	Id int `json:"id"`
}
