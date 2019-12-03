package main

import "strconv"

func limitLength(s string, length int) string {
	if len(s) < length {
		return s
	}

	return s[:length]
}

//https://stackoverflow.com/a/24973180/2030623
func Slice_Atoi(strArr []string) ([]int, error) {
	// NOTE:  Read Arr as Slice as you like
	var str string // O
	var i int      // O
	var err error  // O

	iArr := make([]int, 0, len(strArr))
	for _, str = range strArr {
		i, err = strconv.Atoi(str)
		if err != nil {
			return nil, err // O
		}
		iArr = append(iArr, i)
	}
	return iArr, nil
}

func Slice_Itoa(arr []int) []string {
	// NOTE:  Read Arr as Slice as you like
	var str string
	var i int

	strArr := make([]string, 0, len(arr))
	for _, i = range arr {
		str = strconv.Itoa(i)
		strArr = append(strArr, str)
	}
	return strArr
}
