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




func IsExist( strs []string , target string) bool{
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
	strs := []rune(str)
	return s.ToUpper(string(strs[:1])) + s.ToLower(string(strs[1:]))
}

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


func HandleFlags(inp string) string {
	tokens := []string{"(hex,", "(bin,", "(up,", "(low,", "(cap,"}
	splitedInp := s.Fields(inp)

	for i := 0; i < len(splitedInp); i++ {
		if !IsExist(tokens, splitedInp[i]) {
			continue
		}
		if len(splitedInp) > i+1 && len(splitedInp[i+1]) >= 2 && s.HasSuffix(splitedInp[i + 1], ")"){

			base, err := strconv.Atoi(splitedInp[i+1][ : len(splitedInp[i+1])-1])
			if err != nil || base <= 0 {
				i++
				continue
			}
			if base > i {base = i}
			prevItems := s.Join(splitedInp[i-base : i], " ")
			if prevItems == "" {
				splitedInp[i], splitedInp[i + 1] = "", "" 
				i++
				continue
			}

			switch splitedInp[i] {
				case "(up,":
						splitedInp = append(splitedInp[ : i-base], append([]string{s.ToUpper(prevItems)}, splitedInp[i + 2: ]...)... )
				case "(low,":
						splitedInp = append(splitedInp[ : i-base], append([]string{s.ToLower(prevItems)}, splitedInp[i + 2: ]...)... )
				case "(cap,":
					splitedInp = append(splitedInp[ : i-base], append([]string{Capitalize(prevItems)}, splitedInp[i + 2: ]...)... )
				case "(bin,":
					val, err := strconv.ParseInt(prevItems, 2, 0)
					if err == nil {
						splitedInp = append(splitedInp[ : i-base], append([]string{strconv.Itoa(int(val))}, splitedInp[i + 2: ]...)... )
					}
					i++
				case "(hex,":
					val, err := strconv.ParseInt(prevItems, 16, 0)
					if err == nil {
						splitedInp = append(splitedInp[ : i-base], append([]string{strconv.Itoa(int(val))}, splitedInp[i + 2: ]...)... )
					}
					i++
				}
			}
		}
	return s.Join(splitedInp, " ")
}

// func main(){

// 	fmt.Println(HandleFlags("10 (bin, 1) , 1)"))
// }
