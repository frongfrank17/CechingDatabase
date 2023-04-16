package service

import (
	"fmt"

	repository "CachingDatabase/domain/repository"
)

type DeviceInterface interface {
	InsertData(devicename string, imei string) error
	GetData(imei string) (string, error)
	Test() (string, error)
}
type deviceService struct {
	repo repository.DeviceInterface
}

func NewService(repo repository.DeviceInterface) deviceService {
	return deviceService{repo: repo}
}

func (svr deviceService) InsertData(devicename string, imei string) error {
	fmt.Println("Service : ", imei, " Values : ", devicename)
	err := svr.repo.SetData(devicename, imei)
	if err != nil {
		return err
	}
	err = svr.repo.SetBackUp(devicename, imei)
	if err != nil {
		return err
	}
	return nil
}
func (svr deviceService) Test() (string, error) {
	return "Test", nil
}
func (svr deviceService) GetData(imei string) (string, error) {
	result, err := svr.repo.GetData(imei)
	if err != nil {
		return "", err
	}
	_, err = svr.repo.GetBackUp(imei)
	if err != nil {
		return "", err
	}
	return result, nil
}
