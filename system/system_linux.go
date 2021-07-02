// +build !windows
package system

import (
	"os/exec"
	"runtime"
	"strings"
	"syscall"
	"time"
)

/*
   Return system uptime
*/
func GetUptime() string {
	var sc syscall.Sysinfo_t
	var err error

	err = syscall.Sysinfo(&sc)
	if err != nil {
		return ""
	}

	return string(time.Duration(sc.Uptime) * time.Second)
}

/*
   get os
*/

func GetKernel() string {
	data, err := CmdOutNoLn("uname", "-r")
	if err != nil {
		return ""
	}
	return strings.Trim(data, "\n")
}

func GetOS() string {
	out, err := exec.Command("sh", "-c", "cat /etc/redhat-release").Output()
	if err != nil {
		return ""
	}
	return strings.Trim(strings.Trim(string(out), "\n"), " ")
}

func GetType() string {

	return runtime.GOOS
}

func GetManufacturer() string {
	out, err := exec.Command("sh", "-c", "/usr/sbin/dmidecode -s select-manufacturer").Output()
	if err != nil {
		return ""
	}
	return strings.Trim(strings.Trim(string(out), "\n"), " ")
}

func GetProductName() string {
	out, err := exec.Command("sh", "-c", "/usr/sbin/dmidecode -s system-product-name").Output()
	if err != nil {
		return ""
	}
	return strings.Trim(strings.Trim(string(out), "\n"), " ")
}

func Virtual(Model string) string {
	var ModelLower string

	ModelLower = strings.ToLower(Model)

	if strings.Contains(ModelLower, "vmware") {
		return "vmware"
	} else if strings.Contains(ModelLower, "openstack") {
		return "openstack"
	} else {
		return "physical"
	}

}

func GetMem() string {
	//out, err := exec.Command("sh", "-c", "/usr/sbin/dmidecode |grep -A16 'Memory Device$'|grep 'Size.*MB'|sed 's#^\sSize:\s##'").Output()
	out, err := exec.Command("sh", "-c", "/usr/sbin/dmidecode |grep -A16 'Memory Device$'|grep 'Size.*MB'|awk '{print $(NF-1), $NF}'").Output()
	if err != nil {
		return ""
	}
	return strings.Trim(strings.Trim(string(out), "\n"), " ")

}

func CustomTrim(s string) string {
	return strings.Trim(strings.Trim(s, ":"), ",")
}

func CpuModel() string {

	out, err := exec.Command("sh", "-c", "cat /proc/cpuinfo |grep 'model name'|uniq|awk -F: '{print $2}'").Output()
	//out, err := exec.Command("sh", "-c", "cat /proc/cpuinfo |grep 'model name'|uniq").Output()
	if err != nil {
		return ""
	}
	return strings.Trim(strings.Trim(string(out), "\n"), " ")
}

func Getuuid() string {

	out, err := exec.Command("sh", "-c", "dmidecode -s system-uuid").Output()
	if err != nil {
		return ""
	}
	return strings.Trim(strings.Trim(string(out), "\n"), " ")

}
