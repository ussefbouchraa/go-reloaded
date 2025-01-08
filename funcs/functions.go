package piscine

import (
	h "piscine/helper"
	"strconv"
	s "strings"
)

func AddSuffix(input string) string {
	tokens := []string{"(hex)", "(bin)", "(up)", "(low)", "(cap)"}
	for _, val := range tokens {
		if s.Contains(input, val) {
			token := val[:len(val)-1] + ", 1)"
			input = s.Replace(input, val, token, -1)
		}
	}
	return input
}

func HundleVowel(inp string) string {
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

func HundleQuotes(inp string) string {
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
			trimmed := s.TrimSpace(inp[start+1 : end])
			inp = s.Replace(inp, inp[start+1:end], trimmed, -1)
			i = end
		}
		i++
	}
	return inp
}

func HundlePunct(inp string) string {
	slices := []rune(h.Trim(inp))
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

func HundleFlags(inp string) string {
	tokens := []string{"(hex,", "(bin,", "(up,", "(low,", "(cap,"}
	res := inp
	splitedInp := s.Fields(inp)

	for i := 0; i < len(splitedInp); i++ {
		if !h.IsExist(tokens, splitedInp[i]) {
			continue
		}
		if len(splitedInp) > i+1 && len(splitedInp[i+1]) >= 2 &&
			splitedInp[i+1][len(splitedInp[i+1])-1] == ')' {

			base, err := strconv.Atoi(splitedInp[i+1][:len(splitedInp[i+1])-1])
			if err != nil || base <= 0 {
				i++
				continue
			}
			if base > i {
				base = i
			}
			prev := s.Join(splitedInp[i-base:i], " ")
			oldflag := s.Join(splitedInp[i-base:i+2], " ")
			switch splitedInp[i] {
			case "(up,":
				res = s.Replace(res, oldflag, s.ToUpper(prev), base)
				break
			case "(low,":
				res = s.Replace(res, prev, s.ToLower(prev), base)
				break
			case "(cap,":
				res = s.Replace(res, prev, h.Capitalize(prev), base)
				break
			case "(bin,":
				val, err := strconv.ParseInt(prev, 2, len(prev)*8)
				if err != nil {
					i++
					break
				}
				toString := strconv.Itoa(int(val))
				res = s.Replace(res, prev, toString, base)
				break
			case "(hex,":
				val, err := strconv.ParseInt(prev, 10, len(prev)*8)
				if err != nil {
					i++
					break
				}
				toString := strconv.Itoa(int(val))
				res = s.Replace(res, prev, toString, base)
			}
		}
	}
	return res
}
