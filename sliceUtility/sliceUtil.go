package slice

func IsPresent(s []string, e string) bool {
	for _, val := range s {
		if val == e {
			return true
		}
	}
	return false
}

func ElemNotPresent(s1 []string, s2 []string) []string {
	var sol []string
	for _, e := range s2 {
		//ans := IsPresent(s1, e)
		ans := false
		if !ans {
			sol = append(sol, e)
		}
	}
	return sol
}

func CopiedElements(s1 []string) []string {
	hash := make(map[string]int)
	var sol []string
	for _, v := range s1 {
		if hash[v] == 0 {
			hash[v] = 1
		} else if hash[v] == 1 {
			continue
		}
	}
	for k := range hash {
		sol = append(sol, k)
	}
	return sol
}
