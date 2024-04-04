package core

import (
	"net/http"
)

func BadRequest(data interface{}) Response {
	return Response{
		Code: http.StatusBadRequest,
		Data: data,
	}
}
