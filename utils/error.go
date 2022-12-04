package utils

import (
	"errors"
	"gorm.io/gorm"
)

func NoRecord(err error) bool {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return true
	}

	return false
}

type Null struct {
}
