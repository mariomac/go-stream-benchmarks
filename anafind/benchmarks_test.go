package anafind

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkBaseline(b *testing.B) {
	input := TextAsSlice()
	var out map[string]map[string]struct{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		out = Baseline(input)
	}
	b.StopTimer()
	checkCorrectness(b, out)
}

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

func BenchmarkLambda(b *testing.B) {
	input := TextAsSlice()
	var out map[string]map[string]struct{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		out = Lambda(input)
	}
	b.StopTimer()
	checkCorrectness(b, out)
}

func BenchmarkGoio(b *testing.B) {
	input := TextAsSlice()
	var out map[string]map[string]struct{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		out = Goio(input)
	}
	b.StopTimer()
	checkCorrectness(b, out)
}

func BenchmarkAutomi(b *testing.B) {
	input := TextAsSlice()
	var out map[string]map[string]struct{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		out = Automi(input)
	}
	b.StopTimer()
	checkCorrectness(b, out)
}

func checkCorrectness(b *testing.B, in map[string]map[string]struct{}) {
	assert.Equal(b, map[string]struct{}{"neither": {}, "therein": {}}, in["eehinrt"])
	assert.Equal(b, map[string]struct{}{"there": {}, "three": {}}, in["eehrt"])
	assert.Equal(b, map[string]struct{}{"no": {}, "on": {}}, in["no"])
}

func debugPrint(in interface{}) {
	txt, _ := json.MarshalIndent(in, "", "    ")
	fmt.Println(string(txt))
}
