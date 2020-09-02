package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id       int64
	Username string
	Password string
	Email    string
	Status   uint
	About    string
	Token    string `xorm:"-"`
}

// PassWordCost 密码加密难度
const PassWordCost = 12

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
