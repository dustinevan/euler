package main

import (
	//"github.com/dustinevan/euler/common"
	"fmt"
	//"math"
	"github.com/dustinevan/euler/common"
)

func main() {

	//Numerator Tens Digit
	for a := 1; a < 10; a++ {
		//Numerator Ones Digit
		for b := 1; b < 10; b++ {
			//Denominator Tens Digit
			for c := 1; c < 10; c++ {
				//Denominator Ones Digit
				for d := 1; d < 10; d++ {
					if (a * 10 + b)/(c * 10 + d) < 1 {
						flag, num, den := isCuriousFraction(a, b, c, d)
						if flag {
							fmt.Println(a, b, "/", c, d, " : ", num, "/", den)
						}

					}
				}
			}
		}
	}

}

func isCuriousFraction(a, b, c, d int) (flag bool, num int, den int) {
	if a == c && factoring.CompareIntFractions(b, d, a*10+b, c*10+d) {
		return true, b, d
	} else if a == d && factoring.CompareIntFractions(b, c, a*10+b, c*10+d) {
		return true, b, c
	} else if b == c && factoring.CompareIntFractions(a, d, a*10+b, c*10+d) {
		return true, a, d
	} else if b == d && factoring.CompareIntFractions(a, c, a*10+b, c*10+d) {
		return true, a, c
	} else {
		return false, 0, 0
	}
}