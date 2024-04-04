package entity

type User struct {
	UserId       int     `gorm:"column:user_id;primaryKey"`
	UserPassword string  `gorm:"column:user_password;type:varchar"`
	Username     string  `gorm:"column:username;type:varchar(50);uniqueIndex"`
	UserBalance  float64 `gorm:"column:user_ballance"`
	UserType     string  `gorm:"column:user_type"`
}

func (User) TableName() string {
	return "user"
}

func (u User) Exist() bool {
	return u.UserId != 0
}
