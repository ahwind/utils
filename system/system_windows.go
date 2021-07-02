// +build windows

package system

import (
	"log"
	"strings"
)

type Win32_BaseBoard struct {
	Product string
}

type Win32_BIOS struct {
	SerialNumber string
	Version      string
}

type Win32_OperatingSystem struct {
	Caption string
	//CSDVersion string
}

type Win32_ComputerSystem struct {
	Manufacturer string
	Model        string
}

type Win32_ComputerSystemProduct struct {
	UUID string
}

type Win32_PhysicalMemory struct {
	Name         string
	Capacity     int
	MemoryType   int
	Manufacturer string
}

type Win32_LogicalDisk struct {
	Name      string
	MediaType int
	DriveType int
	Size      int
}

type Win32_Processor struct {
	Name                      string `json:"name"`
	NumberOfCores             int    `json:"numcores"`
	NumberOfLogicalProcessors int    `json:"numlogical"`
}

func GetOS() string {

	var dst []Win32_OperatingSystem
	q := "SELECT Caption FROM Win32_OperatingSystem"
	err := Query(q, &dst)
	if err != nil {
		return ""
	}
	return dst[0].Caption
}

func GetOST() string {

	var dst []Win32_OperatingSystem
	q := "SELECT Caption, CSDVersion FROM Win32_OperatingSystem"
	err := Query(q, &dst)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	return ""
}

func GetBios() []Win32_BIOS {

	var dst []Win32_BIOS
	q := "SELECT SerialNumber, Version FROM Win32_BIOS WHERE (SerialNumber IS NOT NULL)"
	err := Query(q, &dst)
	if err != nil {
		return dst
	}
	return dst

}

func Getuuid() string {

	var dst []Win32_ComputerSystemProduct
	q := "SELECT UUID FROM Win32_ComputerSystemProduct"
	err := Query(q, &dst)
	if err != nil {
		return ""
	}
	return dst[0].UUID

}

func GetProductName() string {

	var dst []Win32_BaseBoard
	q := "SELECT Product FROM Win32_BaseBoard"

	err := Query(q, &dst)
	if err != nil {

		return ""
	}
	return dst[0].Product

}

func GetComputerSystem() []Win32_ComputerSystem {

	var dst []Win32_ComputerSystem
	q := "SELECT Manufacturer, Model FROM Win32_ComputerSystem"
	err := Query(q, &dst)
	if err != nil {
		return dst
	}
	return dst

}

func IsVirtual(Version, Model string) string {
	var ModelLower string

	ModelLower = strings.ToLower(Model)

	if strings.Contains(ModelLower, "vmware") {
		return "vmware"
	} else if strings.Contains(ModelLower, "openstack") {
		return "openstack"
	} else if strings.Contains(Version, "VRTUAL") {
		return "hyperv"
	} else {
		return "physical"
	}

}

func GetMem() []Win32_PhysicalMemory {
	var dst []Win32_PhysicalMemory
	q := "SELECT Name, Capacity, MemoryType, Manufacturer FROM Win32_Physicalmemory"
	err := Query(q, &dst)
	if err != nil {
		return dst
	}
	return dst

}

func CustomTrim(s string) string {
	return strings.Trim(strings.Trim(s, ":"), ",")
}

func GetDisk() []Win32_LogicalDisk {
	var dst []Win32_LogicalDisk
	q := "SELECT Name, MediaType, DriveType, Size FROM Win32_LogicalDisk"
	err := Query(q, &dst)
	if err != nil {
		return dst
	}
	return dst
}

func CpuModel() []Win32_Processor {

	var dst []Win32_Processor
	q := "SELECT Name, NumberOfCores, NumberOfLogicalProcessors FROM Win32_Processor"

	err := Query(q, &dst)
	if err != nil {
		return dst
	}
	return dst
}
