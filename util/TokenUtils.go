package util

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

const TOKEN string = "e7b1c4c689eda0e589f6e9a1013b312e"

// 生成Token
func CreateToken(id int64) (string, error) {

	stringId := strconv.FormatInt(id, 10)

	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Duration(1000) * time.Second).Unix(), // 过期时间
		Issuer:    "SERVICE",                                                // 签发人
		Audience:  stringId,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(TOKEN))
}

func VerifyToken(tokenString string) (interface{}, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("错误的签名方法: %v", token.Header["alg"])
		}

		return []byte(TOKEN), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["aud"], nil
	} else {
		fmt.Println(err)
		return claims["aud"], err
	}
}
