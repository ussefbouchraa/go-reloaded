package main

import (
	os "os"
	f "piscine/funcs"
	h "piscine/helper"
	s "strings"
)

func _parseFileInp(input string) []rune {
	if len(input) == 0 {
		os.Stderr.WriteString("ERR : Empty File !\n")
		os.Exit(0)
	}
	funcs := []func(string) string{f.AddSuffix, f.HandleVowel, f.HandleQuotes, f.HandleFlags, f.HandlePunct}
	for _, f := range funcs {
		input = f(input)
	}
	return []rune(input)
}


func _launching(files ...string) {

	data, err := os.ReadFile(files[0])
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		return
	}

	newbuffer := _parseFileInp(string(data))

	err = os.WriteFile(files[1], []byte(s.TrimRight(string(newbuffer), " ")), 0777)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return
	}
}

func _checkFiles(files ...string) bool {
	for i := 0; i < len(files); i++ {
		_, err := os.Stat(files[i])
		if err != nil {
			os.Stderr.WriteString(err.Error())
			return false
		}
	}
	return true
}

func _checkArgs(args ...string) bool {

	if len(args[0]) == 0 || len(args[1]) == 0 {
		os.Stderr.WriteString("ERR : Invalid Argument !\n")
		return false
	}
	if !h.CheckExtention(args[0]) || !h.CheckExtention(args[1]) {
		os.Stderr.WriteString("ERR : Invalid Format !")
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
	if !_checkArgs(args[0], args[1]) || !_checkFiles(args[0], args[1]) {
		return
	}

	_launching(args[0], args[1])
}
