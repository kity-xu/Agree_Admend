// provides all devices list
package protocol

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
)

type FactoryKind struct {
	Factory string `xml:"Factory,attr"`
}

type DriverCategory struct {
	Name string        `xml:"Name,attr"`
	Kind []FactoryKind `xml:"Kind"`
}

type XmlDevices struct {
	Version  string           `xml:"Version,attr"`
	Category []DriverCategory `xml:"Category"`
}

var AllDevice XmlDevices

//add by yangxiaolong
type ConnMethod struct {
     Method  string    `xml:"Method,attr"`
     Address string    `xml:"Address,attr"`
}

var Connect ConnMethod 

//解析所有的驱动类别
func ParserDrivers() error {
	buffer, err := ioutil.ReadFile("devices/Deviceslist.xml")
	if err != nil {
		return errors.New("open devices.xml file Fail:" + err.Error())
	}
	err = xml.Unmarshal(buffer, &AllDevice)
	if err != nil {
		return errors.New("get config file fail:" + err.Error())
	}
	return nil
}

//add by yangxiaolong
func ParserConnMethod() error {
	buffer, err := ioutil.ReadFile("devices/ConnectMethod.xml")
	if err != nil {
		return errors.New("open ConnectMethod.xml file Fail:" + err.Error())
	}
	err = xml.Unmarshal(buffer, &Connect)
	if err != nil {
		return errors.New("get connect method file Fail:" + err.Error())
	}
	return nil
}