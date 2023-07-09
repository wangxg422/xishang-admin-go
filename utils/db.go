package utils

import (
	"fmt"
	"gorm.io/gorm"
	"reflect"
)

func ConcatLikeWhereCondition(db *gorm.DB, conditions []string, values ...string) {
	length := len(conditions)
	if length == len(values) {
		for i := 0; i < length; i++ {
			if values[i] != "" {
				db.Where(conditions[i]+" like ?", LikeQuery(values[i]))
			}
		}
	}
}

func ConcatEqualsStrWhereCondition(db *gorm.DB, conditions []string, values ...string) {
	length := len(conditions)
	if length == len(values) {
		for i := 0; i < length; i++ {
			if values[i] != "" {
				db.Where(conditions[i]+" = ?", values[i])
			}
		}
	}
}

//func ConcatOneEqualsInt8WhereCondition(db *gorm.DB, column string, value int8) {
//	if value != 0 {
//		db.Where(column+" = ?", value)
//	}
//}

func ConcatOneEqualsStrWhereCondition(db *gorm.DB, column string, value string) {
	if value != "" {
		db.Where(column+" = ?", value)
	}
}

func ConcatTimeRangeWhereCondition(db *gorm.DB, start string, end string) {
	if start == "" && end == "" {
		return
	}

	if start == "" {
		db.Where("create_time <= ?", end+" 23:59:59")
	} else if end == "" {
		db.Where("create_time >= ?", start+" 00:00:00")
	} else {
		db.Where("create_time >= ? AND create_time <= ?", start+" 00:00:00", end+" 23:59:59")
	}
}

func StructToMap(in any) (map[string]any, error) {
	out := make(map[string]any)

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("ToMap only accepts struct or struct pointer; got %T", v)
	}

	t := v.Type()
	// 遍历结构体字段,map的key为结构体字段名称转下划线形式
	for i := 0; i < v.NumField(); i++ {
		fi := t.Field(i)
		out[CamelToCase(fi.Name)] = v.Field(i).Interface()
	}

	return out, nil
}
