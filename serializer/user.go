package serializer

import (
	"blog-api/model"
)

// User 用户序列化器
type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Status   uint   `json:"status"`
	About    string `json:"about"`
	Token    string `json:"token"`
}

// UserResponse 单个用户序列化
type UserResponse struct {
	Response
	Success bool `json:"success"`
	Data    User `json:"data"`
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		ID:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		About:    user.About,
		Status:   user.Status,
		Token:    user.Token,
	}
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user model.User) UserResponse {
	return UserResponse{
		Success: true,
		Data:    BuildUser(user),
	}
}

func CheckLogin() Response {
	return Response{
		Success: false,
		Message: "账户未登录",
	}
}
