package dto

type CreateTopicDTO struct {
	TopicName string `json:"topicName" validate:"required"`
}
