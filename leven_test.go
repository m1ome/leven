package leven

import (
	"fmt"
	agext "github.com/agext/levenshtein"
	agnivade "github.com/agnivade/levenshtein"
	arbovm "github.com/arbovm/levenshtein"
	"testing"
)

var testSource = []struct {
	first, second string
	distance      int
}{
	{"a", "a", 0},
	{"ab", "ab", 0},
	{"ab", "aa", 1},
	{"ab", "aa", 1},
	{"ab", "aaa", 2},
	{"bbb", "a", 3},
	{"kitten", "sitting", 3},
	{"aa", "aü", 1},
	{"Fön", "Föm", 1},
}

var benchmarkWords = map[string]struct {
	first, second string
}{
	"Default":   {"kitten", "sitting"},
	"Prefixed":  {"prefixedkitten", "prefixedsitting"},
	"Postfixed": {"kittenpostfixed", "sittingpostfixed"},
	"Fullscan":  {"kitten", "loyyrm"},
}

func TestDistance(t *testing.T) {
	for i, test := range testSource {
		res := Distance(test.first, test.second)
		if res != test.distance {
			o := fmt.Sprintf("%v \t distance between %v and %v should be %v but returned %v.",
				i, test.first, test.second, test.distance, res)
			t.Errorf(o)
		}
	}
}

var result int

func benckmarkDistance(first string, second string, distance func(s1, s2 string) int, b *testing.B) {
	var r int

	for n := 0; n < b.N; n++ {
		r = distance(first, second)
	}

	result = r
}

// Benchmarking all stuff together
//
// Default distance
//
func BenchmarkDefaultM1omeDistance(b *testing.B) {
	test := benchmarkWords["Default"]
	benckmarkDistance(test.first, test.second, func(s1, s2 string) int {
		return Distance(s1, s2)
	}, b)
}

func BenchmarkDefaultAgextLevenshteinDistance(b *testing.B) {
	test := benchmarkWords["Default"]
	benckmarkDistance(test.first, test.second, func(s1, s2 string) int {
		return agext.Distance(s1, s2, nil)
	}, b)
}

func BenchmarkDefaultAgnivadeLevenshteinDistance(b *testing.B) {
	test := benchmarkWords["Default"]
	benckmarkDistance(test.first, test.second, func(s1, s2 string) int {
		return agnivade.ComputeDistance(s1, s2)
	}, b)
}

func BenchmarkDefaultArbovmLevenshteinDistance(b *testing.B) {
	test := benchmarkWords["Default"]
	benckmarkDistance(test.first, test.second, func(s1, s2 string) int {
		return arbovm.Distance(s1, s2)
	}, b)
}

//
// Prefixed
//
func BenchmarkPrefixedM1omeDistance(b *testing.B) {
	test := benchmarkWords["Prefixed"]
	benckmarkDistance(test.first, test.second, func(s1, s2 string) int {
		return Distance(s1, s2)
	}, b)
}

func BenchmarkPrefixedAgextLevenshteinDistance(b *testing.B) {
	test := benchmarkWords["Prefixed"]
	benckmarkDistance(test.first, test.second, func(s1, s2 string) int {
		return agext.Distance(s1, s2, nil)
	}, b)
}

func BenchmarkPrefixedAgnivadeLevenshteinDistance(b *testing.B) {
	test := benchmarkWords["Prefixed"]
	benckmarkDistance(test.first, test.second, func(s1, s2 string) int {
		return agnivade.ComputeDistance(s1, s2)
	}, b)
}

func BenchmarkPrefixedArbovmLevenshteinDistance(b *testing.B) {
	test := benchmarkWords["Prefixed"]
	benckmarkDistance(test.first, test.second, func(s1, s2 string) int {
		return arbovm.Distance(s1, s2)
	}, b)
}

//
// Postfixed
//

func BenchmarkPostfixedM1omeDistance(b *testing.B) {
	test := benchmarkWords["Postfixed"]
	benckmarkDistance(test.first, test.second, func(s1, s2 string) int {
		return Distance(s1, s2)
	}, b)
}

func BenchmarkPostfixedAgextLevenshteinDistance(b *testing.B) {
	test := benchmarkWords["Postfixed"]
	benckmarkDistance(test.first, test.second, func(s1, s2 string) int {
		return agext.Distance(s1, s2, nil)
	}, b)
}

func BenchmarkPostfixedAgnivadeLevenshteinDistance(b *testing.B) {
	test := benchmarkWords["Postfixed"]
	benckmarkDistance(test.first, test.second, func(s1, s2 string) int {
		return agnivade.ComputeDistance(s1, s2)
	}, b)
}

func BenchmarkPostfixedArbovmLevenshteinDistance(b *testing.B) {
	test := benchmarkWords["Postfixed"]
	benckmarkDistance(test.first, test.second, func(s1, s2 string) int {
		return arbovm.Distance(s1, s2)
	}, b)
}

//
// Fullscan
//

func BenchmarkFullscanM1omeDistance(b *testing.B) {
	test := benchmarkWords["Fullscan"]
	benckmarkDistance(test.first, test.second, func(s1, s2 string) int {
		return Distance(s1, s2)
	}, b)
}

func BenchmarkFullscanAgextLevenshteinDistance(b *testing.B) {
	test := benchmarkWords["Fullscan"]
	benckmarkDistance(test.first, test.second, func(s1, s2 string) int {
		return agext.Distance(s1, s2, nil)
	}, b)
}

func BenchmarkFullscanAgnivadeLevenshteinDistance(b *testing.B) {
	test := benchmarkWords["Fullscan"]
	benckmarkDistance(test.first, test.second, func(s1, s2 string) int {
		return agnivade.ComputeDistance(s1, s2)
	}, b)
}

func BenchmarkFullscanArbovmLevenshteinDistance(b *testing.B) {
	test := benchmarkWords["Fullscan"]
	benckmarkDistance(test.first, test.second, func(s1, s2 string) int {
		return arbovm.Distance(s1, s2)
	}, b)
}
