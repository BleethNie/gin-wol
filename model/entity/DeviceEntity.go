package model

import (
	"reflect"
)

type DeviceEntity struct {
	Ip       string `gorm:"type:varchar(30);not null;comment:ip地址" json:"ip"`
	Mac      string `gorm:"type:varchar(30);not null;comment:Mac地址" json:"mac"`
	NickName string `gorm:"type:varchar(50);not null;comment:备注" json:"nickname"`
	HostName string `gorm:"type:varchar(50);comment:主机名" json:"hostname"`
}

func (DeviceEntity) TableName() string {
	return "device"
}

func (a *DeviceEntity) IsEmpty() bool {
	return reflect.DeepEqual(a, &DeviceEntity{})
}
