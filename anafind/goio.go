package anafind

import (
	"strings"

	"github.com/primetalk/goio/io"
	"github.com/primetalk/goio/slice"
	"github.com/primetalk/goio/stream"
)

func Goio(words []string) map[string]map[string]struct{} {
	wordStr := stream.FromSlice(words)
	fWords := stream.Filter(wordStr, MoreThanOneChar)
	lfWords := stream.Map(fWords, strings.ToLower)
	swAgrs := stream.Map(lfWords, SingleWordAnagrams)
	swAgrsSl := stream.ToSlice(swAgrs)
	swArgsRs, err := io.ObtainResult(io.Continuation[[]Anagrams](swAgrsSl))
	if err != nil {
		panic(err)
	}
	return slice.Reduce(swArgsRs, Accumulate)
}
