package repository

import (
	"time"
)

type Session struct {
	DeviceName string    `json:"DeviceName"`
	Imei       string    `json:"Imei"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
}
type DeviceDomain struct {
	DeviceName string    `json:"DeviceName"`
	Imei       string    `json:"Imei"`
	CreatedAt  time.Time `json:"createdAt" bson:"createdAt"`
}

type DeviceInterface interface {
	SetData(devicename string, imei string) error
	SetBackUp(devicename string, imei string) error
	GetData(imei string) (string, error)
	GetBackUp(imei string) (string, error)
}
