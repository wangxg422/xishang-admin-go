package utils

import (
	"gorm.io/gorm"
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

func ConcatEqualsWhereCondition(db *gorm.DB, conditions []string, values ...string) {
	length := len(conditions)
	if length == len(values) {
		for i := 0; i < length; i++ {
			db.Where(conditions[i]+" = ?", values[i])
		}
	}
}

func ConcatOneEqualsWhereCondition(db *gorm.DB, column string, value any) {
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
