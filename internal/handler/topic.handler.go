package handler

import (
	"simushop/internal/core"
	"simushop/internal/dto"
)

func (h handler) CreateTopic(request core.Request) core.Response {

	var (
		body   dto.CreateTopicDTO
		result = core.Created("Topic Created")
	)

	request.Body(&body)

	if err := h.ValidateStruct(body); err != nil {
		result = core.BadRequest(err)
	}

	err := h.repository.CreateTopic(body.ToTopic())

	if err != nil {
		result = core.BadRequest(err.Error())
	}

	return result
}

func (h handler) UpdateTopic(request core.Request) core.Response {

	var (
		body   dto.UpdateTopicDTO
		result = core.Ok("topic updated")
	)

	request.Body(&body)

	if err := h.ValidateStruct(body); err != nil {
		result = core.BadRequest(err)
	}

	err := h.repository.UpdateTopic(body.ToTopic(), "topic_id = ?", request.GetParam("id"))

	if err != nil {
		result = core.BadRequest(err.Error())
	}

	return result
}
