package middleware

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

func Cors() iris.Handler {
	return cors.New(
		cors.Options{
			AllowedOrigins: []string{"https://lookarm.cn","http://lookarm.cn"},
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders: []string{"*"},
			// , "Content-Type"
			ExposedHeaders: []string{"Content-Length", "text/plain", "Authorization"},
			AllowCredentials: true,
		})
}
