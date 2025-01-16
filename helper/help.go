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

func Split(s string) []string{
	str:= []rune(s)
	data := ""
	slic := []string{}

	if len(s) == 0{
		return slic
	}
	for _, char := range str {
		if char == ' ' { 
			if data != "" {
				slic = append(slic, data)
				data = ""
			}
		} else {
			data += string(char) 
		}
	}
	if data != ""{
		slic = append(slic, data)
	}
	return slic
}

func ToUpper(slic []string) []string {
	for i, str := range slic {
		slic[i] = s.ToUpper(str)
	}
	return slic
}

func ToLower(slic []string) []string {
	for i, str := range slic {
		slic[i] = s.ToLower(str)
	}
	return slic
}
func Capitalize(slic []string) []string {
	for i, str := range slic {
		slic[i] = s.ToUpper(string(str[0])) + s.ToLower(str[1 :])
	}
	return slic

}