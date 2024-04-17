package handler

import (
	"simushop/internal/core"
	"simushop/internal/dto"
)

func (h handler) CreateTopic(request core.Request) core.Response {

	var (
		body   dto.CreateTopicDTO
		result = core.Ok("Topic Created")
	)

	request.Body(&body)

	if err := h.ValidateStruct(body); err != nil {
		result = core.BadRequest(err)
	}

	return result
}
