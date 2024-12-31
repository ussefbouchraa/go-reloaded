package main

import (
	"fmt"
	"os"
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
//     fmt.Println(_hundleQuots("xxxxx ' awesome ' "))
// }

func _hundleQuots(str string) string {
	if s.Count(str, "'") <= 1{
		return "Error  Not Enough Quots "
	}
	for i:= 0; i < len(str) ; i++{
		start := s.Index(str[i: ], "'")
		if start == -1{
			continue
		}
		end := s.Index(str[start + 1: ], "'" )
		if  end == -1{
			continue
		}
		end++
		fmt.Println( string(str[start]), string(str[end]) ,str[start : end] )
		os.Exit(0)
		trimmed := s.TrimSpace(str[start : end])
		str = s.Replace(str, str[start : end],  trimmed , -1)
		
		fmt.Println("---", str[start : end] , "-----" + trimmed)
		i = end
	}
	return str
}

func main(){
fmt.Println(_hundleQuots("' awesome ' "))


}
