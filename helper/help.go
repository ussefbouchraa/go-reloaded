package piscine

import (
	s "strings"
	f "fmt"
)

const ES_Green = " \033[32m"
const ES_Reset = "\033[37m"

func PrintHeader(){
	f.Println(ES_Green + "  ############################################## " + ES_Reset)
	f.Println(ES_Green + "  #                                           	# " + ES_Reset)
	f.Println(ES_Green + "  #          <<--- Go_reloaded --->>           # " + ES_Reset)
	f.Println(ES_Green + "  #                                           	# " + ES_Reset)
	f.Println(ES_Green + "  ############################################## " + ES_Reset)
}

func Capitalize(str string) string {
	strs := []rune(str)
	i := 0
	for strs[i] == '\n'{
		i++
	}
	return s.ToUpper(string(strs[ : i + 1])) + s.ToLower(string(strs[i + 1:]))
	
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
