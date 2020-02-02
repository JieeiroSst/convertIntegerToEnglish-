package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("vim-go")
}
func numberToWords(num int) string {
	to19 := []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve", ",Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
	tens := []string{"Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
	if num == 0 {
		return ""
	}
	if num < 20 {
		return to19[num-1]
	}
	if num < 100 {
		return tens[num/10-2] + " " + numberToWords(num%10)
	}
	if num < 1000 {
		return to19[num/100-1] + " Hundred " + numberToWords(num%100)
	}

	for idx, w := range []string{"Thousand", "Million", "Billion"} {
		p := idx + 1
		if num < pow(1000, (p+1)) {
			return numberToWords(num/pow(1000, p)) + " " + w + " " + numberToWords(num%pow(1000, p))
		}
	}
	return "bug"
}

func pow(i int, p int) int {
	return int(math.Pow(1000, float64(p)))
}
