package main

import (
	"fmt"
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



func _hundleQuots(str string) string {
	if s.Count(str, "'") <= 1{
		return "Error  Not Enough Quots "
	}
	start := s.Index(str, "'")
	end := s.Index(str[start + 1: ], "'" )
	
	res :=  s.TrimSpace(str[start + 1 : end + 1])
	// return  s.Replace(str, str[start + 1 : end + 
}
func main() {
    fmt.Println(_hundleQuots(" ' awesome '"))
}
