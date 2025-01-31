package goreload

import (
	F "fmt"
	"strings"
)

const ES_Green = " \033[32m"
const ES_Reset = "\033[37m"

func PrintHeader(){
	F.Println(ES_Green + "  ############################################## " + ES_Reset)
	F.Println(ES_Green + "  #                                           	# " + ES_Reset)
	F.Println(ES_Green + "  #          <<--- Go_reloaded --->>           # " + ES_Reset)
	F.Println(ES_Green + "  #                                           	# " + ES_Reset)
	F.Println(ES_Green + "  ############################################## " + ES_Reset)
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

func handleWhiteSpaces(inp string)string{
	WhiteSpaces := []string {"\t", "\r", "\v","\f"}
	for _,v := range WhiteSpaces{
		inp = strings.Replace(inp, v, " ", -1)
	}
	return inp
}

func Trim(inp string) string {
	res := ""
	inp = handleWhiteSpaces(inp)
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

func isNewLine(str string) bool{
	if len(str) == 0{
		return false
	}
	for _, val := range(str){
		if val != '\n'{
			return false
		}
	}
	return true
}

func GetPrevious(slic []string, indx int, size int) ([]string, int) {
	res , rev := []string{} ,[]string{}

	if size > indx {
		size = indx
	}

	if len(slic) == 0 || indx <= 0 || size >= len(slic) {
		return res, size
	}

	for i := indx - 1; i >= 0 &&  len(res) < size; i-- {
		if isNewLine(slic[i]) {
			size++
		}
		res = append(res, slic[i])
	}

	for j:= len(res) - 1 ; j >= 0; j--{
		rev = append(rev, res[j])
	}

	return rev, len(rev)
}