package filesystem

import (
	"os/exec"
	"strings"
	"syscall"
)

var (
	ReportedFSTypes = []int64{
		EXT4_SUPER_MAGIC,
		NTFS_SB_MAGIC}
)

func GetFileSystems() ([]string, error) {
	var fs []string
	var buffer syscall.Statfs_t

	out, err := exec.Command("sh", "-c", "df").Output()

	if err != nil {
		return fs, err
	} else {
		for _, i := range strings.Split(string(out), "\n") {
			j := strings.Fields(i)
			if len(j) > 5 {
				tmp := j[5]
				err = syscall.Statfs(tmp, &buffer)
				if IsReportedFSTypes(buffer.Type) {
					fs = append(fs, tmp)
				}
			}
		}
	}
	return fs, nil
}

func IsReportedFSTypes(fsType int64) bool {
	for _, i := range ReportedFSTypes {
		if fsType == int64(i) {
			return true
		}
	}
	return false
}

func GetFileSystemType(typeFs int64) string {
	var typeString string
	switch typeFs {
	case EXT4_SUPER_MAGIC:
		typeString = "ext4"
	case NTFS_SB_MAGIC:
		typeString = "ntfs"
	}
	return typeString
}
