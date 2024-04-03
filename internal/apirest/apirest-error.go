package apirest

import (
	"net/http"
	"simushop/internal/core"
)

func BadRequest(response interface{}) core.Fail {
	return core.Fail{
		ErrorCode: http.StatusBadRequest,
		Response:  response,
	}
}
