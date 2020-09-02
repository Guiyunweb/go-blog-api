package service

import (
	"blog-api/cache"
	"blog-api/model"
	"blog-api/serializer"
	"blog-api/util"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

type UserRegisterService struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	About    string `json:"about"`
}

type UserLoginService struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Valid 验证表单
func (service *UserRegisterService) Valid() *serializer.Response {

	count, err := model.DB.Where("username = ?", service.Username).Count(&model.User{})
	if err != nil {
		return &serializer.Response{
			Success: false,
			Message: "服务器繁忙",
		}
	}
	if count > 0 {
		return &serializer.Response{
			Success: false,
			Message: "用户名已被注册",
		}
	}

	return nil
}

// 注册方法
func (service *UserRegisterService) Register() (model.User, *serializer.Response) {
	user := model.User{
		Id:       util.GetSnowflakeId(),
		Username: service.Username,
		Email:    service.Email,
		About:    service.About,
	}
	// 表单验证
	if err := service.Valid(); err != nil {
		return user, err
	}

	if err := user.SetPassword(service.Password); err != nil {
		return user, &serializer.Response{
			Success: false,
			Message: "密码加密失败",
		}
	}

	// 创建用户
	if _, err := model.DB.Insert(&user); err != nil {
		return user, &serializer.Response{
			Success: false,
			Message: "注册失败",
		}
	}
	return user, nil
}

// 登录方法
func (service UserLoginService) Login() (model.User, *serializer.Response) {
	var user model.User
	if _, err := model.DB.Where("username = ?", service.Username).Get(&user); err != nil {
		return user, &serializer.Response{
			Success: false,
			Message: "账户名错误",
		}
	}

	if user.CheckPassword(service.Password) == false {
		return user, &serializer.Response{
			Success: false,
			Message: "密码错误",
		}
	}

	token, _ := util.CreateToken(user.Id)
	user.Token = token

	// 将用户信息保持至redis
	if userString, err := json.Marshal(user); err == nil {
		_, err := cache.RedisClient.Do("SET", user.Id, userString, "EX", "5000")
		if err != nil {
			fmt.Println(err)
		}
	} else {
		return user, &serializer.Response{
			Success: false,
			Message: "登录失败",
		}
	}

	return user, nil
}

func UserInfo(token string) (model.User, *serializer.Response) {
	var user model.User
	if userId, err := util.VerifyToken(token); err == nil {
		values, _ := redis.Bytes(cache.RedisClient.Do("GET", userId))
		json.Unmarshal(values, &user)
		return user, nil
	} else {
		return user, &serializer.Response{
			Success: false,
			Message: "用户信息获取失败",
		}
	}

}
