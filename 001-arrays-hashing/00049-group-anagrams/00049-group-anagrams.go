// 49: Group Anagrams
// https://leetcode.com/problems/group-anagrams/

package main

import (
	"fmt"
	"sort"
)

// SOLUTION

// Solving method: Hash Map combined with sorting

/*
# Approach
- Sort each string: The key idea is that two strings that are anagrams will have the same sorted string. For example, "eat", "tea", and "ate" will all become "aet" when sorted.
- HashMap (or Map in Go): Use a map to group strings that have the same sorted string as their key. The value of each key will be a slice containing all the strings (anagrams) that map to that key.
- Return the grouped anagrams: Finally, return the values of the map, which will contain the groups of anagrams.
*/

// groupAnagrams groups anagrams together in the given array of strings.

func groupAnagrams(strs []string) [][]string {

	// Step 1: Create a map to store the groups of anagrams
	anagramMap := make(map[string][]string)

	// Step 2: Iterate through each string in the input list
	for _, str := range strs {

		// Step 2a: Convert the string to a slice of runes (to handle multi-byte characters safely)
		chars := []rune(str)

		// Step 2b: Sort the characters of the string
		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})

		// Step 2c: Convert the sorted slice of runes back to a string
		sortedStr := string(chars)

		// Step 3: Add the original string to the anagram group corresponding to the sorted string
		anagramMap[sortedStr] = append(anagramMap[sortedStr], str)

	}

	// Step 4: Collect all the anagram groups

	var result [][]string
	for _, groups := range anagramMap {
		result = append(result, groups)
	}

	return result
}

/*

# Time Complexity
- Sorting each string takes O(k log k) time, where k is the length of the longest string in the input list.
- Iterating through each string in the list takes O(n) time, where n is the total number of strings.
- Therefore, the overall time complexity is O(n * k log k).


# Space Complexity
- The space complexity is O(n * k)  where n is the number of strings, and k is the average length of the strings because the map will store up to n groups of strings, and each group will contain a string of length up to k.
- Auxiliary space for sorting: Sorting each string requires O(k) space for the slice of characters.
*/

func main() {
	// Example 1
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := groupAnagrams(strs)
	fmt.Println(result) // Output: [["eat", "tea", "ate"], ["tan", "nat"], ["bat"]]

	// Example 2
	strs2 := []string{""}
	result2 := groupAnagrams(strs2)
	fmt.Println(result2) // Output: [[""]]

	// Example 3
	strs3 := []string{"a"}
	result3 := groupAnagrams(strs3)
	fmt.Println(result3) // Output: [["a"]]
}
