package main

import (
	os "os"
	F "goreload/funcs"
	H "goreload/helper"
	S "strings"
)

func _parseFileInp(input string) []rune {
	if len(input) == 0 {
		os.Stderr.WriteString("ERR : Empty File !\n")
		os.Exit(0)
	}
	funcs := []func(string) string{ F.AddSuffix, F.HandleVowel, F.HandleFlags, F.HandlePunct, F.HandleQuotes}
	for _, f := range funcs {
		input = f(input)
	}
	return []rune(input)
}

func _launching(files ...string) {

	H.PrintHeader()
	data, err := os.ReadFile(files[0])
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		return
	}
	newbuffer := _parseFileInp(string(data))

	err = os.WriteFile(files[1], []byte(S.TrimRight(string(newbuffer), " ")), 0777)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return
	}
}

func _checkArgs(args ...string) bool {

	if len(args[0]) == 0 || len(args[1]) == 0 {
		os.Stderr.WriteString("ERR : Invalid Argument !\n")
		return false
	}
	if !H.CheckExtention(args[0]) || !H.CheckExtention(args[1]) {
		os.Stderr.WriteString("ERR : Invalid File Format !")
		return false
	}
	if args[0] == args[1] {
		os.Stderr.WriteString("ERR : Duplicated Files !\n")
		return false
	}
	return true
}

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		os.Stderr.WriteString("ERR : Not Enough Parameters !\n")
		return
	}
	if !_checkArgs(args[0], args[1]) {
		return
	}

	_launching(args[0], args[1])

}
