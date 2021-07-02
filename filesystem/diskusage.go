package filesystem

import (
	"bytes"
	"encoding/json"
	"syscall"
)

/*
	Get disk usage with a path

	returns:

	[0] TotalStorage
	[1] TotalFileNodes
	[2] UsedStorage
	[3] FreeStorage
	[4] UsedFileNodes
	[5] FreeFileNodes
*/
func GetDiskUsage(path string) ([]float64, error) {
	var usage []float64 = make([]float64, 6) // Data to return
	var sc syscall.Statfs_t                  // Filesystem stat
	var err error                            // Error catching
	var buffer bytes.Buffer                  // Buffer for json indent
	var content []byte                       // Json's content

	// Get filesystem stat
	err = syscall.Statfs(path, &sc)
	if err != nil {
		return usage, err
	}

	// Convert structure => json
	content, err = json.Marshal(sc)
	if err != nil {
		return usage, err
	}

	// Indent json
	json.Indent(&buffer, content, "", "   ")

	// Set data to return
	usage[0] = float64(sc.Bsize) * float64(sc.Blocks) // TotalStorage
	usage[1] = float64(sc.Files)                      // TotalFileNodes
	usage[3] = float64(sc.Bsize) * float64(sc.Bfree)  // FreeStorage
	usage[2] = usage[0] - usage[3]                    // UsedStorage
	usage[5] = float64(sc.Ffree)                      // FreeFileNodes
	usage[4] = usage[1] - usage[5]                    // UsedFileNodes

	return usage, nil
}
