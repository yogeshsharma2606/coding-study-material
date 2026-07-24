package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	resultMap := make(map[string][]string)

	for _, s := range strs {
		r := []rune(s)
		sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
		key := string(r)
		resultMap[key] = append(resultMap[key], s)
	}

	result := [][]string{}
	for _, v := range resultMap {
		result = append(result, v)
	}
	return result
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

// output will be map[eat:[eat tea ate] tan:[tan nat] bat:[bat]]
