package day25

import (
	"math"
)

var snafuToDec = map[string]int{
	"=": -2,
	"-": -1,
	"0": 0,
	"1": 1,
	"2": 2}

var decToSnafu = map[int]string{
	-4: "1",
	-3: "2",
	-2: "=",
	-1: "-",
	0:  "0",
	1:  "1",
	2:  "2",
	3:  "=",
	4:  "-",
	5: "0"}

func ToDec(snafu string) int {
	col := 0
	sum := 0

	snafu = reverse(snafu)
	for i := range snafu {
		sum += int(math.Pow(5, float64(col))) * snafuToDec[string(snafu[i])]
		col++
	}

	return sum
}

func ToSnafu(dec int) string {
	num := dec
	result := ""
	carry := 0

	for num > 0 {
		r := num % 5 + carry
		num = num / 5

		if r > 2 {
			carry = 1
		} else {
			carry = 0
		}

		result += decToSnafu[r]
	}

	if carry > 0 {
		result += decToSnafu[carry]
	}
	return reverse(result)
}

func SumSnafu(list []string) string {
	var result string

	for index := range list {
		list[index] = reverse(list[index])
	}

	for cont, index, carry := true, 0, 0; cont; index++ {
		cont = false

		sum := carry
		for _, num := range list {
			if len(num) > index {
				sum += snafuToDec[string(num[index])]
				cont = true
			}
		}

		v := sum / 5
		r := sum % 5

		if r == 3 || r == 4 {
			v++
		}

		if r < 0 {
			v--
		}

		carry = v
		if cont || r > 0 {
			result += decToSnafu[r]
		}

	}

	return reverse(result)
}

func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}
