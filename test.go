package main

import (
	"fmt"
	// "os"
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

func _hundleQuotes(str string) string {
	if s.Count(str, "'") <= 1{
		return str
	}
	i, start := 0 , 0
	for  i < len(str) {
		if str[i] == 39 && len(str) > i + 1{
			start = i
			end := s.Index(str[start+1 : ], "'" )
			if  end == -1{
				break
			}
			end += start + 1
			trimmed := s.TrimSpace(str[start + 1 : end])
			str = s.Replace(str, str[start + 1 : end],  trimmed , -1)
			i = end 
		}
			i++
	}
	return str
}

func main(){
	fmt.Println(_hundleQuotes("'   xxxx  ' '"))
}



