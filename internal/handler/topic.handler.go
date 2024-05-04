package handler

import (
	"simushop/internal/core"
	"simushop/internal/dto"
	"simushop/internal/entity"
)

func (h handler) CreateTopic(request core.Request) core.Response {

	var (
		body dto.CreateTopicDTO
	)

	request.Body(&body)

	if err := h.ValidateStruct(body); err != nil {
		return core.BadRequest(err)
	}

	topic, err := h.repository.CreateTopic(body.ToTopic())

	if err != nil {
		return core.BadRequest(err.Error())
	}

	return core.Created(topic)
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

func (h handler) ListTopics(request core.Request) core.Response {
	var (
		topics    []entity.Topic
		result    = core.InternalError("error on list topics.")
		topicName = request.GetQueryParam("topic_name")
		query     = h.repository.ListTopics
		paginate  = core.NewPaginate(request)
	)

	if len(topicName) > 0 {
		topics = query(paginate, "topic_name LIKE ?", "%"+topicName+"%")
	} else {
		topics = query(paginate, "")
	}

	result = core.Ok(topics)

	return result
}
