package anafind

import (
	"strings"

	"github.com/kamstrup/fn"
)

func Fn(words []string) map[string]map[string]struct{} {
	nzWords := fn.ArrayOf(words).Where(MoreThanOneChar)
	lower := fn.MapOf(nzWords, strings.ToLower)
	singleWords := fn.MapOf(lower, SingleWordToMap)
	return fn.Into(Anagrams{}, Accumulate, singleWords)
}
