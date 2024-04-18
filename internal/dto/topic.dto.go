package dto

import "simushop/internal/entity"

type CreateTopicDTO struct {
	TopicName string `json:"topicName" validate:"required"`
}

type UpdateTopicDTO struct {
	TopicName string `json:"topicName" validate:"omitempty,min=1"`
}

func (c CreateTopicDTO) ToTopic() entity.Topic {
	return entity.Topic{
		TopicName: c.TopicName,
	}
}

func (u UpdateTopicDTO) ToTopic() entity.Topic {
	return entity.Topic{
		TopicName: u.TopicName,
	}
}
