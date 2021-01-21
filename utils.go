package networkservices

//
//import (
//	"core/services/win32esk"
//	"encoding/binary"
//	"fmt"
//	"syscall"
//	"unsafe"
//)
//
//var (
//	ipHelperMod = syscall.NewLazyDLL("Iphlpapi.dll")
//
//	procGetTCPTable2 = ipHelperMod.NewProc("GetTcpTable2")
//)
//
//func GetTCPTable2() *MIB_TCPTABLE2 {
//	var n uint32
//	err, _, _ := procGetTCPTable2.Call(uintptr(unsafe.Pointer(&MIB_TCPTABLE2{})), uintptr(unsafe.Pointer(&n)), 1)
//	if syscall.Errno(err) != syscall.ERROR_INSUFFICIENT_BUFFER {
//		fmt.Printf("Error calling GetTcpTable2: %v", syscall.Errno(err))
//	}
//
//	b := make([]byte, n)
//	err, _, _ = procGetTCPTable2.Call(uintptr(unsafe.Pointer(&b[0])), uintptr(unsafe.Pointer(&n)), 1)
//	if err != 0 {
//		fmt.Printf("Error calling GetTcpTable2: %v", syscall.Errno(err))
//	}
//	readerRes := NewClassReader(b)
//	table := NewTCPTable(readerRes)
//	fmt.Println(table)
//	fmt.Println("GetTCPTable2 Call End")
//	return table
//}
//
//func NewClassReader(byteCode []byte) *ClassReader {
//	fmt.Println("NewClassReader", byteCode)
//	return &ClassReader{Bytecode: byteCode}
//}
//
//func (this *ClassReader) ReadUint32() uint32 {
//	value := bigEndian.Uint32(this.Bytecode[:4])
//	this.Bytecode = this.Bytecode[4:]
//	return value
//}
//
//func (this *ClassReader) ReadBytes(len int) []byte {
//	bytes := this.Bytecode[:len]
//	this.Bytecode = this.Bytecode[len:]
//	return bytes
//}
//
//func (this *ClassReader) ReadIp(bytes []byte) string {
//	return fmt.Sprintf("%d.%d.%d.%d", bytes[0], bytes[1], bytes[2], bytes[3])
//}
//
//func (this *ClassReader) ReadPort(bytes []byte) uint16 {
//	return binary.BigEndian.Uint16(bytes[0:2])
//}
//
//type MIB_TCPROW2 struct {
//	DwState        win32esk.DWORD
//	DwLocalAddr    win32esk.DWORD
//	DwLocalPort    win32esk.DWORD
//	DwRemoteAddr   win32esk.DWORD
//	DwRemotePort   win32esk.DWORD
//	DwOningPid     win32esk.DWORD
//	DwOffLoadState win32esk.TCP_CONNECTION_OFFLOAD_STATE
//}
//
//type MIB_TCPTABLE2 struct {
//	DwNumEntries win32esk.DWORD
//	Table        []*MIB_TCPROW2
//}
//
//func (r *MIB_TCPROW2) DisplayIP(val win32esk.DWORD) string {
//	return fmt.Sprintf("%d.%d.%d.%d", byte(val), byte(val>>8), byte(val>>16), val>>24)
//}
//
//func (r *MIB_TCPROW2) DisplayPort(val win32esk.DWORD) uint16 {
//	return binary.BigEndian.Uint16([]byte{byte(val), byte(val >> 8)})
//}
//
//func NewTCPRow(r *ClassReader) *MIB_TCPROW2 {
//	return &MIB_TCPROW2{win32esk.DWORD(r.ReadUint32()), win32esk.DWORD(r.ReadUint32()), win32esk.DWORD(r.ReadUint32()), win32esk.DWORD(r.ReadUint32()), win32esk.DWORD(r.ReadUint32()), win32esk.DWORD(r.ReadUint32()), win32esk.TCP_CONNECTION_OFFLOAD_STATE(r.ReadUint32())}
//}
//
//func NewTCPTable(r *ClassReader) *MIB_TCPTABLE2 {
//	t := &MIB_TCPTABLE2{}
//	fmt.Println("NewTCPTable", r)
//	t.DwNumEntries = win32esk.DWORD(r.ReadUint32())
//	table := make([]*MIB_TCPROW2, t.DwNumEntries)
//	for i := uint32(0); i < uint32(t.DwNumEntries); i++ {
//		table[i] = NewTCPRow(r)
//	}
//	t.Table = table
//	return t
//}
//
//func CloseTCPEntry(row *MIB_TCPROW2) error {
//	row.DwState = 12
//	if err, _, _ := syscall.NewLazyDLL("Iphlpapi.dll").NewProc("SetTcpEntry").Call(uintptr(unsafe.Pointer(row))); err != 0 {
//		return syscall.Errno(err)
//	}
//	return nil
//}
