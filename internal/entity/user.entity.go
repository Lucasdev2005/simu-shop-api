package entity

import (
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserId       int     `gorm:"column:user_id;primaryKey" json:"userId"`
	UserPassword string  `gorm:"column:user_password;type:varchar" json:"userPassword"`
	Username     string  `gorm:"column:username;type:varchar(50);uniqueIndex" json:"username"`
	UserBalance  float64 `gorm:"column:user_ballance" json:"userBalance"`
	UserType     string  `gorm:"column:user_type" json:"userType"`
}

func (User) TableName() string {
	return "user"
}

func (u User) Exist() bool {
	return u.UserId != 0
}

func (u *User) SetPassword() {
	hashSalts, _ := strconv.Atoi(os.Getenv("HASH_SALTS"))
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.UserPassword), hashSalts)
	if err != nil {
		panic(err.Error())
	} else {
		u.UserPassword = string(bytes)
	}
}

func (u User) ComparePassword(passwordNotEncrypt string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.UserPassword), []byte(passwordNotEncrypt))
	return err == nil
}
