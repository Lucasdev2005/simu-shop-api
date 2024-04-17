package repository

import (
	"simushop/internal/entity"
)

func (r repository) CreateTopic(topic entity.Topic) error {
	var result error = nil

	err := findExistingTopicByName(topic, r)

	if err != nil {
		return err
	} else {
		result = r.db.Create(&topic).Error
	}

	return result
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

func (r repository) ListTopics(where string, args ...interface{}) []entity.Topic {
	var (
		topics []entity.Topic
		limit  = 10
	)

	if len(where) > 0 {
		r.queryWhere(where, args).Limit(limit).Find(&topics)
	} else {
		r.db.Find(&topics).Limit(limit)
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
