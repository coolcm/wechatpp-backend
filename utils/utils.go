package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 验证请求参数是否为空
func VerifyParams(c *gin.Context, params map[string]string) (bool) {
	result := make([]string, 0)
	for k, v := range params {
		if v == "" {
			result = append(result, k)
		}
	}
	if cap(result) == 0 {
		return true
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "msg": fmt.Sprintf("%v can not be empty", result)})
		return false
	}
}