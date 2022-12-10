package utils

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

func CheckStrNil(params ...any) bool {
	for _, param := range params {
		param := fmt.Sprint(param)
		fmt.Println("param:", param)
		if param == "" {
			return true
		}
	}
	return false
}

// ReadBodyToMap 读取请求体中的数据到Map
func ReadBodyToMap(c *gin.Context) (map[string]interface{}, error) {
	data, err := c.GetRawData()
	if err != nil {
		return nil, err
	}

	body := make(map[string]interface{})
	err = json.Unmarshal(data, &body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
