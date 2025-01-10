package main

// import "fmt"

// import (
// 	"fmt"
// 	"strconv"
// 	s "strings"
// )

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

// func HandleFlags(inp string) string {
// 	tokens := []string{"(hex,", "(bin,", "(up,", "(low,", "(cap,"}
// 	splitedInp := s.Fields(inp)

// 	for i := 0; i < len(splitedInp); i++ {
// 		if !IsExist(tokens, splitedInp[i]) {
// 			continue
// 		}
// 		if len(splitedInp) > i+1 && len(splitedInp[i+1]) >= 2 && s.HasSuffix(splitedInp[i + 1], ")"){

// 			base, err := strconv.Atoi(splitedInp[i+1][:len(splitedInp[i+1])-1])
// 			if err != nil || base <= 0 {
// 				i++
// 				continue
// 			}
// 			if base > i {base = i}
// 			prev := s.Join(splitedInp[i-base:i], " ")
// 			PrevWithFlag := s.Join(splitedInp[i - base : i + 2], " ")

// 			switch splitedInp[i] {
// 				case "(up,":
// 					inp = s.Replace(inp, PrevWithFlag, s.ToUpper(prev), base)
// 				case "(low,":
// 					inp = s.Replace(inp, PrevWithFlag, s.ToLower(prev), base)
// 				case "(cap,":
// 					inp = s.Replace(inp, PrevWithFlag, Capitalize(prev), base)
// 				case "(bin,":
// 					val, err := strconv.ParseInt(prev, 2, 0)
// 					if err != nil {
// 						i++
// 						break
// 					}
// 					inp = s.Replace(inp, PrevWithFlag,  strconv.Itoa(int(val)), base)
// 				case "(hex,":
// 					val, err := strconv.ParseInt(prev, 16, 0)
// 					if err != nil {
// 						i++
// 						break
// 					}
// 					inp = s.Replace(inp, PrevWithFlag, strconv.Itoa(int(val)), base)
// 				}
// 			}
// 		}
// 	return inp
// }

// const ES_Green = " \033[32m"
// const ES_Reset = " \033[37m"

// func main(){
// 	fmt.Println(ES_Green + " " + "<<--- Go_reloaded --->>" + ES_Reset  )
// 	// fmt.Println(HandleFlags(" (low, 1) (cap, 1)"))      
// }
