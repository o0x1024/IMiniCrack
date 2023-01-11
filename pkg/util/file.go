package util

import (
	"os"
	"strings"
)

func CheckFileIsExist(filename string) bool {
	var flag = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		flag = false
	}
	return flag
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func GetParentDirectory(dirctory string) string {
	if strings.LastIndex(dirctory, "/") == -1 {
		return substr(dirctory, 0, strings.LastIndex(dirctory, "\\"))

	}
	return substr(dirctory, 0, strings.LastIndex(dirctory, "/"))

}

func GetFileName(path string) string {
	if strings.LastIndex(path, "/") == -1 {
		return substr(path, strings.LastIndex(path, "\\")+1, len(path))
	}
	return substr(path, strings.LastIndex(path, "/")+1, len(path))
}

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}
