package primes

import "fmt"

var c *calculator

func init() {
	c = NewPrimeCalcualtor(0)
}

type (
	calculator struct {
		numberList []bool
		primeList []int
		highestPrime int
		internalLimit int
		currentLevel int
	}
)

func NewPrimeCalcualtor(limit int) *calculator {

	internalLimit := 1000000000
	numberList := make([]bool, internalLimit )
	numberList[0] = true //true means not prime in this system
	numberList[1] = true

	primeList := make([]int, 1)
	primeList[0] = 2

	highestPrime := 2

	c := calculator{numberList: numberList, primeList: primeList, highestPrime: highestPrime,
		internalLimit: internalLimit, currentLevel: 2}
	/*TODO: the limit doesn't work right now, it just calcs to the internal limit,
	 make this iterative. */
	c.eliminateNonPrimes(0)
	fmt.Println("Primes Done! Found all primes to ", internalLimit)
	return &c
}

func IsPrime(x int) bool {
	return c.isPrime(x)
}

func (c *calculator) isPrime(x int) bool {
	if x < c.internalLimit {
		return !c.numberList[x]
	}
	fmt.Println("!!!!ERROR!!!!", x, "is not inbounds")
	return false

}

func (c *calculator) FindPrimes(limit int) []int {
	c.eliminateNonPrimes(limit)
	return c.primeList
}

func FindPowerOfTwo(power int) int {
	num := 2

	for i := 2; i <= power; i++  {
		num *= 2
	}
	return num
}

func (c *calculator) eliminateNonPrimes(limit int) {
	for {
		for index, prime := range c.primeList {
			c.eliminate(index, prime, 2)
		}
		if !c.findNextPrime() {
			return
		} else {
			c.currentLevel++
		}
	}
}

func (c *calculator) eliminate(index, multiplier, level int) {
	for i := index; i < len(c.primeList); i++ {
		x := multiplier * c.primeList[i]
		if x < c.internalLimit {
			c.numberList[x] = true
			if level < c.currentLevel {
				c.eliminate(i, x, level+1)
			}
		} else {
			return
		}
	}
}

func (c *calculator) findNextPrime() bool{
 	highestPossiblePrime := FindPowerOfTwo(c.currentLevel)
	result := false
	for i := c.highestPrime + 1; i < highestPossiblePrime && i < c.internalLimit; i++ {
		if c.numberList[i] == false {
			c.primeList = append(c.primeList, i)
			c.highestPrime = i
			result = true
		}
	}
	return result
}