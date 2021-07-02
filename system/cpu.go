package system

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"runtime"
	"strings"

	"github.com/toolkits/file"
)

/*
	Return system load average
*/
func GetNumCPU() float64 {
	return float64(runtime.NumCPU())
}

/*
	Return program's architecture
*/
func GetArch() string {
	return runtime.GOARCH
}

func CpuMHz() (mhz string, err error) {
	f := "/proc/cpuinfo"
	var bs []byte
	bs, err = ioutil.ReadFile(f)
	if err != nil {
		return
	}

	reader := bufio.NewReader(bytes.NewBuffer(bs))

	for {
		var lineBytes []byte
		lineBytes, err = file.ReadLine(reader)
		if err == io.EOF {
			return
		}

		line := string(lineBytes)
		if !strings.Contains(line, "MHz") {
			continue
		}

		arr := strings.Split(line, ":")
		if len(arr) != 2 {
			return "", fmt.Errorf("%s content format error", f)
		}

		return strings.TrimSpace(arr[1]), nil
	}

	return "", fmt.Errorf("no MHz in %s", f)
}
