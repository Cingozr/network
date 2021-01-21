package networkservices

import (
	"fmt"
	"strings"
	"syscall"
	"unicode/utf16"
	"unsafe"
)

func GetInternetProcess() {
	fmt.Println("GetInternetProcess")
	processNameMap := GetProcesNameMap()
	table := getTCPTable()
	fmt.Println(table)
	for i := uint32(0); i < uint32(table.dwNumEntries); i++ {
		row := table.table[i]
		ip := row.displayIP(row.dwRemoteAddr)
		port := row.displayPort(row.dwRemotePort)
		if row.dwOwningPid <= 0 {
			continue
		}

		if port != 80 && port != 443 {
			continue
		}
		process := strings.ToLower(processNameMap[uint32((row.dwOwningPid))])
		fmt.Printf("TCP Connections: \n Process= %v,\n Pid = %v,\n Addr = %v:%v \n", process, row.dwOwningPid, ip, port)
	}

}

func GetProcesNameMap() map[uint32]string {
	snapShot, err := syscall.CreateToolhelp32Snapshot(syscall.TH32CS_SNAPPROCESS, 0)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer syscall.CloseHandle(snapShot)

	var procEntry syscall.ProcessEntry32
	procEntry.Size = uint32((unsafe.Sizeof(procEntry)))
	if err = syscall.Process32First(snapShot, &procEntry); err != nil {
		fmt.Println(err)
		return nil
	}
	processNameMap := make(map[uint32]string)
	for {
		processNameMap[procEntry.ProcessID] = parseProcessName(procEntry.ExeFile)
		if err = syscall.Process32Next(snapShot, &procEntry); err != nil {
			if err == syscall.ERROR_NO_MORE_FILES {
				return processNameMap
			}
			fmt.Printf("Fail to syscall processNext: %v", err)
			return nil
		}
	}
}

func parseProcessName(exeFile [syscall.MAX_PATH]uint16) string {
	for i, v := range exeFile {
		if v <= 0 {
			return string(utf16.Decode(exeFile[:i]))
		}
	}
	return ""
}
