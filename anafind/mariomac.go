package anafind

import (
	"strings"

	"github.com/mariomac/gostream/stream"
)

func Mariomac(words []string) map[string]map[string]struct{} {
	wordStream := stream.OfSlice(words).
		Filter(MoreThanOneChar).
		Map(strings.ToLower)
	groupedWords, _ := stream.Map(wordStream, SingleWordAnagrams).
		Reduce(Accumulate)
	return groupedWords
}
