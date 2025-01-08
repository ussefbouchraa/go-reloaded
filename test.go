package main

import (
	// "fmt"
	// "strconv"
	// s "strings"
	// "unsafe"
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




// func _isExist( strs []string , target string) bool{
// 	if len(strs) == 0 || len(target) == 0{
// 		return false
// 	}
// 	for _, val := range(strs){
// 		if val == target{
// 			return true
// 		}
// 	}
// 	return false
// }

// func Capitalize(str string) string {
// 	return s.ToUpper(str[:1]) + s.ToLower(str[1:])
// }

// func  _hundleFlags(inp string) string{
// 	tokens := []string{"(hex,", "(bin,", "(up,", "(low,", "(cap,"}
// 	res := inp
// 	splitedInp := s.Fields(inp)
		
// 		for i:= 0; i < len(splitedInp);i++ {
// 			if !_isExist(tokens, splitedInp[i]){
// 				continue
// 			}
// 			if len(splitedInp) > i + 1 && len(splitedInp[i + 1]) >= 2 &&
// 				splitedInp[i + 1][len(splitedInp[i + 1]) - 1] == ')' {

// 				base ,err := strconv.Atoi(splitedInp[i+1][ :len(splitedInp[i+1]) - 1])
// 				if err != nil || base <= 0 {
// 					i++
// 					continue
// 				}
// 				if base > i { base = i }
// 				prev := s.Join(splitedInp[i-base : i ], " ")
// 				oldflag := s.Join(splitedInp[i-base : i+2], " ")
// 				switch splitedInp[i] {
// 					case "(up,":
// 							res = s.Replace(res, oldflag, s.ToUpper(prev), base)
// 							break
// 						case "(low,":
// 							res = s.Replace(res, prev, s.ToLower(prev), base)
// 							break
// 						case "(cap,":
// 							res = s.Replace(res, prev, Capitalize(prev), base)
// 							break
// 						case "(bin,":
// 							val, err := strconv.ParseInt(prev, 2,  len(prev)*8)
// 							if(err != nil){
// 								i++
// 								break
// 							}
// 							toString :=  strconv.Itoa(int(val))
// 							res = s.Replace(res, prev, toString, base)
// 							break
// 						case "(hex,":
// 							val, err := strconv.ParseInt(prev, 10,  len(prev)*8)
// 							if(err != nil){
// 								i++
// 								break
// 							}
// 							toString :=  strconv.Itoa(int(val))
// 							res = s.Replace(res, prev, toString, base)
// 						}
// 					}
// 				}
// 	return res
// }

// func main(){
// s := []rune("🙂")
// fmt.Println(len(s))
// fmt.Println(cap(s))
// fmt.Println((s))
// 	// fmt.Println(_hundleFlags(" 🙂gg (cap, 2) It has been xx (up, 1) years(up, 1)"))
// }
