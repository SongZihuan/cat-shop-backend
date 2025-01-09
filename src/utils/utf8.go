package utils

import "unicode/utf8"

func IsUTF8(b []byte) bool {
	return utf8.Valid(b)
}

func IsUTF8String(s string) bool {
	return utf8.ValidString(s)
}

func HasUTF8BOM(data []byte) bool {
	if len(data) >= 3 && data[0] == 0xEF && data[1] == 0xBB && data[2] == 0xBF {
		return true
	}
	return false
}

func RemoveBOMIfExists(data []byte) []byte {
	if HasUTF8BOM(data) {
		return data[3:]
	}
	return data
}
