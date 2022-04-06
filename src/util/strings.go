package util

import (
	"strings"
	"unsafe"
)

func Str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func Bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// IsEmpty ("") == true
// IsEmpty (" ") == false
// IsEmpty ("\n") == false
// IsEmpty ("\t") == false
func IsEmpty(s string) bool {
	return s == ""
}

// IsBlank ("") == true
// IsBlank (" ") == true
// IsBlank ("\n") == true
// IsBlank ("\t") == true
func IsBlank(s string) bool {
	temp := strings.Trim(s, " ")
	temp = strings.Trim(temp, "\n")
	temp = strings.Trim(temp, "\t")
	return IsEmpty(temp)
}
