// +build windows

// http://support.microsoft.com/en-gb/kb/823179
// This step-by-step article describes how to access serial ports and how to access parallel ports by using Microsoft Visual Basic .NET.
// Use platform invoke services to call Win32 API functions in Visual Basic .NET to access serial and parallel ports

package parallel

import (
	"fmt"
	"os"
	"sync"
	"syscall"
	"time"
	"unsafe"
)

type Port struct {
	f  *os.File
	fd syscall.Handle
	rl sync.Mutex
	wl sync.Mutex
	ro *syscall.Overlapped
	wo *syscall.Overlapped
}

//'Declare structures.
//   Public Structure DCB
//   Public DCBlength As Int32
//   Public BaudRate As Int32
//   Public fBitFields As Int32 'See Comments in Win32API.Txt
//   Public wReserved As Int16
//   Public XonLim As Int16
//   Public XoffLim As Int16
//   Public ByteSize As Byte
//   Public Parity As Byte
//   Public StopBits As Byte
//   Public XonChar As Byte
//   Public XoffChar As Byte
//   Public ErrorChar As Byte
//   Public EofChar As Byte
//   Public EvtChar As Byte
//   Public wReserved1 As Int16 'Reserved; Do Not Use
//End Structure
type structDCB struct {
	DCBlength, BaudRate                            uint32
	flags                                          [4]byte
	wReserved, XonLim, XoffLim                     uint16
	ByteSize, Parity, StopBits                     byte
	XonChar, XoffChar, ErrorChar, EofChar, EvtChar byte
	wReserved1                                     uint16
}

//Public Structure COMMTIMEOUTS
//   Public ReadIntervalTimeout As Int32
//   Public ReadTotalTimeoutMultiplier As Int32
//   Public ReadTotalTimeoutConstant As Int32
//   Public WriteTotalTimeoutMultiplier As Int32
//   Public WriteTotalTimeoutConstant As Int32
//End Structure
type structTimeouts struct {
	ReadIntervalTimeout         uint32
	ReadTotalTimeoutMultiplier  uint32
	ReadTotalTimeoutConstant    uint32
	WriteTotalTimeoutMultiplier uint32
	WriteTotalTimeoutConstant   uint32
}

func openPort(name string) (p *Port, err error) {
	if len(name) > 0 && name[0] != '\\' {
		name = "\\\\.\\" + name
	}

	// hParallelPort = CreateFile("LPT1",
	//	GENERIC_READ Or GENERIC_WRITE,
	//  0,
	//  IntPtr.Zero,
	//	OPEN_EXISTING,
	//	FILE_ATTRIBUTE_NORMAL,
	//	IntPtr.Zero)

	//h, err := syscall.CreateFile(syscall.StringToUTF16Ptr(name),
	//	syscall.GENERIC_READ|syscall.GENERIC_WRITE,
	//	0,
	//	nil,
	//	syscall.OPEN_EXISTING,
	//	syscall.FILE_ATTRIBUTE_NORMAL|syscall.FILE_FLAG_OVERLAPPED,
	//	0)

	//CreateFile(string lpFileName,
	//	uint dwDesiredAccess,
	//	uint dwShareMode,
	//	IntPtr lpSecurityAttributes,
	//	uint dwCreationDisposition,
	//  uint dwFlagsAndAttributes,
	//	IntPtr hTemplateFile);

	//Public Const GENERIC_READ As Int32 = &H80000000
	//Public Const GENERIC_WRITE As Int32 = &H40000000
	//Public Const OPEN_EXISTING As Int32 = 3
	//Public Const FILE_ATTRIBUTE_NORMAL As Int32 = &H80
	//Public Const NOPARITY As Int32 = 0
	//Public Const ONESTOPBIT As Int32 = 0

	// https://msdn.microsoft.com/en-us/library/windows/desktop/aa363201(v=vs.85).aspx
	//hCom = CreateFile( pcCommPort,
	//    GENERIC_READ | GENERIC_WRITE,
	//    0,      //  must be opened with exclusive-access
	//    NULL,   //  default security attributes
	//    OPEN_EXISTING, //  must use OPEN_EXISTING
	//    0,      //  not overlapped I/O
	//    NULL ); //  hTemplate must be NULL for comm devices

	h, err := syscall.CreateFile(syscall.StringToUTF16Ptr(name),
		syscall.GENERIC_READ|syscall.GENERIC_WRITE,
		0,
		nil,
		syscall.OPEN_EXISTING,
		syscall.FILE_ATTRIBUTE_NORMAL,
		0)

	// Verify that the obtained handle is valid.
	if err != nil {
		return nil, err
	}
	f := os.NewFile(uintptr(h), name)
	defer func() {
		if err != nil {
			f.Close()
		}
	}()

	fmt.Println("Serial Port Handler is ready for ", name)

	fmt.Println("--- Get State ---")
	if err = getCommState(h); err != nil {
		return
	}

	fmt.Println("--- State ---")
	if err = setCommState(h); err != nil {
		return
	}

	fmt.Println("--- Setup ---")
	//if err = setupComm(h, 64, 64); err != nil {
	if err = setupComm(h, 1, 1); err != nil {
		return
	}

	fmt.Println("--- Timeouts ---")
	if err = setCommTimeouts(h, 0); err != nil {
		return
	}

	fmt.Println("Mask")
	if err = setCommMask(h); err != nil {
		return
	}

	ro, err := newOverlapped()
	if err != nil {
		return
	}
	wo, err := newOverlapped()
	if err != nil {
		return
	}
	port := new(Port)
	port.f = f
	port.fd = h
	port.ro = ro
	port.wo = wo

	return port, nil
}

func (p *Port) Close() error {
	return p.f.Close()
}

func (p *Port) Write(buf []byte) (int, error) {
	p.wl.Lock()
	defer p.wl.Unlock()

	fmt.Println("Reset Event")

	if err := resetEvent(p.wo.HEvent); err != nil {
		return 0, err
	}

	fmt.Println("Reset Event Done")

	var n uint32

	fmt.Println("WriteFile")
	//Public Declare Auto Function WriteFile Lib "kernel32.dll" (ByVal hFile As IntPtr, _
	//   ByVal lpBuffer As Byte(), ByVal nNumberOfBytesToWrite As Int32, _
	//      ByRef lpNumberOfBytesWritten As Int32, ByVal lpOverlapped As IntPtr) As Boolean
	// Success = WriteFile(hParallelPort, Buffer, Buffer.Length, BytesWritten, IntPtr.Zero)
	err := syscall.WriteFile(p.fd, buf, &n, p.wo)

	fmt.Println("WriteFile Done")
	if err != nil && err != syscall.ERROR_IO_PENDING {
		return int(n), err
	}

	return getOverlappedResult(p.fd, p.wo)
}

func (p *Port) Read(buf []byte) (int, error) {
	if p == nil || p.f == nil {
		return 0, fmt.Errorf("Invalid port on read %v %v", p, p.f)
	}

	p.rl.Lock()
	defer p.rl.Unlock()

	fmt.Println("Reset Event")

	if err := resetEvent(p.ro.HEvent); err != nil {
		return 0, err
	}

	fmt.Println("Reset Event Done")

	var done uint32

	//Public Declare Auto Function ReadFile Lib "kernel32.dll" (ByVal hFile As IntPtr, _
	//   ByVal lpBuffer As Byte(), ByVal nNumberOfBytesToRead As Int32, _
	//      ByRef lpNumberOfBytesRead As Int32, ByVal lpOverlapped As IntPtr) As Boolean
	err := syscall.ReadFile(p.fd, buf, &done, p.ro)

	fmt.Println("WriteFile Done")

	if err != nil && err != syscall.ERROR_IO_PENDING {
		return int(done), err
	}
	return getOverlappedResult(p.fd, p.ro)
}

var (
	nSetCommState,
	nGetCommState,
	nSetCommTimeouts,
	nSetCommMask,
	nSetupComm,
	nGetOverlappedResult,
	nCreateEvent,
	nResetEvent,
	nGetLastError uintptr
)

func init() {
	k32, err := syscall.LoadLibrary("kernel32.dll")
	if err != nil {
		panic("LoadLibrary " + err.Error())
	}
	defer syscall.FreeLibrary(k32)

	// Get Procedure Addresses from Library
	nGetCommState = getProcAddr(k32, "GetCommState")
	nSetCommState = getProcAddr(k32, "SetCommState")
	nSetCommTimeouts = getProcAddr(k32, "SetCommTimeouts")
	nSetCommMask = getProcAddr(k32, "SetCommMask")
	nSetupComm = getProcAddr(k32, "SetupComm")
	nGetOverlappedResult = getProcAddr(k32, "GetOverlappedResult")
	nCreateEvent = getProcAddr(k32, "CreateEventW")
	nResetEvent = getProcAddr(k32, "ResetEvent")
	nGetLastError = getProcAddr(k32, "GetLastError")
}

func getProcAddr(lib syscall.Handle, name string) uintptr {
	addr, err := syscall.GetProcAddress(lib, name)
	if err != nil {
		panic(name + " " + err.Error())
	}
	return addr
}

func getCommState(h syscall.Handle) (err error) {
	var params structDCB
	params.DCBlength = uint32(unsafe.Sizeof(params))

	//Public Declare Auto Function GetCommState Lib "kernel32.dll" (ByVal nCid As IntPtr, _
	//   ByRef lpDCB As DCB) As Boolean
	r, _, err := syscall.Syscall(nGetCommState, 2, uintptr(h), uintptr(unsafe.Pointer(&params)), 0)

	fmt.Println("GetCommState: ", r, err)

	if r == 0 {
		_, _, err := syscall.Syscall(nGetLastError, 1, uintptr(h), 0, 0)
		fmt.Println("LastError : ", err)
		return nil
	}
	return nil
}

func setCommState(h syscall.Handle) error {
	var params structDCB
	params.DCBlength = uint32(unsafe.Sizeof(params))

	//params.flags[0] = 0x01  // fBinary
	//params.flags[0] |= 0x10 // Assert DSR

	params.flags[0] = 0x00

	params.BaudRate = uint32(9600)
	params.ByteSize = 8
	params.Parity = 0
	params.StopBits = 0
	//Public Declare Auto Function SetCommState Lib "kernel32.dll" (ByVal nCid As IntPtr, _
	//   ByRef lpDCB As DCB) As Boolean
	r, _, err := syscall.Syscall(nSetCommState, 2, uintptr(h), uintptr(unsafe.Pointer(&params)), 0)

	fmt.Println("SetCommState : ", r, err)

	if r == 0 {
		_, _, err := syscall.Syscall(nGetLastError, 1, uintptr(h), 0, 0)
		fmt.Println("LastError : ", err)
		return nil
	}
	return nil
}

func setCommTimeouts(h syscall.Handle, readTimeout time.Duration) error {
	var timeouts structTimeouts
	// const MAXDWORD = 1<<32 - 1

	//if readTimeout > 0 {
	//	// non-blocking read
	//	timeoutMs := readTimeout.Nanoseconds() / 1e6
	//	if timeoutMs < 1 {
	//		timeoutMs = 1
	//	} else if timeoutMs > MAXDWORD {
	//		timeoutMs = MAXDWORD
	//	}
	//	timeouts.ReadIntervalTimeout = 0
	//	timeouts.ReadTotalTimeoutMultiplier = 0
	//	timeouts.ReadTotalTimeoutConstant = uint32(timeoutMs)
	//} else {
	//	// blocking read
	//	timeouts.ReadIntervalTimeout = MAXDWORD
	//	timeouts.ReadTotalTimeoutMultiplier = MAXDWORD
	//	timeouts.ReadTotalTimeoutConstant = MAXDWORD - 1
	//}

	/* From http://msdn.microsoft.com/en-us/library/aa363190(v=VS.85).aspx

		 For blocking I/O see below:

		 Remarks:

		 If an application sets ReadIntervalTimeout and
		 ReadTotalTimeoutMultiplier to MAXDWORD and sets
		 ReadTotalTimeoutConstant to a value greater than zero and
		 less than MAXDWORD, one of the following occurs when the
		 ReadFile function is called:

		 If there are any bytes in the input buffer, ReadFile returns
		       immediately with the bytes in the buffer.

		 If there are no bytes in the input buffer, ReadFile waits
	               until a byte arrives and then returns immediately.

		 If no bytes arrive within the time specified by
		       ReadTotalTimeoutConstant, ReadFile times out.
	*/
	timeouts.ReadIntervalTimeout = 0
	timeouts.ReadTotalTimeoutConstant = 0
	timeouts.ReadTotalTimeoutMultiplier = 0
	timeouts.WriteTotalTimeoutConstant = 0
	timeouts.WriteTotalTimeoutMultiplier = 0
	//Public Declare Auto Function SetCommTimeouts Lib "kernel32.dll" (ByVal hFile As IntPtr, _
	//   ByRef lpCommTimeouts As COMMTIMEOUTS) As Boolean
	r, _, err := syscall.Syscall(nSetCommTimeouts, 2, uintptr(h), uintptr(unsafe.Pointer(&timeouts)), 0)

	fmt.Println("SetCommTimeouts : ", r, err)

	if r == 0 {
		_, _, err := syscall.Syscall(nGetLastError, 1, uintptr(h), 0, 0)
		fmt.Println("LastError :", err)
		return nil
	}
	return nil
}

func setupComm(h syscall.Handle, in, out int) error {
	// Public Const OPEN_EXISTING As Int32 = 3
	// Public Declare Auto Function SetCommState Lib "kernel32.dll" (ByVal nCid As IntPtr, _
	// ByRef lpDCB As DCB) As Boolean

	r, _, err := syscall.Syscall(nSetupComm, 3, uintptr(h), uintptr(in), uintptr(out))

	fmt.Println("SetComm : ", r, err)

	if r == 0 {
		_, _, err := syscall.Syscall(nGetLastError, 1, uintptr(h), 0, 0)
		fmt.Println("LastError :", err)
		return nil
	}
	return nil
}

func setCommMask(h syscall.Handle) error {
	const EV_RXCHAR = 0x0001
	r, _, err := syscall.Syscall(nSetCommMask, 2, uintptr(h), EV_RXCHAR, 0)

	fmt.Println("SetCommMask : ", r, err)

	if r == 0 {
		_, _, err := syscall.Syscall(nGetLastError, 1, uintptr(h), 0, 0)
		fmt.Println("LastError :", err)
		return nil
	}
	return nil
}

func resetEvent(h syscall.Handle) error {
	r, _, err := syscall.Syscall(nResetEvent, 1, uintptr(h), 0, 0)

	fmt.Println("SetCommMask : ", r, err)

	if r == 0 {
		_, _, err := syscall.Syscall(nGetLastError, 1, uintptr(h), 0, 0)
		fmt.Println("LastError :", err)
		return nil
	}
	return nil
}

func newOverlapped() (*syscall.Overlapped, error) {
	var overlapped syscall.Overlapped
	r, _, err := syscall.Syscall6(nCreateEvent, 4, 0, 1, 0, 0, 0, 0)
	if r == 0 {
		return nil, err
	}
	overlapped.HEvent = syscall.Handle(r)
	return &overlapped, nil
}

func getOverlappedResult(h syscall.Handle, overlapped *syscall.Overlapped) (int, error) {
	var n int
	r, _, err := syscall.Syscall6(nGetOverlappedResult, 4,
		uintptr(h),
		uintptr(unsafe.Pointer(overlapped)),
		uintptr(unsafe.Pointer(&n)), 1, 0, 0)
	if r == 0 {
		return n, err
	}

	return n, nil
}
