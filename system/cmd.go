// +build windows
package system

import (
        "bytes"
        "os/exec"
	"strings"
)

func CmdOut(name string, arg ...string) (string, error) {
        cmd := exec.Command(name, arg...)
        var out bytes.Buffer
        cmd.Stdout = &out
        err := cmd.Run()
        return out.String(), err
}
func CmdOutNoLn(name string, arg ...string) (out string, err error) {
        out, err = CmdOut(name, arg...)
        if err != nil {
                return
        }

        return strings.TrimSpace(string(out)), nil
}
