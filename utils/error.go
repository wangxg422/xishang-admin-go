package utils

import (
	"errors"
	"gorm.io/gorm"
)

func NoRecord(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
