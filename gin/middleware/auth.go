package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

/** 注册 */
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("before middleware 0")
		c.Set("request", "clinet_request")
		c.Next()
		fmt.Println("before middleware 1")
	}
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("auth")
		c.Next()
	}
}

func ApiV1() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("v1")
		c.Next()
	}
}
func ApiV12() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("v12")
		c.Next()
	}
}
