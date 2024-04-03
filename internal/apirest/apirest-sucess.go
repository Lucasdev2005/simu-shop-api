package apirest

import (
	"net/http"
	"simushop/internal/core"
)

func Created(data interface{}) core.Success {
	return core.Success{
		SuccessCode: http.StatusCreated,
		Response:    data,
	}
}
