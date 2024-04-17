package handler_test

import (
	"simushop/internal/core"
	"simushop/internal/dto"
	"testing"
)

func TestRequiredFieldsCreateTopic(t *testing.T) {
	t.Log("testing Required Fields from create Topic.")

	res := createTopic(dto.CreateTopicDTO{})

	BadRequest(res, t)
}

func createTopic(data dto.CreateTopicDTO) core.Response {
	return handlerInstance.CreateTopic(core.Request{
		Body: func(obj any) error {
			createTopicDTO, ok := obj.(*dto.CreateTopicDTO)

			if ok {
				createTopicDTO.TopicName = data.TopicName
			}
			return nil
		},
	})
}
