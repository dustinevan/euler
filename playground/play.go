package main

import (
	"fmt"
	"github.com/dustinevan/euler/common/primes"
	"time"
	//"github.com/dustinevan/euler/common/operations"
)


func main() {

	now := time.Now()
	c := primes.NewPrimeCalcualtor(100)
	new := c.FindPrimes(1)
	elapsed := time.Since(now)
	fmt.Println(elapsed)

	/*now = time.Now()
	old := operations.GetAllPrimesBelow(1000000)
	elapsed = time.Since(now)*/

	fmt.Println(elapsed, len(new))
	//fmt.Println(len(new), len(old))
}


