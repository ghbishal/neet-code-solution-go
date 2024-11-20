// 00659 Encode and decode strings
// https://www.lintcode.com/problem/659/
// https://leetcode.com/problems/encode-and-decode-strings/

package main

import (
	"fmt"
	"strings"
)

// SOLUTION

// Solving method: Delimiter Separated Encoding

/*
# Approach
This method can be classified as a "Delimiter Encoding" approach. We essentially transform
a list of strings into a single string using a delimiter, which can later be split
back into the original list. This approach is efficient and works well given the constraints.
*/

// Encode function: Takes a list of strings and returns a single encoded string.
func Encode(str []string) string {
	// Use strings.Join to join each string in the slice with a unique delimiter
	// Here, we use "#!" as a delimiter since it's unlikely to be in the input strings.
	// This helps ensure the delimiter won't cause errors during decoding.
	return strings.Join(str, "#!")
}

// Decode function: Takes the encoded string and decodes it back into a list of strings.
func Decode(s string) []string {
	// Use strings.Split to split the encoded string by the delimiter we used in Encode.
	// This will return an array of original strings.
	if s == "" {
		return []string{} // Handle the edge case of an empty input string
	}

	return strings.Split(s, "#!")
}

/*

# Time Complexity
- Encode : O(n * m) where n is the number of strings and m is the average length of the strings.
- Decode : O(n * m) where n is the number of strings in the encoded and m is the average length of the strings.


# Space Complexity
- Encode: O(n⋅m), as we create a new encoded string with all characters from the input.
- Decode: O(n⋅m), since we produce a new slice of strings, with each element requiring additional memory.
*/

func main() {
	// Test case 1
	encoded := Encode([]string{"neet", "code", "love", "you"})
	fmt.Println("Encoded:", encoded) // Shows the encoded format
	decoded := Decode(encoded)
	fmt.Println("Decoded:", decoded) // Should output the original array

	// Test case 2
	encoded = Encode([]string{"we", "say", ":", "yes"})
	fmt.Println("Encoded:", encoded)
	decoded = Decode(encoded)
	fmt.Println("Decoded:", decoded)
}
