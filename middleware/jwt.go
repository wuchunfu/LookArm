package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris/v12"
	"lookarm/config"
	"lookarm/utils/message"
	"strings"
	"time"
)

var JwtKey = []byte(config.JwtKey)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// 生成token
func SetToken(username string) (string, int) {
	expireTime := time.Now().Add(7 * 24 * time.Hour)
	SetClaims := MyClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "ginblog",
		},
	}

	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	token, err := reqClaim.SignedString(JwtKey)
	if err != nil {
		return "", message.ERROR
	}
	return token, message.SUCCSES

}

// 验证token

func CheckToken(token string) (*MyClaims, int) {
	var claims MyClaims

	setToken, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (i interface{}, e error) {
		return JwtKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, message.ERROR_TOKEN_WRONG
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return nil, message.ERROR_TOKEN_RUNTIME
			} else {
				return nil, message.ERROR_TOKEN_TYPE_WRONG
			}
		}
	}
	if setToken != nil {
		if key, ok := setToken.Claims.(*MyClaims); ok && setToken.Valid {
			return key, message.SUCCSES
		} else {
			return nil, message.ERROR_TOKEN_WRONG
		}
	}
	return nil, message.ERROR_TOKEN_WRONG
}

// jwt中间件
func JwtToken() iris.Handler {
	return func(c iris.Context) {
		var code int
		tokenHeader := c.GetHeader("Authorization")
		if tokenHeader == "" {
			code = message.ERROR_TOKEN_EXIST
			c.JSON(iris.Map{
				"status":  code,
				"message": message.GetErrMsg(code),
			})

			return
		}
		checkToken := strings.Split(tokenHeader, " ")
		if len(checkToken) == 0 {
			code = message.ERROR_TOKEN_TYPE_WRONG
			c.JSON(iris.Map{
				"status":  code,
				"message": message.GetErrMsg(code),
			})

			return
		}

		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = message.ERROR_TOKEN_TYPE_WRONG
			c.JSON(iris.Map{
				"status":  code,
				"message": message.GetErrMsg(code),
			})

			return
		}
		_, tCode := CheckToken(checkToken[1])
		if tCode != message.SUCCSES {
			code = tCode
			c.JSON(iris.Map{
				"status":  code,
				"message": message.GetErrMsg(code),
			})

			return
		}
		//c.Set("username", key)
		c.Next()
	}
}