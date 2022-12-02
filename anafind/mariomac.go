package anafind

import (
	"strings"

	"github.com/mariomac/gostream/item"
	"github.com/mariomac/gostream/stream"
)

func Mariomac(words []string) map[string]map[string]struct{} {
	wordStream := stream.OfSlice(words).Filter(func(w string) bool {
		return len(w) > 1
	}).Map(strings.ToLower)
	groupedWords, _ := stream.Map(wordStream, func(w string) Anagrams {
		a := Anagrams{}
		a.AddWord(w)
		return a
	}).Reduce(Accumulate)

	actualAnagrams := stream.OfMap(groupedWords).
		Filter(func(entry item.Pair[string, map[string]struct{}]) bool {
			return len(entry.Val) > 1
		})
	return stream.ToMap(actualAnagrams)
}
