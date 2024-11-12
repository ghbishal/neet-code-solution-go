// 242. Valid Anagram
// https://leetcode.com/problems/valid-anagram/

package main

import "fmt"

// SOLUTION

// Solving method: Hash Map (or Frequency Counting)

// isAnagram determines if two strings are anagrams of each other.
func isAnagram(s string, t string) bool {

	// Step 1: Check if the strings are of the same length.
	// If not, they cannot be anagrams, so we return false immediately.
	if len(s) != len(t) {
		return false
	}

	// Step 2: Initialize a map to keep track of the frequency of each character in `s`.
	// `Seen` will store each character in `s` as a key, with its value being the count of occurrences.
	// rune is an alias for int32 and is commonly used to represent a Unicode code point.
	seen := make(map[rune]int)

	// Step 3: Iterate through the first string `s`, adding each character to the `seen` map.
	// Each character's count is incremented by 1 for every occurrence.
	for _, char := range s {
		seen[char]++
	}
	// Step 4: Iterate through the second string `t`.
	// For each character in `t`, decrement the corresponding count in the `seen` map.
	for _, char := range t {
		seen[char]--
	}

	// Step 5: Check if all counts in the `seen` map are zero.
	// If any value is not zero, it means `s` and `t` do not have the same characters in the same counts.
	// so we return false in the case.
	for _, value := range seen {
		if value != 0 {
			return false
		}
	}

	return true
}

/*
#Time Complexity

- Length Check: Checking len(s) != len(t) takes constant time, O(1).
- Building the Map (seen): The first for loop iterates over each character in s, which takes O(n) time, where n is the length of s.
- Updating the Map with t: The second for loop iterates over each character in t, which also takes O(n) time.
- Final Map Check: The final for loop checks each value in the seen map. Since the map has at most 26 keys (for lowercase alphabetic characters), this check is effectively O(1) (constant time) because 26 is a fixed upper bound.
- Overall Time Complexity: O(n) since we process each string a single time in linear time.

# Space Complexity
- Auxiliary Data Structure (seen): The seen map stores up to 26 entries (for lowercase letters). Since this is a constant upper limit, the space complexity for the seen map is O(1).
- Overall Space Complexity: O(1).
*/

func main() {
	// INPUT
	s := "anagram"
	t := "nagaram"

	// OUTPUT
	result := isAnagram(s, t)
	fmt.Println(result)
}
