package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func RandomStr(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}

func PrintJson(o any) {
	bs, _ := json.Marshal(o)
	var out bytes.Buffer
	_ = json.Indent(&out, bs, "", "\t")
	fmt.Println(out.String())
}

// FirstUpper 字符串首字母大写
func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// FirstLower 字符串首字母小写
func FirstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

func LikeQuery(s string) string {
	return "%" + s + "%"
}

func StrToInt64Array(list []string) ([]int64, error) {
	if len(list) == 0 {
		return []int64{}, nil
	}

	var arr []int64
	for _, s := range list {
		v, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil, err
		}
		arr = append(arr, v)
	}

	return arr, nil
}

// CamelToCase 驼峰式转下划线
func CamelToCase(s string) string {
	if s == "" {
		return s
	}

	var v byte = 'a' - 'A'

	length := len(s)
	var builder strings.Builder
	for i := 0; i < length; i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			if i == 0 {
				builder.WriteByte(s[i] + v)
			} else {
				builder.WriteByte('_')
				builder.WriteByte(s[i] + v)
			}
		} else {
			builder.WriteByte(s[i])
		}
	}
	return builder.String()
}

// CaseToCamel 下划线转驼峰
func CaseToCamel(s string) string {
	// TODO 用到时实现
	return s
}
