package kvm

import (
	"cn.agree/driverlayer"
	"cn.agree/utils"
	"encoding/hex"
)

//KVM通用实现
type KVMGeneral struct {
}

func (jst *KVMGeneral) Initdriver(pin *driverlayer.DriverArg) {
}

func (jst *KVMGeneral) Deinit(pin *driverlayer.DriverArg) error {
	return nil
}

func (jst *KVMGeneral) GetFactoryName() string {
	return "通用驱动"
}

func (jst *KVMGeneral) DeSync(pin *driverlayer.DriverArg, timeout int) error {
	utils.Debug("KVMGeneral received DeSync request")
	s, _ := hex.DecodeString("fa5a1102000067")
	b, e := driverlayer.WritePortAndReadWithLen(pin.Port, pin.Baud, s, 0, timeout, nil)
	if e != nil {
		return e
	} else {
		utils.Debug("KVMGeneral DeSync data %s,begin to process", b)
		return nil
	}
}

func (jst *KVMGeneral) OutSyncInner(pin *driverlayer.DriverArg, timeout int) error {
	utils.Debug("KVMGeneral received OutSyncInner request")
	s, _ := hex.DecodeString("fa5a1103000068")
	b, e := driverlayer.WritePortAndReadWithLen(pin.Port, pin.Baud, s, 0, timeout, nil)
	if e != nil {
		return e
	} else {
		utils.Debug("KVMGeneral OutSyncInner data %s,begin to process", b)
		return nil
	}
}

func (jst *KVMGeneral) InnSyncOuter(pin *driverlayer.DriverArg, timeout int) error {
	utils.Debug("KVMGeneral received InnSyncOuter request")
	s, _ := hex.DecodeString("fa5a1101000066")
	b, e := driverlayer.WritePortAndReadWithLen(pin.Port, pin.Baud, s, 0, timeout, nil)
	if e != nil {
		return e
	} else {
		utils.Debug("KVMGeneral InnSyncOuter data %s,begin to process", b)
		return nil
	}
}
