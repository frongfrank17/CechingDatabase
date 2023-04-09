package repository

type DeviceDomain struct {
	DeviceName string `json:"DeviceName"`
	Imei       string `json:"Imei"`
}

type DeviceInterface interface {
	SetData(devicename string, imei string) error
	GetData(imei string) (string, error)
}
