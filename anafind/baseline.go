package anafind

import "strings"

func Baseline(words []string) map[string]map[string]struct{} {
	var swa []Anagrams
	for _, w := range words {
		if !MoreThanOneChar(w) {
			continue
		}
		swa = append(swa, SingleWordAnagrams(strings.ToLower(w)))
	}
	if len(swa) == 0 {
		return nil
	}
	seed := swa[0]
	for i := range swa[1:] {
		seed = Accumulate(seed, swa[i])
	}
	return seed
}
