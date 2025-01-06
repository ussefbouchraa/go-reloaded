package main

import (
	"fmt"
	"strconv"
	s "strings"
)

// func main() {

// 	slice := "qb_cdefghijkl;nopkK"
// //qbcd_c
// 	res := slice[:]
// 	fmt.Println(res)

// }

// func main() {
// slices := []string{"aaaa","bbbb","cccc","dddd","eeeee"}
// fmt.Println(len(slices))
// fmt.Println(cap(slices))
// 	res := append(slices[0: 1],  slices[4:]...)
// 	fmt.Println(res)

// // }

// func main() {
//     s := "  ---Hello Geeks---  "
//     // result := strings.TrimPrefix(s, "---")
//     // result = strings.TrimSuffix(result, "---")
//     result := strings.TrimSpace(s)
//     fmt.Println(result)
// }

//switch statement

// func main() {
// 	str := "jkggfn relkjshjhk vrl (up) kygfhjmg"
// 	help := s.Split(str, " ")
// 	for i := 0 ; i < len(help); i++ {
// 			switch help[i] {
// 				case "(up," :

// 					help = append(help[:i], help[i+1:]...)
// 				}
// 	}
//     fmt.Println(_hundleQuotes("xxxxx ' awesome ' "))
// }

// func main(){
// 	fmt.Println(_hundleQuotes("'   xxxx  ' '"))
// }


func _isExist( strs []string , target string) bool{
	if len(strs) == 0 || len(target) == 0{
		return false
	}
	for _, val := range(strs){
		if val == target{
			return true
		}
	}
	return false
}

func Capitalize(str string) string {
	return s.ToUpper(str[:1]) + s.ToLower(str[1:])
}

func  _hundleFlags(inp string) string{
	tokens := []string{"(hex,", "(bin,", "(up,", "(low,", "(cap,"}
	res := ""
	splitedInp := s.Fields(inp)
		
		for i:= 0; i < len(splitedInp);i++ {
			if !_isExist(tokens, splitedInp[i]){
				res += splitedInp[i] + " "
				continue
			}
		
			if len(splitedInp) > i + 1 && len(splitedInp[i + 1]) >= 2 && splitedInp[i + 1][len(splitedInp[i + 1]) - 1] == ')' {
				base ,err := strconv.Atoi(splitedInp[i+1][ :len(splitedInp[i+1]) - 1])
					if err != nil || base <= 0 {
						res += splitedInp[i] + " "
						continue
					}
						if base > i{
							base = i
						}
						prev := s.Join(splitedInp[i-base : i], " ")
						switch splitedInp[i] {
							case "(up,":
								res = s.Replace(res, prev, s.ToUpper(prev), base)
							case "(low,":
								res = s.Replace(res, prev, s.ToLower(prev), base)
							case "(cap,":
								res = s.Replace(res, prev, Capitalize(prev), base)
							// case "(bin,":
							// 	res = s.Replace(res, prev, Capitalize(prev), base)
							// case "(hex,":
							// 	res = s.Replace(res, prev, Capitalize(prev), base)
							}	
						i++
				}else{
					res += splitedInp[i] + " "
				}
			}
	return res
}

func main(){

	fmt.Println(_hundleFlags("zxxxx (cap, -2) xxxxxxxx  yyyyy"))
}
