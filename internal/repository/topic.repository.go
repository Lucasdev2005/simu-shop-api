package repository

import (
	"log"
	"simushop/internal/core"
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

func (r repository) ListTopics(paginate core.Paginate, where string, args ...interface{}) []entity.Topic {
	var (
		topics []entity.Topic
		limit  = 10
	)

	log.Println("Repository [ListTopics] paginate", paginate)
	if len(where) > 0 {
		err := r.paginate(paginate, &topics).Limit(limit).Where(where, args...)

		log.Println("Repository [ListTopics] error ", err.Error)
	} else {
		r.paginate(paginate, &topics).Limit(limit)
	}

	log.Println("[ListTopics] topics", topics)
	return topics
}

func findExistingTopicByName(topic entity.Topic, r repository) error {
	r.queryFirst(&topic, "topic_name = ?", topic.TopicName)

	if topic.Exist() {
		return newError("topic with name " + topic.TopicName + " exists.")
	}

	return nil
}
