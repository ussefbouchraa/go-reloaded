package goreload

import (
	h "goreload/helper"
	t "goreload/tools"
	"strconv"
	s "strings"
)

func AddSuffix(input string) string {
	tokens := []string{" (up) ", " (low) ", " (cap) "}
	for _, val := range tokens {
		if s.Contains(input, val) {
			token := val[:len(val)-2] + ", 1) "
			input = s.Replace(input, val, token, -1)
		}
	}
	return input
}

func HandleVowel(inp string) string {
	vowel, res := "aeiouhAEIOUH", ""
	inpp := []rune(inp)
	for i, val := range inpp {
		if (val == 'A' || val == 'a') && len(inpp) > i+2 && inpp[i+1] == ' ' {
			if s.Contains(vowel, string(inpp[i+2])) {
				res += string(inpp[i]) + "n"
				continue
			}
		}
		res += string(inpp[i])
	}
	return res
}

func HandleQuotes(inp string) string {
	if s.Count(inp, "'") <= 1 {
		return inp
	}
	i, start := 0, 0
	for i < len(inp) {
		if inp[i] == 39 && len(inp) > i+1 {
			start = i
			end := s.Index(inp[start+1:], "'")
			if end == -1 {
				break
			}
			end += start + 1
			trimmed :=  "'" + s.TrimSpace(inp[start+1 : end]) + "' "
			inp = s.Replace(inp, inp[start:end + 1], trimmed , -1)
			i += end 
		}
		i++
	}
	return inp
}

func HandlePunct(inp string) string {
	slices := []rune(inp)
	puncts := ".,!?:;"

	for i := 0; i < len(slices); i++ {
		if slices[i] == ' ' && len(slices) > i+1 {
			if slices[i+1] == ' ' {
				slices = append(slices[:i], slices[i+1:]...)
				i--
				continue
			}
			if s.Index(puncts, string(slices[i+1])) != -1 {
				slices[i], slices[i+1] = slices[i+1], slices[i]
			}
		}
	}
	return s.TrimRight(string(slices), " ")
}


func HandleFlags(inp string) string {
	tokens:= []string{"(hex)", "(bin)", "(up,", "(low,", "(cap,"}
	splitedInp := s.Split(inp, " ")

	for i := 0; i < len(splitedInp); i++ {
		if !h.IsExist(tokens, splitedInp[i]) {
			continue
		}

		if splitedInp[i] == "(bin)" || splitedInp[i] == "(hex)" {
			splitedInp, i = t.HandleUniFlags(splitedInp, i)
			continue
		}

		if len(splitedInp) > i+1 && len(splitedInp[i+1]) >= 2 && s.HasSuffix(splitedInp[i+1], ")") {

			base, err := strconv.Atoi(splitedInp[i+1][:len(splitedInp[i+1])-1])
			if  err != nil || base <= 0 {
				continue
			}

			prevItems, size := h.GetPrevious(splitedInp, i, base)
			switch splitedInp[i] {
			case "(up,":
				splitedInp = append(splitedInp[:i-size], append(t.ToUpper(prevItems), splitedInp[i+2:]...)...)
			case "(low,":
				splitedInp = append(splitedInp[:i-size], append(t.ToLower(prevItems), splitedInp[i+2:]...)...)
			case "(cap,":
				splitedInp = append(splitedInp[:i-size], append(t.Capitalize(prevItems), splitedInp[i+2:]...)...)
			}
			i -= 2
		}
	}
	return s.Join((splitedInp), " ")
}
