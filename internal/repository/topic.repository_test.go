package repository_test

import (
	"simushop/internal/core"
	"simushop/internal/entity"
	"strconv"
	"testing"
)

func TestCreatingTopic(t *testing.T) {
	t.Log("testing create Topic.")

	err := createTopic(entity.Topic{
		TopicName: "Topic Test create",
	})

	equal(t, nil, err, "error on create a topic.")
}

func TestUpdateTopic(t *testing.T) {
	t.Log("testing update topic")

	topicName := "new name"

	createTopic(entity.Topic{TopicName: topicName})

	err := mockRepository.UpdateTopic(entity.Topic{TopicName: "new name actual"}, "topic_name = ?", topicName)

	equal(t, nil, err, "error on update a product")
}

func TestCreatingTopicsWithSameName(t *testing.T) {
	t.Log("testing create topics with same names.")

	topic := entity.Topic{
		TopicName: "Topic Test",
	}

	createTopic(topic)
	err := createTopic(topic)

	notEqual(t, nil, err, "creating topics with same name.")
}

func TestListTopics(t *testing.T) {
	t.Log("testing list Topics")

	for _, i := range []int{1, 2, 3} {
		createTopic(entity.Topic{
			TopicName: "topic " + strconv.Itoa(i),
		})
	}

	topics := mockRepository.ListTopics(core.NewPaginate(core.Request{
		GetQueryParam: func(key string) string {
			if key == "page" {
				return "1"
			} else {
				return "2"
			}
		},
	}), "topic_name LIKE ? ", "%"+"topic"+"%")

	t.Error("[]", topics)
	equal(t, true, len(topics) > 0, "error on List Topics.")
}

func createTopic(t entity.Topic) error {
	return mockRepository.CreateTopic(t)
}
