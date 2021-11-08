package response

import "net/http"

type ServiceResponse struct {
	Code   int `json:"code"`
	Message string `json:"message"`
}

func SetResponse(Code int, Message string) ServiceResponse {
	var sr ServiceResponse
	sr.Code = Code
	if Message == ""{
		sr.Message = http.StatusText(Code)
	}else {
		sr.Message = Message
	}

	return sr
}