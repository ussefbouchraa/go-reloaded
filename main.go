package main

import (
	// f  "fmt"
	os "os"
	s "strings"
)


// type S_tokens struct {
// 	hex string
// 	bin string
// 	up string
// 	low string
// 	cap string
// }

type S_command struct {
	buff 	string
	flag 	string 
	number 	uint  
}


func _checkExtention(arg string) bool {
	if len(arg) > 4{
		if arg[len(arg) - 4 :] == ".txt"{
			return true
		}
	}
	return false
}

func _checkArgs(args ...string) bool {
	
	if len(args[0]) == 0 || len(args[1]) == 0 {
		os.Stderr.WriteString("ERR : Invalid Argument !\n")
		return false
	}
	if !_checkExtention(args[0]) || !_checkExtention(args[1]) {
		os.Stderr.WriteString("ERR : Invalid Format !")
		return false
	}
	if args[0] == args[1]{
		os.Stderr.WriteString("ERR : Duplicated Files !\n")
		return false
	}
	
	return true
}

func _checkFiles(files ...string) bool{	
	
	for i := 0 ; i < len(files) ; i++{
		_, err := os.Stat(files[i])
		if err != nil{
			os.Stderr.WriteString(err.Error())
			return false
		}
	}
	return true
}


func _convertToBytes(strs []string) []byte{

	var bytes []byte 
		for i := 0 ; i < len(strs); i++{
			bytes = append(bytes, []byte(strs[i] + " ")...)
		}
	return bytes
}
  
func _hundleVowel(inp string) string{

	vowel, res := "aeiouhAEIOUH" , ""
	for i, val := range(inp){
		if (val == 'A' || val == 'a') && len(inp) > i + 2 && inp[i + 1] == ' ' { 
				if (s.Contains(vowel, string(inp[i + 2]))) {
					res += string(inp[i]) + "n"
					continue
				}
			}
		res += string(inp[i])
	}
	return res
}


func  _spliting(inp string) string{
	tokens, res := []string{"(hex,", "(bin,", "(up,", "(low,", "(cap,"}, ""	
	splitedInp := s.Fields(inp)
		
			for i:= 0; i < len(splitedInp) ; i++{
			for _,tkn := range (tokens){
				if s.Contains(splitedInp[i], tkn) && len(splitedInp) > i + 1 {
					res += splitedInp[i] + splitedInp[i + 1] + " "
					i+=2
					continue
				}
			}
			res += splitedInp[i] + " "
		}
	return res
}


func _addSuffix( input string) string{

	tokens := []string{"(hex)", "(bin)", "(up)", "(low)", "(cap)"}	
		for _, val := range(tokens){
			if (s.Contains(input, val)){
				token := "(" + val[1 : len(val) - 1] + ", 1)" 
				input = s.Replace(input, val, token, -1)
			}
		}
		return input 
}

func _trim( inp string) string{
	res := ""
	for i, val := range(inp){
		if val == ' ' && len(inp) > i + 1  && inp[i + 1] == ' '{
			continue
		}else{
			res += string(val)
		}
	}

	return res
	

}

func _hundlePunct(inp string) string {
	slices := []rune(_trim(inp))
	puncts := ".,!?:;"

	for i := 0; i < len(slices) ; i++{
		if slices[i] == ' ' && len(slices) > i + 1 {
			if slices[i + 1] == ' '{
				slices = append(slices[: i] , slices[i+1:]...)			
				i--
				continue
			}
			if s.Index(puncts,string(slices[i+1])) != -1 {
				slices[i] ,slices[i + 1] = slices[i + 1], slices[i] 
			}
		}
	}
		return s.TrimRight(string(slices), " ")
}

	
func _parseFileInp(input string ) []byte{	
	
	if len(input) == 0{
		os.Stderr.WriteString("ERR : Empty File !\n")
		os.Exit(1)
	}

	// ---- addingSuffix
	input = _addSuffix(input)
	// f.Println(input)
	
	// ----- handle the vowel
	input = _hundleVowel(input)
	// f.Println(input)

	// ----- custom spliting 
	input = _spliting(input)
	// f.Println(input)
	
	// ------ hundle punctuations
	input = _hundlePunct(input)
	return []byte(input)

	// ----- insert slice of structs 

	return _convertToBytes(s.Fields(input))
}


func _launching(files ...string){

	data , err := os.ReadFile(files[0])
	if err != nil{
		os.Stderr.WriteString(err.Error() + "\n")		
		return
	}
	
	newbuffer := _parseFileInp(string(data))
	
	err = os.WriteFile(files[1], []byte(s.TrimRight(string(newbuffer), " ")), 0777)
	if err != nil{
		os.Stderr.WriteString(err.Error() )
		return
	}
}

func main(){
	args := os.Args[1:]
	if len(args) != 2{
		os.Stderr.WriteString("ERR : Not Enough Parameters !\n")
		return 
	}
	if !_checkArgs(args[0], args[1]) || !_checkFiles(args[0], args[1]) {
		return	
	}

	_launching(args[0], args[1])

}