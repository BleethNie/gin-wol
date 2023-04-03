package utils

import (
	"container/list"
	"encoding/json"
	"fmt"
	"github.com/BleethNie/gin-wol/config"
	model "github.com/BleethNie/gin-wol/model/entity"
	bolt "go.etcd.io/bbolt"
	"log"
)

var db *bolt.DB

const (
	BUCKET = "DEVICE"
)

func InitDB() *bolt.DB {
	if db == nil {
		path := config.Cfg.Db.Path
		db, err := bolt.Open(path, 0600, nil)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
	}
	return db
}

func QueryDeviceByMac(mac string) model.DeviceEntity {
	// 查询数据
	var deviceEntity model.DeviceEntity

	var queryFn = func(tx *bolt.Tx, bucketName, keyName string) error {
		// 按照 bucket 名称查找 bucket
		b := tx.Bucket([]byte(BUCKET))
		if b == nil {
			return fmt.Errorf("Bucket not found")
		}

		// 按照 key 查找 value
		byteDevice := b.Get([]byte(mac))
		if byteDevice == nil {
			return fmt.Errorf("Key not found")
		}
		if err := json.Unmarshal(byteDevice, &deviceEntity); err != nil {
			panic(err)
		}
		return nil
	}
	// 开启读事务并执行查询函数
	err := db.View(func(tx *bolt.Tx) error {
		return queryFn(tx, BUCKET, mac)
	})
	if err != nil {
		panic(err)
	}
	return deviceEntity
}

func QueryAllDevice() *list.List {
	// 创建或获取 bucket
	deviceList := list.New()

	var queryFn = func(tx *bolt.Tx, bucketName string) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return err
		}
		return b.ForEach(func(k, v []byte) error {
			mac := string(k)
			device := QueryDeviceByMac(mac)
			deviceList.PushBack(device)
			return nil
		})
	}
	// 开启读事务并执行查询函数
	err := db.View(func(tx *bolt.Tx) error {
		return queryFn(tx, BUCKET)
	})
	if err != nil {
		panic(err)
	}
	return deviceList
}

func UpdateDevice(deviceEntity model.DeviceEntity) {
	mac := deviceEntity.Mac
	var updateFn = func(tx *bolt.Tx, bucketName, keyName string) error {
		// 创建或获取 bucket
		b, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return err
		}

		// 将结构体转换为字节数组
		valueBytes, err := json.Marshal(deviceEntity)
		if err != nil {
			panic(err)
		}

		// 存储数据
		err = b.Put([]byte(keyName), valueBytes)
		if err != nil {
			return err
		}
		return nil
	}
	// 开启读事务并执行查询函数
	err := db.Update(func(tx *bolt.Tx) error {
		return updateFn(tx, BUCKET, mac)
	})
	if err != nil {
		panic(err)
	}
}
