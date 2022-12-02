package anafind

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkMariomac(b *testing.B) {
	input := TextAsSlice()
	var out map[string]map[string]struct{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		out = Mariomac(input)
	}
	b.StopTimer()
	checkCorrectness(b, out)
}

func checkCorrectness(b *testing.B, in map[string]map[string]struct{}) {
	assert.Equal(b, map[string]map[string]struct{}{
		"eehinrt": {"neither": {}, "therein": {}},
		"eehrt":   {"there": {}, "three": {}},
		"no":      {"no": {}, "on": {}},
	}, in)
}
func debugPrint(in interface{}) {
	txt, _ := json.MarshalIndent(in, "", "    ")
	fmt.Println(string(txt))
}
