package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)
//第一个middleware
func MiddlewareOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("before middlewareOne: "+time.Now().String())
		c.Next()
		fmt.Println("after middlewareOne: "+time.Now().String())
		return
	}
}
//第二个middleware
func MiddlewareTwo() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("before middlewareTwo: "+time.Now().String())
		c.Next()
		fmt.Println("after middlewareTwo: "+time.Now().String())
		return
	}
}
//第三个middleware
func MiddlewareThree() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("before MiddlewareThree: "+time.Now().String())
		c.Next()
		fmt.Println("after MiddlewareThree: "+time.Now().String())
		return
	}
}
