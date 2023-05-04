package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var JwtSecret = []byte("lykSignHere")

type Claims struct {
	jwt.RegisteredClaims
	UserId   uint   `json:"user_id"`
	UserName string `json:"user_name"`
}

// GenerateToken 签发Token
func GenerateToken(userID uint, username string) (tokenString string, err error) { // 错误抛给上级处理
	claims := Claims{
		UserId:   userID,
		UserName: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(48 * time.Hour)), // 调试的时候过期时间可以设久点
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "GoAsk OAuth2 Server", // 其实没有专门的OAuth2服务器
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // 用HS256进行对称加密
	tokenString, err = token.SignedString(JwtSecret)           // 用密匙签发
	return
}

// ParseToken 验证用户token
func ParseToken(tokenString string) (claims *Claims, err error) { // err=nil时表示验证成功，错误抛给上级处理
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})
	if err != nil {
		Logger.Infoln(err)
	}
	if token != nil {
		if claims, ok := token.Claims.(*Claims); ok && token.Valid { // 通过这种语法，可以把token的Claims指向的claims取出来
			return claims, nil
		}
	}
	return nil, err
}
