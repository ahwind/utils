package system

import (
	"strings"
)

func SliceContain(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}

	return false
}

func SliceContainLike(s []string, e string) bool {

	for _, a := range s {
		if strings.Contains(a, e) {
			return true
		}
	}

	return false
}
