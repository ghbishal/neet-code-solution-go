// 238: Product of Array Except Self
// https://leetcode.com/problems/product-of-array-except-self/

package main

import "fmt"

// SOLUTION

// Solving method:

/*
# Approach

*/

// productExceptSelf calculates the product of all elements in the array except itself.

func productExceptSelf(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)

	// Step 1: Fill in the prefix products
	// Initialize the first element of answer as 1 since there's no element to the left of the first element
	answer[0] = 1
	for i := 1; i < n; i++ {
		// answer[i] contains the product of all elements to the left of nums[i]
		answer[i] = answer[i-1] * nums[i-1]
	}

	// Step 2: Fill in the suffix products and multiply with the prefix products
	// Initialize a variable to store the running product of elements to the right of the current element
	suffix := 1
	for i := n - 1; i >= 0; i-- {
		// Multiply the current value in answer (which is the prefix product) by the suffix product
		answer[i] *= suffix
		// Update suffix to include nums[i] for the next iteration
		suffix *= nums[i]
	}

	return answer

}

/*
# Time Complexity
- O(n) where n is the length of the input array.
We iterate through the array twice, once to calculate the prefix products and once to calculate the suffix products.


# Space Complexity
- O(1) since we only use a single array to store the output and a few variables for calculations.
*/

func main() {

	// Example 1
	input1 := []int{1, 2, 3, 4}
	result1 := productExceptSelf(input1)
	fmt.Println(result1) // Output: [24,12,8,6]

	// Example 2
	input2 := []int{-1, 1, 0, -3, 3}
	result2 := productExceptSelf(input2)
	fmt.Println(result2) // Output: [0,0,9,0,0]

}
