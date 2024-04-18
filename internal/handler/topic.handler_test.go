package handler_test

import (
	"log"
	"simushop/internal/core"
	"simushop/internal/dto"
	"testing"
)

func TestRequiredFieldsCreateTopic(t *testing.T) {
	t.Log("testing Required Fields from create Topic.")

	res := createTopic(dto.CreateTopicDTO{})

	BadRequest(res, t)
}

func TestCreateTopicWithExistingName(t *testing.T) {
	t.Log("testing create Topic")

	res := createTopic(dto.CreateTopicDTO{
		TopicName: "topic 1",
	})

	BadRequest(res, t)
}

func TestCreateTopic(t *testing.T) {
	t.Log("testing creating topic.")

	res := createTopic(dto.CreateTopicDTO{
		TopicName: "topic to Create",
	})

	Created(res, t)
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

func TestUpdateTopic(t *testing.T) {
	t.Log("updating Topics")

	res := updateTopic(dto.UpdateTopicDTO{
		TopicName: "topic to Update",
	})

	Ok(res, t)
}

func TestUpdateTopicWithExistingName(t *testing.T) {
	t.Log("testing update topic with existing name")

	res := updateTopic(dto.UpdateTopicDTO{
		TopicName: "topic 1",
	})

	BadRequest(res, t)
}

func updateTopic(dtoToUpdate dto.UpdateTopicDTO) core.Response {
	log.Println("[UpdateTopic]")

	return handlerInstance.UpdateTopic(core.Request{
		Body: func(obj any) error {
			updateTopicDTO, ok := obj.(*dto.UpdateTopicDTO)

			if ok {
				updateTopicDTO.TopicName = dtoToUpdate.TopicName
			}
			return nil
		},
		GetParam: func(key string) string {
			return "1"
		},
	})
}
