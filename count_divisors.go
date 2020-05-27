/*
Function should return the count of divisors of some integer n

Examples

	Divisors(4)  == 3  //  1, 2, 4
	Divisors(5)  == 2  //  1, 5
	Divisors(12) == 6  //  1, 2, 3, 4, 6, 12
	Divisors(30) == 8  //  1, 2, 3, 5, 6, 10, 15, 30

My approach is brute force.  The only simplifying assumption made is that the divisors
can't be > n/2.  I wonder also if this approach isn't flawed - i.e. will it find divisors that are
themselves factorable?
 */
package main

import "fmt"

func Divisors(n int) int {

	count := 1  // count of divisors

	for i := 1; i <= n/2 ; i++ {

		if n % i == 0 { // i is a divisor

			count++

		}
	}

	return count
}

/*
	Testing - create an array? of test values and result
 */


type testPair struct {
	value int
	expected int
}

func main() {

	tests := []testPair {
		{4, 3},
		{5, 2},
		{12, 6},
		{30, 8},
	}

	for _, pair := range tests {

		fmt.Printf("%d found %d divisors, expected %d\n", pair.value, Divisors(pair.value), pair.expected)

	}

}
