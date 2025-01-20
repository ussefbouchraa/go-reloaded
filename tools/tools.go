package piscine

import (
	"strconv"
	s "strings"
)

func ToUpper(slic []string) []string {
	if len(slic) == 0{
		return []string{}
	}
	for i, str := range slic {
		slic[i] = s.ToUpper(str)
	}
	return slic
}

func ToLower(slic []string) []string {
	if len(slic) == 0{
		return []string{}
	}
	for i, str := range slic {
		slic[i] = s.ToLower(str)
	}
	return slic
}

func Capitalize(slic []string) []string {
	if len(slic) == 0{
		return []string{}
	}
	for i, ss := range slic {
		str := []rune(ss)
		j := 0 ;
		for len(str) > j && str[j] == '\n' {
			j++
		}
		if len(str) > j {
			slic[i] = string(str[ : j]) + s.ToUpper(string(str[j])) + s.ToLower(string(str[j+1 :]))
		}
	}
	return slic

}

func HandleUniFlags(splitedInp []string, i int) ([]string, int) {
	var val int64 = 0
	var err error = nil

	if i-1 >= 0 {
		if splitedInp[i] == "(bin)" {
			val, err = strconv.ParseInt(splitedInp[i-1], 2, 0)
		} else if splitedInp[i] == "(hex)" {
			val, err = strconv.ParseInt(splitedInp[i-1], 16, 0)
		}

		if err == nil {
			splitedInp = append(splitedInp[:i-1], append([]string{strconv.Itoa(int(val))}, splitedInp[i+1:]...)...)
			i -= 2
			return splitedInp, i
		}
	}
	splitedInp = append(splitedInp[:i], splitedInp[i+1:]...)
	i--

	return splitedInp, i
}