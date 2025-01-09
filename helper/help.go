package piscine

import (
	s "strings"
)
func Capitalize(str string) string {
	strs := []rune(str)
	return s.ToUpper(string(strs[:1])) + s.ToLower(string(strs[1:]))
}

func IsExist(strs []string, target string) bool {
	if len(strs) == 0 || len(target) == 0 {
		return false
	}
	for _, val := range strs {
		if val == target {
			return true
		}
	}
	return false
}

func Trim(inp string) string {
	res := ""
	inpp := []rune(inp)
	for i, val := range inpp {
		if val == ' ' && len(inpp) > i+1 && inpp[i+1] == ' ' {
			continue
		} else {
			res += string(val)
		}
	}
	return res
}

func CheckExtention(arg string) bool {
	if len(arg) > 4 {
		if arg[len(arg)-4:] == ".txt" {
			return true
		}
	}
	return false
}
