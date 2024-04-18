package dto

import "simushop/internal/entity"

type CreateTopicDTO struct {
	TopicName string `json:"topicName" validate:"required"`
}

func (c CreateTopicDTO) ToTopic() entity.Topic {
	return entity.Topic{
		TopicName: c.TopicName,
	}
}
