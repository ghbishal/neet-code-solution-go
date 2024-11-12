// 1: Two Sum
// https://leetcode.com/problems/two-sum/

package main

import "fmt"

// SOLUTION

// Solving method: Hash Map

/*
# Approach
- Brute Force Approach: A simple approach would be to check every pair of numbers to see if they add up to the target, but that would be inefficient (O(n^2) time complexity).

# Optimal Approach Using a Hash Map:

- As we traverse the array, for each number, we calculate the complement (i.e., target - num).
- We check if this complement is already in a hash map. If it is, that means we have already seen the number that, together with the current number, adds up to the target.
- If the complement is not found, we store the current number and its index in the hash map.
*/

// twoSum finds the indices of the two numbers that add up to a target sum.
func twoSum(nums []int, target int) []int {

	// Create a map to store the number and its index
	numMap := make(map[int]int)

	// Loop through each number in the array
	for i, num := range nums {
		// Calculate the complement we need to find.
		complement := target - num

		// Check if the complement is already in the map
		if index, found := numMap[complement]; found {
			// If found, return the pairs of indices
			return []int{i, index}
		}

		// Otherwise, store the number and its index in the map
		numMap[num] = i

		fmt.Println(numMap)

	}

	// Return an empty slice if no solution is found
	return nil

}

/*
# Time Complexity
- O(n) where n is the length of the nums array. We iterate over nums once, and each lookup and insertion in the map has an average time complexity of O(1).
- Therefore, the entire solution runs in linear time.

# Space Complexity
- O(n) as well, because in the worst case (where there are no duplicates), we need to store every element of nums in the map.
*/

func main() {

	// INPUT

	nums := []int{0, 7, 11, 15, 2, 3}
	target := 9

	// OUTPUT
	result := twoSum(nums, target)
	fmt.Println(result)
}
