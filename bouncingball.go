package main

import (
	"fmt"
	"math"
)

func BouncingBall(hfloor, a, hwindow float64) int {
	/*
		hfloor is height of the floor from which ball is dropped
		hwindow	is height of window with observer - looking for bounces
		a is the bounce factor - and it must be < 1

		after looking at a few examples will realize that generally are looking for
		n where

		a^n <= hwindow/hfloor

		This can be solved w/logs:

		n <= log-a(r); r = hwindow/hfloor, log-w = log base w

		but, math packages probably don't have the ability to solve
		logs in arbitrary bases.  So, using the conversion of bases
		trick (which I totally didn't remember - had to google)
		you can solve it:

		log-b(x) = log-bb(x) / log-bb(b); log-w = log with base w

	*/

	// check for invalid params
	if hfloor > hwindow && hfloor > 0 && a > 0 && a < 1 {

		r := hwindow/hfloor

		n := math.Log10(r)/math.Log10(a)

		if math.Floor(n) == n {  // bounced exactly to window - which doesn't count according to the constraints given

			n--  // account for one too many bounces then

		}

		return 2 * int(math.Floor(n)) + 1


	} else {  // 1+ invalid params

		return -1
	}
}

func main() {
	/*
			testequal(3, 0.66, 1.5, 3)
		      testequal(40, 0.4, 10, 3)
		      testequal(10, 0.6, 10, -1)
		      testequal(40, 1, 10, -1)
		      testequal(5, -1, 1.5, -1)
	*/


	fmt.Println(BouncingBall(3, 0.66, 1.5))
	fmt.Println(BouncingBall(40, 0.4, 10))
	fmt.Println(BouncingBall(10, 0.6, 10))
	fmt.Println(BouncingBall(40, 1, 10))
	fmt.Println(BouncingBall(5, -1, 1.5))
	fmt.Println(BouncingBall(2, 0.25, 1))
	fmt.Println(BouncingBall(2, 0.5, 1))

	// now some random tests
	//r :=rand.New(rand.NewSource(19700920))
	//for i := 1; i < 100; i++ {
	//
	//	f := 100 * r.Float64()
	//	w := 80 * r.Float64()
	//	a := 3 * r.Float64() - 1.5
	//
	//	fmt.Println(BouncingBall(f, a, w))
	//}
}
