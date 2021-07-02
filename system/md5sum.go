package system

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"os/exec"
)

func Md5sumCheck(workdir, md5file string) bool {
	cmd := exec.Command("md5sum", "-c", md5file)
	cmd.Dir = workdir
	err := cmd.Run()
	if err != nil {
		log.Printf("cd %s; md5sum -c %s fail", workdir, md5file)
		return false
	}
	return true
}

func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))

	return hex.EncodeToString(h.Sum(nil))
}
