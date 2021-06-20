package component

import "gorm.io/gorm"

type Mysql struct {
	Write *gorm.DB
	Read  *gorm.DB
}
