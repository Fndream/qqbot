package reward

import (
	"gorm.io/gorm"
	"qqbot/pkg/db"
)

func DB(tx *gorm.DB) *gorm.DB {
	if tx == nil {
		return db.DB
	}
	return tx
}
