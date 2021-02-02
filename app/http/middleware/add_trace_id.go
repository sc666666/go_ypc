package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// 添加跟踪 ID
func AddTraceId() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 请求前
		c.Header("trace_id", uniqid(""))
		c.Next()
		// 请求后
	}
}

func uniqid(prefix string) string {
	now := time.Now()
	sec := now.Unix()
	usec := now.UnixNano() % 0x100000
	return fmt.Sprintf("%s%08x%05x", prefix, sec, usec)
}
