package protocol

import (
	"cn.agree/driverlayer"
	"cn.agree/driverlayer/finger"
	"cn.agree/driverlayer/iccard"
	"cn.agree/driverlayer/idccard"
	"cn.agree/driverlayer/kvm"
	"cn.agree/driverlayer/msf"
	"cn.agree/driverlayer/pin"
	"cn.agree/driverlayer/pingjia"
	"cn.agree/driverlayer/pr2"
	"cn.agree/utils"
	"errors"
	"strconv"
	"strings"
)

//设备类别
const (
	PIN = iota
	MSF
	IC
	IDR
	FP
	PINGJIA
	PR2PRT
	KVM
	MEDIA
	SCN
	MAX_DEVICE
)

//设备状态
const (
	DEVICE_DISABLE = iota //设备未启用
	DEVICE_IDLE           //设备空闲
	DEVICE_BUSY           //设备正忙
)

//设备的总的描述结构
type device_node struct {
	DevKind     int
	State       int
	Category    string
	FactoryName string
	IsOccupy    bool
	drv         driverlayer.DriverBase
	drvarg      *driverlayer.DriverArg
}

type dev_con_relation struct {
	DevName   string
	DevKind   int
	DevConfig string
}

var DEV_NOT_FOUND = errors.New("not found this kind device ")
var DEV_ALREADY_IMPLEMENTED = errors.New("this kind of device has implemented")
var DEV_NOT_IMPLEMENTED = errors.New("this kind of device has not implemented")

var alldevice = make([]device_node, MAX_DEVICE)
var DevKinds = []dev_con_relation{
	dev_con_relation{"Pin", PIN, "devices/Pinconfig.xml"},
	dev_con_relation{"Pr2prt", PR2PRT, "devices/Pr2prtconfig.xml"},
	dev_con_relation{"Fp", FP, "devices/Fpconfig.xml"},
	dev_con_relation{"Pingjia", PINGJIA, "devices/Pingjiaconfig.xml"},
	dev_con_relation{"Ic", IC, "devices/Icconfig.xml"},
	dev_con_relation{"Idr", IDR, "devices/Idrconfig.xml"},
	dev_con_relation{"Msf", MSF, "devices/Msfconfig.xml"},
	dev_con_relation{"Scn", SCN, "devices/Scnconfig.xml"},
	dev_con_relation{"Kvm", KVM, "devices/Kvmconfig.xml"}}

func init() {
	var (
		i = PIN
		s string
	)

	//设置所有的硬件类别
	for ; i < MAX_DEVICE; i++ {
		s, _ = getDeviceCategory(i)
		alldevice[i].Category = s
	}
}

//获得设备实现
func getDeviceImp(DevKind int) (driverlayer.DriverBase, *driverlayer.DriverArg, error) {
	if DevKind >= 0 && DevKind < MAX_DEVICE {
		if alldevice[DevKind].IsOccupy == false {
			return nil, nil, DEV_NOT_IMPLEMENTED
		}
		return alldevice[DevKind].drv, alldevice[DevKind].drvarg, nil
	}
	return nil, nil, DEV_NOT_FOUND
}

//根据下标获得设备
func getDevice(DevKind int) (*device_node, error) {
	if DevKind >= 0 && DevKind < MAX_DEVICE {
		return &alldevice[DevKind], nil
	}
	return nil, DEV_NOT_FOUND
}

//获得所有设备
func getAllDevice() *[]device_node {
	return &alldevice
}

//设置设备实现
func setDeviceImp(DevKind int, drv driverlayer.DriverBase, drvarg *driverlayer.DriverArg) error {
	if DevKind >= 0 && DevKind < MAX_DEVICE {
		if alldevice[DevKind].IsOccupy == false {
			alldevice[DevKind].IsOccupy = true
			alldevice[DevKind].drv = drv
			alldevice[DevKind].FactoryName = drv.GetFactoryName()
			alldevice[DevKind].drvarg = drvarg
			alldevice[DevKind].State = DEVICE_IDLE
			return nil
		} else {
			return DEV_ALREADY_IMPLEMENTED
		}
	}
	return DEV_NOT_FOUND
}

//把某项设备标为正在使用
func MarkDeviceUse(DevKind int) (bool, error) {
	if DevKind >= 0 && DevKind < MAX_DEVICE {
		if alldevice[DevKind].State == DEVICE_IDLE {
			alldevice[DevKind].State = DEVICE_BUSY
			return true, nil
		} else {
			utils.Debug("MarkDeviceUse %s is not in correct state  %d state ", alldevice[DevKind].Category, alldevice[DevKind].State)
		}
		return false, nil
	}
	return false, DEV_NOT_FOUND
}

//把某项设备标为空闲
func MarkDeviceIDLE(DevKind int) (bool, error) {
	if DevKind >= 0 && DevKind < MAX_DEVICE {
		if alldevice[DevKind].State == DEVICE_BUSY {
			alldevice[DevKind].State = DEVICE_IDLE
			return true, nil
		} else {
			utils.Warn(" MarkDeviceIDLE %s is not in correct state %d ", alldevice[DevKind].Category, alldevice[DevKind].State)
		}
		return false, nil
	}
	return false, DEV_NOT_FOUND
}

//获得类型的下标值
func getDeviceIndex(DevName string) (int, error) {
	for _, s := range DevKinds {
		if strings.EqualFold(DevName, s.DevName) {
			return s.DevKind, nil
		}
	}
	return -1, DEV_NOT_FOUND
}

//获得下标对应的类型
func getDeviceCategory(DevKind int) (string, error) {
	for _, s := range DevKinds {
		if s.DevKind == DevKind {
			return s.DevName, nil
		}
	}
	return "", DEV_NOT_FOUND
}

//组装arg,如果有不是通用参数的,则放入extractarg中
func populateArg(ps []DevParam) *driverlayer.DriverArg {
	t := driverlayer.DriverArg{}
	t.Baud = 9600
	for _, v := range ps {
		switch v.Name {
		case "Factory":
			t.FactoryName = v.Value
			break

		case "Port":
			t.Port = v.Value
			break

		case "ExtPort":
			t.ExtPort = v.Value
			break

		case "Baud":
			p, err := strconv.Atoi(v.Value)
			if err != nil {
				utils.Error("parse Baud error,set default value 9600 : %+v", v)
				continue
			}
			t.Baud = p
			break

		default:
			ext := driverlayer.DriverExtraParam{Name: v.Name, Value: v.Value}
			t.ExtraParam = append(t.ExtraParam, ext)
		}
	}
	return &t
}

func findFactoryParam(ps []DevParam) string {
	for _, s := range ps {
		if strings.EqualFold(s.Name, "Factory") {
			return s.Value
		}
	}
	return ""
}

func StartDriverService(drivers []UseDriverCategory) error {
	var err error
	var v UseDriverCategory
	var k int
	var drvbase driverlayer.DriverBase
	var facname string

	Rollback := func(driver driverlayer.DriverBase, arg *driverlayer.DriverArg) {
		if err != nil {
			driver.Deinit(arg)
		}
	}

	for _, v = range drivers {
		drvbase = nil
		k, err = getDeviceIndex(v.Category)
		if err != nil {
			utils.Error("can't find Category impl:%s", v.Category)
			continue
		}
		facname = findFactoryParam(v.Kind)
		if facname == "" {
			utils.Error("can't find Category:%v's factory", v.Category)
			continue
		}

		switch k {
		case PIN:
			switch facname {
			case "南天":

				break

			case "通用类型":
				drvbase = &pin.PinGeneral{}
				break

			default:

			}

		case FP:
			switch facname {
			case "通用类型":
				drvbase = &finger.ZZFinger{}
				break

				//			case "天诚":
				//				drvbase = &finger.TCFinger{}
				//				break

			case "中正":
				drvbase = &finger.ZZFinger{}
				break

			default:

			}

		case PINGJIA:
			switch facname {
			case "南天":

				break

			case "升腾":
				drvbase = &pingjia.STPingjia{}
				break

			case "通用类型":
				drvbase = &pingjia.HTPingjia{}

			default:
			}

		case MSF:
			switch facname {
			case "实达":
				drvbase = &bankcard.SDMsf{}
				break

			default:

			}

		case IC:
			switch facname {
			case "通用类型":
				drvbase = &iccard.ICGeneral{}
				break
				//			case "升腾":
				//				drvbase = &iccard.STICCard{}
				//				break
			default:

			}

		case IDR:
			switch facname {
			case "通用类型":
				drvbase = &idccard.IDCGeneral{}
				break

				//			case "升腾":
				//				drvbase = &idccard.STIdr{}
				//				break
			default:

			}

		case PR2PRT:
			switch facname {
			case "PR2":
				drvbase = &pr2.OlevittePrinter{}
				break
			case "OKI":
				drvbase = &pr2.OkiPrinter{}
			default:

			}

		case KVM:
			switch facname {
			case "通用类型":
				drvbase = &kvm.KVMGeneral{}
				break
			}

		}
		//进行初始化
		if drvbase != nil {
			cat, _ := getDeviceCategory(k)
			utils.Info("%s  use %s driver", cat, drvbase.GetFactoryName())
			t := populateArg(v.Kind)
			setDeviceImp(k, drvbase, t)
			drvbase.Initdriver(t)
			defer Rollback(drvbase, t)
		} else {
			utils.Error(`Category %s could not find factory %s`, v.Category, facname)
		}
	}
	return err
}
