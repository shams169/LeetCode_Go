package main

var RomanMap = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func RomanToInt(roman string) int {
	var result int

	if len(roman) == 0 {
		return 0
	}

	l := len(roman)
	result = RomanMap[string(roman[l-1])]

	for i := l - 2; i >= 0; i-- {
		if RomanMap[string(roman[i])] < RomanMap[string(roman[i+1])] {
			result = result - RomanMap[string(roman[i])]
		} else {
			result = result + RomanMap[string(roman[i])]
		}
	}

	return result
}

// func RomanToInt(roman string) int {
// 	var result int

// 	l := len(roman)

// 	result += RomanMap[string(roman[l-1])]

// 	for i := l - 2; i >= 0; i-- {
// 		if RomanMap[string(roman[i])] < RomanMap[string(roman[i+1])] {
// 			result = result - RomanMap[string(roman[i])]
// 		} else {
// 			result = result + RomanMap[string(roman[i])]
// 		}

// 	}
// 	//fmt.Println(result)
// 	return result
// }
