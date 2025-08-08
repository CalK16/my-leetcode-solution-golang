package main

import (
	"slices"
	"testing"
)

func sortStr(s string) string {
	bytes := []byte(s)
	slices.Sort(bytes)
	return string(bytes)
}

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, word := range strs {
		s := sortStr(word)
		m[s] = append(m[s], word)
	}
	ans := make([][]string, 0)
	for _, val := range m {
		ans = append(ans, val)
	}
	return ans
}

func TestGroupAnagrams(t *testing.T) {
	expected := [][]string{[]string{"bat"}, []string{"nat", "tan"}, []string{"ate", "eat", "tea"}}
	ans := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	if !isSame2DSlice(ans, expected) {
		t.Errorf("TEST 1 FAILED")

	}

	expected = [][]string{[]string{""}}
	ans = groupAnagrams([]string{""})
	if !isSame2DSlice(ans, expected) {
		t.Errorf("TEST 2 FAILED")
	}

}
