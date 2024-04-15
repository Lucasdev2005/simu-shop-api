package entity

// TODO Topics and many2many with products
type Topic struct {
	TopicId   int    `gorm:"column:topic_id" json:"topicId"`
	TopicName string `gorm:"column:topic_name;uniqueIndex" json:"topicName"`
}
