package utils

import "strconv"

func StringSliceToIntSlice(s []string) []int {
	var i []int
	for _, v := range s {
		n, err := strconv.Atoi(v)
		if err == nil {
			i = append(i, n)
		}
	}
	return i
}
