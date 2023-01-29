package model

import (
	"time"
)

// 不包含逻辑删除的模型
type CommonEntity struct {
	ID        int       `gorm:"primary_key;auto_increment" json:"id" mapstructure:"id"`
	CreatedOn time.Time `json:"created_on" mapstructure:"-"`
	UpdatedOn time.Time `json:"updated_on" mapstructure:"-"`
	IsDeleted int8      `json:"is_deleted" mapstructure:"-"`
}
