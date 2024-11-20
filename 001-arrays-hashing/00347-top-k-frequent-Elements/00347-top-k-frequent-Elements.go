// 347. Top K Frequent Elements
// https://leetcode.com/problems/top-k-frequent-elements/

package main

import "fmt"

// SOLUTION

// Solving method: Hash Map for Counting and Heap for Sorting

/*
# Approach
This approach is a combination of Hash Map for Counting and Heap for Sorting by Frequency.
Specifically, it uses the Bucket Sort method, which is optimized for this kind of problem where
we need to retrieve the top k frequent elements quickly without sorting the entire array.
*/

// topKFrequent finds the k most frequent elements in the given array.
func topKFrequent(nums []int, k int) []int {

	// Step 1: Count the frequency of each elements using a map
	frequency := make(map[int]int)
	for _, num := range nums {
		frequency[num]++
	}

	// Step 2: Create a slices (buckets), where the index represents the frequency
	// The length of `buckets` is len(nums)+1, as frequencies range from 0 to len(nums)
	buckets := make([][]int, len(nums)+1)
	for num, freq := range frequency {
		buckets[freq] = append(buckets[freq], num)
	}

	// Step 3: Gather the top k frequent elements starting from the highest frequency bucket
	result := []int{}
	for i := len(buckets) - 1; i >= 0 && len(result) < k; i-- {
		// Add elements from the current bucket to the result until we have k elements
		if len(buckets[i]) > 0 {
			result = append(result, buckets[i]...)
		}
	}

	return result[:k]
}

/*

# Time Complexity
- The algorithm operates in O(n) time complexity, where n is the length of the nums array.
	- Counting frequency with a hash map is O(n)
	- Inserting elements into a bucket and flattening the bucket to find the top k elements is also O(n)

# Space Complexity
- The space complexity is also O(n) due to the hash map and bucket array used to store frequencies.
*/

func main() {
	// Example usage
	nums := []int{1, 1, 1, 2, 2, 4, 4, 5, 3}
	k := 3
	fmt.Println(topKFrequent(nums, k)) // Output: [1,2]
}
