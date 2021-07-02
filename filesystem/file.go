package filesystem

import (
	"io/ioutil"
	"os"
)

var ()

/*
	Return the file (param) content in a string.
*/
func FileReadAll(filepath string) (string, error) {
	contents, err := ioutil.ReadFile(filepath)
	return string(contents), err
}

/*
	Write a string in a file.
*/
func FileWriteAll(filepath, content string) error {
	return ioutil.WriteFile(filepath, []byte(content), 0664)
}

/*
	Test if a dir exists.
	(if the target is a file, the function returns false)
*/
func DirExists(path string) bool {
	if info, err := os.Stat(path); err == nil {
		if info.IsDir() {
			return true
		}
	}
	return false
}
