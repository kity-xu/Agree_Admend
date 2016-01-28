package dllinterop

import (
	"syscall"
)

//加载dll
func LoadDll(dllName string) (dll *syscall.DLL, err error) {
	return syscall.LoadDLL(dllName)
}

//卸载dll
func UnloadDll(dll *syscall.DLL) {
	if dll != nil {
		dll.Release()
	}
}

//加载dll中的函数
func LoadProc(dll *syscall.DLL, procName string) (uintptr, error) {
	return syscall.GetProcAddress(dll.Handle, procName)
}

//为调用DLL中的方法提供一个统一的接口
func CallProc(handle uintptr, args ...uintptr) (r1, r2 uintptr, err syscall.Errno) {
	l := uintptr(len(args))
	switch {
	case l <= 3:
		tmp := make([]uintptr, 3)
		for i, v := range args {
			tmp[i] = v
		}
		return syscall.Syscall(handle, l, tmp[0], tmp[1], tmp[2])
	case l <= 6:
		tmp := make([]uintptr, 6)
		for i, v := range args {
			tmp[i] = v
		}
		return syscall.Syscall6(handle, l, tmp[0], tmp[1], tmp[2],
			tmp[3], tmp[4], tmp[5])
	case l <= 9:
		tmp := make([]uintptr, 9)
		for i, v := range args {
			tmp[i] = v
		}
		return syscall.Syscall9(handle, l, tmp[0], tmp[1], tmp[2],
			tmp[3], tmp[4], tmp[5], tmp[6],
			tmp[7], tmp[8])
	case l <= 12:
		tmp := make([]uintptr, 12)
		for i, v := range args {
			tmp[i] = v
		}
		return syscall.Syscall12(handle, l, tmp[0], tmp[1], tmp[2],
			tmp[3], tmp[4], tmp[5], tmp[6],
			tmp[7], tmp[8], tmp[9],
			tmp[10], tmp[11])
	case l <= 15:
		tmp := make([]uintptr, 15)
		for i, v := range args {
			tmp[i] = v
		}
		return syscall.Syscall15(handle, l, tmp[0], tmp[1], tmp[2],
			tmp[3], tmp[4], tmp[5], tmp[6],
			tmp[7], tmp[8], tmp[9],
			tmp[10], tmp[11], tmp[12],
			tmp[13], tmp[14])
	default:
		return uintptr(0), uintptr(0), syscall.E2BIG
	}
}
