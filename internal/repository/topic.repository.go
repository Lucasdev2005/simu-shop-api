package repository

import (
	"simushop/internal/core"
	"simushop/internal/entity"
)

func (r repository) CreateTopic(topic entity.Topic) (entity.Topic, error) {
	err := findExistingTopicByName(topic, r)

	if err != nil {
		return entity.Topic{}, err
	} else {
		errOnCreate := r.db.Create(&topic).Error
		if errOnCreate != nil {
			return entity.Topic{}, errOnCreate
		}
	}

	return topic, nil
}

func (r repository) UpdateTopic(topic entity.Topic, where string, args ...interface{}) error {
	var (
		applyUpdate bool  = true
		result      error = nil
	)

	if len(topic.TopicName) > 0 {
		err := findExistingTopicByName(topic, r)
		if err != nil {
			applyUpdate = false
			result = err
		}
	}

	if applyUpdate {
		result = r.update(&topic, where, args...).Error
	}

	return result
}

func (r repository) ListTopics(paginate core.Paginate, where string, args ...interface{}) []entity.Topic {
	var (
		topics []entity.Topic
	)

	if len(where) > 0 {
		r.paginate(paginate).Where(where, args...).Find(&topics)
	} else {
		r.paginate(paginate).Find(&topics)
	}

	return topics
}

func findExistingTopicByName(topic entity.Topic, r repository) error {
	r.queryFirst(&topic, "topic_name = ?", topic.TopicName)

	if topic.Exist() {
		return newError("topic with name " + topic.TopicName + " exists.")
	}

	return nil
}
