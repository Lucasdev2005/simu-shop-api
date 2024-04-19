package core

import "net/http"

func Created(data interface{}) Response {
	return Response{
		Code: http.StatusCreated,
		Data: data,
	}
}

func Ok(data interface{}) Response {
	return Response{
		Code: http.StatusOK,
		Data: data,
	}
}

func InternalError(data interface{}) Response {
	return Response{
		Code: http.StatusInternalServerError,
		Data: data,
	}
}
