package main

import (
	"fmt"
	"sort"
)

func topKFrequent(words []string, k int) []string {
	cnt := map[string]int{}
	for _, w := range words {
		cnt[w]++
	}
	uniqueWords := make([]string, 0, len(cnt))
	for w := range cnt {
		uniqueWords = append(uniqueWords, w)
	}
	sort.Slice(uniqueWords, func(i, j int) bool {
		s, t := uniqueWords[i], uniqueWords[j]
		return cnt[s] > cnt[t] || cnt[s] == cnt[t] && s < t
	})
	return uniqueWords[:k]
}

func main() {
	words := []string{"love", "love", "a", "cat", "a", "dog"}
	k := 2
	res := topKFrequent(words, k)
	fmt.Println(res)
}
