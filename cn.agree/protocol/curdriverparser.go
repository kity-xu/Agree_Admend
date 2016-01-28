// provides parse current use device list
package protocol

import (
	"cn.agree/utils"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"strings"
)

var DRIVE_CONFIG_FILE_NOT_FOUND = errors.New("not found driver config file")

type DevParam struct {
	Name  string `xml:"Name,attr"`
	Value string `xml:"Value,attr"`
	Desc  string `xml:"Desc,attr"`
}

type UseDriverCategory struct {
	Category string     `xml:"Category,attr"`
	Kind     []DevParam `xml:"Param"`
}

type UseDevice struct {
	Device []UseDriverCategory
}

var CurDrivers UseDevice

//解析所有的使用配置文件
func ParseUseDriver() error {

	fs, err := GetAllDrvConfig()
	err = xml.Unmarshal(fs, &CurDrivers)
	if err != nil {
		return err
	}
	return nil
}

func GetAllDrvConfig() ([]byte, error) {
	var (
		err error
		fs  []byte
	)
	buffer := make([]byte, 0, 700)
	buffer = append(buffer, []byte("<Root>")...)
	for _, s := range DevKinds {
		fs, err = ioutil.ReadFile(s.DevConfig)
		if err != nil {
			utils.Error("can't read config file,category is [%s],config file is [%s]", s.DevName, s.DevConfig)
		}
		buffer = append(buffer, fs...)
	}
	buffer = append(buffer, []byte("</Root>")...)
	return buffer, nil
}

func GetAllDeviceList() ([]byte, error) {
	var (
		err error
	)
	buffer := make([]byte, 0, 700)
	buffer, err = ioutil.ReadFile("devices/Deviceslist.xml")
	if err != nil {
		utils.Error("can't read config file,config file is [devices/Devicelist.xml]")
	}
	return buffer, nil
}

//获取设备配置
func GetDrvConfig(drv string) ([]byte, error) {
	for _, s := range DevKinds {
		if strings.EqualFold(drv, s.DevName) {
			fs, err := ioutil.ReadFile(s.DevConfig)
			if err != nil {
				return nil, err
			}
			return fs, nil
		}
	}
	return nil, DRIVE_CONFIG_FILE_NOT_FOUND
}

//保存设备配置
func SaveDrvConfig(drv string, config string) error {
	for _, s := range DevKinds {
		if strings.EqualFold(drv, s.DevName) {
			err := ioutil.WriteFile(s.DevConfig, []byte(config), 0644)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return DRIVE_CONFIG_FILE_NOT_FOUND
}
