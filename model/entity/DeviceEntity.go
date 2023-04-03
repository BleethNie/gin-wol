package model

import (
	"reflect"
)

type DeviceEntity struct {
	CommonEntity
	Ip       string `gorm:"type:varchar(30);not null;comment:ip地址" json:"ip"`
	Mac      string `gorm:"type:varchar(30);not null;comment:Mac地址" json:"mac"`
	NickName string `gorm:"type:varchar(50);not null;comment:备注" json:"nick_name"`
	HostName string `gorm:"type:varchar(200);comment:主机名" json:"host_name"`
}

func (DeviceEntity) TableName() string {
	return "device"
}

func (a *DeviceEntity) IsEmpty() bool {
	return reflect.DeepEqual(a, &DeviceEntity{})
}
