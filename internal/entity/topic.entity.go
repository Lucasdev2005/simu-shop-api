package entity

// TODO Topics and many2many with products
type Topic struct {
	TopicId   int       `gorm:"column:topic_id;primaryKey" json:"topicId"`
	TopicName string    `gorm:"column:topic_name;uniqueIndex" json:"topicName"`
	Products  []Product `gorm:"many2many:product_topic"`
}

func (Topic) TableName() string {
	return "topic"
}

func (t Topic) Exist() bool {
	return t.TopicId != 0
}
