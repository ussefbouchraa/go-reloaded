package goreload

import (
	H "goreload/helper"
	T "goreload/tools"
	"strconv"
	S "strings"
)

func AddSuffix(input string) string {
	tokens := []string{" (up) ", " (low) ", " (cap) "}
	for _, val := range tokens {
		if S.Contains(input, val) {
			token := val[:len(val)-2] + ", 1) "
			input = S.Replace(input, val, token, -1)
		}
	}
	return input
}

func HandleVowel(inp string) string {
	vowel, res := "aeiouhAEIOUH", ""
	inpp := []rune(inp)
	for i, val := range inpp {
		if (val == 'A' || val == 'a') && len(inpp) > i+2 && inpp[i+1] == ' ' {
			if S.Contains(vowel, string(inpp[i+2])) {
				res += string(inpp[i]) + "n"
				continue
			}
		}
		res += string(inpp[i])
	}
	return res
}

func HandleQuotes(inp string) string {
	if S.Count(inp, "'") <= 1 {
		return inp
	}
	inp = S.Replace(inp, "' '", "''" , -1)
	i, start := 0, 0
	for i < len(inp) {
		if inp[i] == 39 && len(inp) > i+1 {
			start = i
			end := S.Index(inp[start+1:], "'")
			if end == -1 {
				break
			}
			end += start + 1
			trimmed :=  S.TrimSpace(inp[start+1 : end ])
			if trimmed != ""{
				inp = S.Replace(inp, inp[start+1 : end], trimmed , -1)
				i = end - 2
			}else{
				i++
			}
		}
		i++
	}

	return inp
}

func HandlePunct(inp string) string {
	slices := []rune(inp)
	puncts := ".,!?:;'"

	for i := 0; i < len(slices); i++ {
		if S.Contains(puncts, string(slices[i])) && len(slices) > i+1 &&  slices[i+1] != ' '{
			slices = append(slices[:i+1], append([]rune(" "), slices[i+1:]...)...)
			i++
		}
		if slices[i] == ' ' && len(slices) > i+1 {
			if slices[i+1] == ' ' {
				slices = append(slices[:i], slices[i+1:]...)
			}
			if S.Index(".,!?:;", string(slices[i+1])) != -1 {
				slices[i], slices[i+1] = slices[i+1], slices[i]
			}
			if slices[i+1] == '\n' {
				slices = append(slices[:i], slices[i+1:]...)
			}
		}
	}
	return string(slices)
}

func HandleFlags(inp string) string {
	tokens := []string{"(hex)", "(bin)", "(up,", "(low,", "(cap,"}
	splitedInp := S.Split(H.Trim(inp), " ")

	for i := 0; i < len(splitedInp); i++ {
		if !H.IsExist(tokens, splitedInp[i]) {
			continue
		}

		if splitedInp[i] == "(bin)" || splitedInp[i] == "(hex)" {
			splitedInp, i = T.HandleUniFlags(splitedInp, i)
			continue
		}

		if len(splitedInp) > i+1 && len(splitedInp[i+1]) >= 2 && S.HasSuffix(splitedInp[i+1], ")") {

			nbr, err := strconv.Atoi(splitedInp[i+1][:len(splitedInp[i+1])-1])
			if err != nil || nbr <= 0 {
				continue
			}

			prevItems, size := H.GetPrevious(splitedInp, i, nbr)
			switch splitedInp[i] {
			case "(up,":
				splitedInp = append(splitedInp[:i-size], append(T.ToUpper(prevItems), splitedInp[i+2:]...)...)
			case "(low,":
				splitedInp = append(splitedInp[:i-size], append(T.ToLower(prevItems), splitedInp[i+2:]...)...)
			case "(cap,":
				splitedInp = append(splitedInp[:i-size], append(T.Capitalize(prevItems), splitedInp[i+2:]...)...)
			}
			i -= 2
		}
	}
	return S.Join((splitedInp), " ")
}
