package main

var wordMap = map[int]string{
	1:   "one",
	2:   "two",
	3:   "three",
	4:   "four",
	5:   "five",
	6:   "six",
	7:   "seven",
	8:   "eight",
	9:   "nine",
	10:  "ten",
	11:  "eleven",
	12:  "twelve",
	15:  "quarter",
	20:  "twenty",
	25:  "twenty five",
	30:  "half",
	100: "midnight",
}

func getMin(i int) string {
	return wordMap[i]
}

func getHours(i int) string {
	if val, ok := wordMap[i]; ok {
		return val
	}
	return "midnight"
}
