package middleware

import (
	"GoAsk/serializer"
	"GoAsk/utils"
	e "GoAsk/utils/error"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"net/http"
)

// JwtAuth 用于验证JWT是否正确的中间件
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 通过cookies读取jwt token
		//token := c.GetHeader("Cookie")
		token := c.GetHeader("Authorization")
		if token == "" { // token缺失
			c.JSON(http.StatusUnauthorized, serializer.BuildResponse(e.ErrorTokenMissing, nil))
			c.Abort()
			return
		}
		claims, err := utils.ParseToken(token)
		if err != nil { // token解析错误
			if errors.Is(err, jwt.ErrTokenExpired) { // token过期
				c.JSON(http.StatusForbidden, serializer.BuildResponse(e.ErrorTokenExpired, nil))
			} else { // 其他错误
				c.JSON(http.StatusForbidden, serializer.BuildResponse(e.ErrorTokenParsing, nil))
			}
			c.Abort()
			return
		}
		c.Set("user_id", claims.UserId) // 通过Set和Get来在中间件中传递参数
		c.Set("user_name", claims.UserName)
	}
}
