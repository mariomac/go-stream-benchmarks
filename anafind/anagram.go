package anafind

import (
	"regexp"
	"sort"
)

// Anagrams Key: sorted lowercase characters of a given word
// Value: set of lowercase words matching the key as anagram
type Anagrams map[string]map[string]struct{}

// AddWord does not differentiate between upper and lowercase letters
func (a *Anagrams) AddWord(word string) {
	wb := []byte(word)
	sort.Slice(wb, func(i, j int) bool {
		return wb[i] < wb[j]
	})
	key := string(wb)
	if matches, ok := (*a)[key]; ok {
		matches[word] = struct{}{}
	} else {
		(*a)[key] = map[string]struct{}{word: {}}
	}
}

// Accumulate merges src into dst and returns dst
func Accumulate(dst, src Anagrams) Anagrams {
	for sorted, matches := range src {
		dstMatches, ok := dst[sorted]
		if !ok {
			dstMatches = map[string]struct{}{}
			dst[sorted] = dstMatches
		}
		for word := range matches {
			dstMatches[word] = struct{}{}
		}
	}
	return dst
}

var text = `
On the other part are two rocks, whereof the one reaches with sharp peak to the
wide heaven, and a dark cloud encompasses it; this never streams away, and
there is no clear air about the peak neither in summer nor in harvest tide. No
mortal man may scale it or set foot thereon, not though he had twenty hands and
feet. For the rock is smooth, and sheer, as it were polished. And in the midst
of the cliff is a dim cave turned to Erebus, towards the place of darkness,
whereby ye shall even steer your hollow ship, noble Odysseus. Not with an arrow
from a bow might a man in his strength reach from his hollow ship into that
deep cave. And therein dwelleth Scylla, yelping terribly. Her voice indeed is
no greater than the voice of a new-born whelp, but a dreadful monster is she,
nor would any look on her gladly, not if it were a god that met her. Verily she
hath twelve feet all dangling down; and six necks exceeding long, and on each a
hideous head, and therein three rows of teeth set thick and close, full of
black death. Up to her middle is she sunk far down in the hollow cave, but
forth she holds her heads from the dreadful gulf, and there she fishes,
swooping round the rock, for dolphins or sea-dogs, or whatso greater beast she
may anywhere take, whereof the deep-voiced Amphitrite feeds countless flocks.
Thereby no sailors boast that they have fled scatheless ever with their ship,
for with each head she carries off a man, whom she hath snatched from out the
dark-prowed ship.
`

func TextAsSlice() []string {
	splitter := regexp.MustCompile("[^a-zA-Z]")
	return splitter.Split(text, -1)
}

func TextAsChannel() <-chan string {
	slice := TextAsSlice()
	ret := make(chan string, len(slice))
	for _, v := range slice {
		ret <- v
	}
	return ret
}
