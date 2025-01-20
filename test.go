package main

import (
	// "fmt"
	// "fmt"
	"strconv"
	s "strings"
)

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



func Capitalize(str string) string {
	strs := []rune(str)
	return s.ToUpper(string(strs[:1])) + s.ToLower(string(strs[1:]))
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



func HandleFlags(inp string) string {
	
	tokens := []string{"(hex,", "(bin,", "(up,", "(low,", "(cap,"}
	splitedInp := s.Split(inp, " ")
	
	for i := 0 ; i < len(splitedInp); i++ {
		if !IsExist(tokens, splitedInp[i]) {
			continue
		}
		if len(splitedInp) > i+1 && len(splitedInp[i+1]) >= 2 && s.HasSuffix(splitedInp[i + 1], ")"){

			base, err := strconv.Atoi(splitedInp[i+1][ : len(splitedInp[i+1])-1])
			if err != nil || base <= 0 { continue }
			
			if base > i {base = i}
			prevItems := s.Join(splitedInp[i-base : i], " ")
			if prevItems == "" {
				splitedInp = append(splitedInp[ : i], splitedInp[i + 2: ]...) 
				i--
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
					if err != nil || splitedInp[i+1] != "1)" { continue }
					splitedInp = append(splitedInp[ : i-base], append([]string{strconv.Itoa(int(val))}, splitedInp[i + 2: ]...)... )
				case "(hex,":
					val, err := strconv.ParseInt(prevItems, 16, 0)
					if err != nil || splitedInp[i+1] != "1)"   { continue }
					splitedInp = append(splitedInp[ : i-base], append([]string{strconv.Itoa(int(val))}, splitedInp[i + 2: ]...)... )
		
				}
				i -= 2 
			}
		}
	return s.Join(splitedInp, " ")
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

// func GetPrevious(slic []string, indx int, size int) ([]string ,int) {
//     res := []string{}
//     cp := 0

// 	if size > indx{size = indx}
//     if len(slic) == 0 || indx <= 0 || size >= len(slic) {
//         return res, size
//     }
    
//     for _, val := range(slic[: indx]){
//         if isNewLine(val) {
// 			cp++
// 		}
//     }
// 	newsize := size + cp
// 	if (indx - newsize) < len(slic){
// 		return slic[ indx - newsize : indx] , newsize
// 	}

//     return slic[ indx - size : indx] , size
// }



func GetPrevious(slic []string, indx int, size int) ([]string, int) {
	res := []string{}

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

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res, len(res)
}

// func main() {
//     res := []string{"\n", "\n", "\n", "\n","(up, 2)"}
//     fmt.Println(GetPrevious(res, 4, 1))
// }