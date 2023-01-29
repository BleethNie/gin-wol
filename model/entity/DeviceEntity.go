package model

import (
	"reflect"
)

type DeviceEntity struct {
	CommonEntity
	UserId      int    `gorm:"type:int;not null;comment:用户 ID" json:"user_id"`
	Title       string `gorm:"type:varchar(100);not null;comment:文章标题" json:"title"`
	Uuid        string `gorm:"type:varchar(50);not null;comment:唯一标识" json:"uuid"`
	Desc        string `gorm:"type:varchar(200);comment:文章描述" json:"desc"`
	Content     string `gorm:"type:longtext;comment:文章内容" json:"content"`
	Img         string `gorm:"type:varchar(100);comment:封面图片地址" json:"img"`
	Type        int8   `gorm:"type:tinyint;comment:类型(1-原创 2-转载 3-翻译)" json:"type"`
	Status      int8   `gorm:"type:tinyint;comment:状态(1-公开 2-私密)" json:"status"`
	IsTop       *int8  `gorm:"type:tinyint;not null;default:0;comment:是否置顶(0-否 1-是)" json:"is_top"`
	OriginalUrl string `gorm:"type:varchar(100);comment:源链接" json:"original_url"`
}

func (a *DeviceEntity) IsEmpty() bool {
	return reflect.DeepEqual(a, &DeviceEntity{})
}
