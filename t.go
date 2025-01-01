
//1------ falgs --> (cap) || (cap, 1)
flags := []{"(cap,", "(low,"}


//2------ split  [ (cap,|1) ] 


//3------ parsing
loop str{
	if strings.contain(flags, str[i] ){
		if len(str[i + 1]) == 3 {
			 strconv.Atoi(str[i+1][0]) == nil && str[i + 1][1] == ")"    
		}else{
			return str	
		}
	}
}