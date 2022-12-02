package anafind

import (
	"strings"

	"github.com/koss-null/lambda/pkg/pipe"
)

func Lambda(words []string) map[string]map[string]struct{} {
	wordStream := pipe.Slice(words).Parallel(1).
		Filter(MoreThanOneChar).
		Map(strings.ToLower)
	groupedWords := pipe.Map(wordStream, SingleWordToMap).
		Reduce(Accumulate)
	return *groupedWords
}
