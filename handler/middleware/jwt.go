package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/marshhu/file-store-server/handler/resp"
	"github.com/marshhu/file-store-server/util"
	"net/http"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Token")
		if len(token) <= 0{
			c.JSON(http.StatusBadRequest, resp.Response{
				Code: resp.Unauthorized,
				Msg:  "未经授权，无访问权限",
				Data: nil,
			})
			c.Abort()
			return
		}

		_, err := util.ParseToken(token)
		if err != nil{
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				c.JSON(http.StatusUnauthorized, resp.Response{
					Code: resp.Unauthorized,
					Msg:  "token已过期",
					Data: nil,
				})
				c.Abort()
				return
			default:
				c.JSON(http.StatusUnauthorized, resp.Response{
					Code: resp.Unauthorized,
					Msg:  "token验证失败",
					Data: nil,
				})
				c.Abort()
				return
			}
		}
		c.Next()
	}
}