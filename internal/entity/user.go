package entity

type User struct {
	UserId int `gorm:"primaryKey"`
}

func (User) TableName() string {
	return "user"
}
