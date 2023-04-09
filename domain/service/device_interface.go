package service

type DeviceInterface interface {
	InsertData(devicename string, imei string) error
	GetData(imei string) (string, error)
	Test() (string, error)
}
