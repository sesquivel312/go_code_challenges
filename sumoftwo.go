/*
	Write a function that takes an array of numbers (integers for the tests) and
	a target number. It should find two different items in the array that, when
	added together, give the target value. The indices of these items should then
	be returned in a tuple like so: (index1, index2).

	For the purposes of this kata, some tests may have multiple answers; any valid
	solutions will be accepted.  The input will always be valid (numbers will be an
	array of length 2 or greater, and all of the items will be numbers; target will
	always be the sum of two different items from that array).

	Based on: http://oj.leetcode.com/problems/two-sum/

	twoSum [1, 2, 3] 4 === (0, 2)

	My Thoughts:
		nested loops can't be an efficient way to do this.  Given the above, there's
		nothing preventing the given array from containing values greater than the
		target number - i.e. don't need to look at that.  Regardless sorting might
		improve performance?

 */
package main

import (
		"fmt"
)

/*
For each value in the array, in index order, check it against the remaining
values in the array.  Once a value at an index is checked, it doesn't need
to be checked again, i.e.
	check v1 against v2, ... (v1 = value @ index 1, etc.)
	check v2 against v3, .. (don't go back)
*/
func TwoSum(number []int, target int) [2]int {

	for i := 0; i < len(number); i++ {  // could have used range here

		for j:=i+1; j<len(number); j++ {  //

			if number[i] + number[j] == target {
				return [2]int{i, j,}

			}

		}

	}

	return [2]int{-1,-1}
}

func TwoSum2(numbers []int, target int) [2]int {
	var complements = make(map[int]int)
	for i, number := range numbers {
		if index, ok := complements[number]; ok {
			return [2]int{index, i}
		}else{
			/*
			this bit is creating a map using the complements as keys and the
			index of the value that generated that complement as values.  Then
			when it checks the map (if part) if the complement is in the map at
			the right index then it has the index of two values that satisfy the
			problem constraints - from the current index of range (i) and the
			value at map[value @ numbers[i]
			todo time the two methods on large random arrays
			 */
			complements[target-number] = i
			fmt.Println(complements)
		}
	}
	return [2]int{}
}


func main () {
	nums := []int{23,12,19,78,22,5,9,}
	fmt.Println(TwoSum(nums, 83))
	fmt.Println(TwoSum2(nums, 83))
}