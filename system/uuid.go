package system

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"net"
)

func ID() string {

	var uuid string
	intfs, _ := net.Interfaces()
	for _, i := range intfs {
		log.Println(i.HardwareAddr.String())
		uuid += i.HardwareAddr.String()
	}
	hasher := md5.New()
	hasher.Write([]byte(uuid))
	return hex.EncodeToString(hasher.Sum(nil))
}

func Uuid() string {
	var uuid string
	intfs, _ := net.Interfaces()
	for _, i := range intfs {
		uuid += i.HardwareAddr.String()
	}
	hasher := md5.New()
	hasher.Write([]byte(uuid))
	return hex.EncodeToString(hasher.Sum(nil))[:16]

}
