// 217: Contains Duplicate
// https://leetcode.com/problems/contains-duplicate/

package main

import "fmt"

// SOLUTION

// Solving method: Hash Map

// containsDuplicate determines if there are any duplicates in the given array.
func containsDuplicate(nums []int) bool {

	// Create a set (using a map with an empty struct as values) to store unique elements.
	// We use a map here because checking for existence in a map has an average O(1) time complexity.
	set := make(map[int]struct{})

	// Loop through each number in the array
	for _, num := range nums {
		// Check if the current number already exists in the set
		if _, hasNum := set[num]; hasNum {
			// If the number is already in the set, it means we've found a duplicate
			return true
		}
		// Otherwise, add the number to the set
		// Using struct{}{} as the value, which takes up zero memory space
		set[num] = struct{}{}

	}
	// If we loop through the entire array and find no duplicates, return false
	return false
}

/*
# Time Complexity
O(n), where n is the length of the nums array. We iterate over nums once, and each lookup and insertion in the map has an average time complexity of O(1).
Therefore, the entire solution runs in linear time.

# Space Complexity
O(n) as well, because in the worst case (where there are no duplicates), we need to store every element of nums in the map.
*/

func main() {
	// INPUT

	nums := []int{1, 2, 3, 1}

	// OUTPUT
	result := containsDuplicate(nums)
	fmt.Println(result)
}
