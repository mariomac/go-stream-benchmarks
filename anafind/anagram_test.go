package anafind

import "testing"
import "github.com/stretchr/testify/assert"

func TestAnagrams_Accumulate(t *testing.T) {
	a, b := Anagrams{}, Anagrams{}
	a.AddWord("hello")
	a.AddWord("olleh")
	a.AddWord("friend")
	b.AddWord("rifend")
	b.AddWord("mario")
	a = Accumulate(a, b)
	assert.Equal(t, Anagrams{
		"ehllo":  map[string]struct{}{"hello": {}, "olleh": {}},
		"definr": map[string]struct{}{"friend": {}, "rifend": {}},
		"aimor":  map[string]struct{}{"mario": {}},
	}, a)
}
