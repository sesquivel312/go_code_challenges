/*
A child is playing with a ball on the nth floor of a tall building.
The height of this floor, h, is known. He drops the ball out of the
window. The ball bounces (for example), to two-thirds of its height
(a bounce of 0.66).

His mother looks out of a window 1.5 meters from the ground.

How many times will the mother see the ball pass in front of her
window (including when it's falling and bouncing?

Three conditions must be met for a valid experiment:

    Float parameter "h" in meters must be greater than 0
    Float parameter "bounce" must be greater than 0 and less than 1
    Float parameter "window" must be less than h.

If all three conditions above are fulfilled, return a positive integer, otherwise return -1.

Note:

    The ball can only be seen if the height of the rebounding ball is strictly greater than the window parameter.
        Example:
	    - h = 3, bounce = 0.66, window = 1.5, result is 3
	    - h = 3, bounce = 1, window = 1.5, result is -1

	    (Condition 2) not fulfilled)

*/

package main

import "fmt"

func bball(hfloor, a, hwin float64) int {

	/*
	   assuming the bounce height on bounce n is a * bounce height
	   on bounce n-1
	*/

	// first check that hfloor is > hwin - i.e. kid not dropping ball from
	// lower floor

	if a >= 1 {
		return -1
	} else if hfloor > hwin {  // dropping from higher floor whan window - proceed

		// let's brute force this first

		count := 1            // must see it on the way down at least
		r := hwin / hfloor  // just making things a littler easier to write

		bheight := a * hfloor // get the first bounce height

		for bheight > hwin {  // this is go's while loop syntax - ick - b/c we already have first bounce ht, this will fall through if that's not as high as the window

			bheight *= a  // new bounce height is a * last one

			count += 2
		}

		return count

	} else { // kid dropping from same or lower floor, won't ever be seen
		return -1
	}
}

func main(){

	// - h = 3, bounce = 0.66, window = 1.5, result is 3
	fmt.Println(bball(3, 0.66, 1.5))
	// - h = 3, bounce = 1, window = 1.5, result is -1
	fmt.Println(bball(3, 1, 1.5))
}