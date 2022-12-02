package anafind

import (
	"strings"

	"github.com/vladimirvivien/automi/collectors"
	"github.com/vladimirvivien/automi/stream"
)

func Automi(words []string) map[string]map[string]struct{} {
	var res Anagrams
	str := stream.New(words).
		Filter(MoreThanOneChar).
		Map(strings.ToLower).
		Map(SingleWordToMap).
		Reduce(Anagrams{}, Accumulate).
		Into(collectors.Func(func(i interface{}) error {
			res = i.(Anagrams)
			return nil
		}))
	if err := <-str.Open(); err != nil {
		panic(err)
	}
	return res
}
