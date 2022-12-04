package anafind

import (
	"strings"

	"github.com/kamstrup/fn"
)

func Fn(words []string) map[string]map[string]struct{} {
	wordsLower := fn.ArrayOf(words).
		Where(MoreThanOneChar).
		Shape(strings.ToLower)
	singleWords := fn.MapOf(wordsLower, SingleWordToMap)
	return fn.Into(Anagrams{}, Accumulate, singleWords)
}
