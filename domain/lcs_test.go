package domain

import (
	"testing"
)

func TestLcs(t *testing.T) {
	testCases := []struct {
		stringA  string
		stringB  string
		expected int
	}{
		{"AAABBAAA", "AAABBBAA", 7},
		{"XAYZABC", "ABCXYZ", 3},
		{"AGGTAB", "GXTXAYB", 4},
		{"TRYSTAAAAAAN", "TRYSTAN", 7},
	}

	for _, tc := range testCases {
		result := Lcs(tc.stringA, tc.stringB, len(tc.stringA), len(tc.stringB))
		if result != tc.expected {
			t.Errorf("Lcs(%s, %s): expected %d, result %d", tc.stringA, tc.stringB, tc.expected, result)
		}
	}
}

func TestLcsDP(t *testing.T) {
	testCases := []struct {
		stringA  string
		stringB  string
		expected int
	}{
		{"AAABBAAA", "AAABBBAA", 7},
		{"XAYZABC", "ABCXYZ", 3},
		{"AGGTAB", "GXTXAYB", 4},
		{"TRYSTAAAAAAN", "TRYSTAN", 7},
		{"This is a test", "This is a another test", 14},
		{"It was the best of times, it was the worst of times.", "It was the age of wisdom, it was the age of foolishness.", 39},
	}

	for _, tc := range testCases {
		result := LcsDP(tc.stringA, tc.stringB)
		if result != tc.expected {
			t.Errorf("LcsDP(%s, %s): expected %d, result %d", tc.stringA, tc.stringB, tc.expected, result)
		}
	}
}

func BenchmarkLcs(b *testing.B) {
	stringA := "It was the best of times, it was the worst of times. Then it was the age of wisdom, it was the age of foolishness. It was the epoch of belief, it was the epoch of incredulity. It was the season of Light, it was the season of Darkness. It was the spring of hope, it was the winter of despair."
	stringB := "It was the best of times, it was the worst of times. Then it was the age of wisdom, it was the age of foolishness. It was the epoch of belief, it was the epoch of incredulity. Put here another sentence.  It was the spring of hope, it was the winter of despair."
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Lcs(stringA, stringB, len(stringA), len(stringB))
	}
}

func BenchmarkLcsDP(b *testing.B) {
	stringA := "It was the best of times, it was the worst of times. Then it was the age of wisdom, it was the age of foolishness. It was the epoch of belief, it was the epoch of incredulity. It was the season of Light, it was the season of Darkness. It was the spring of hope, it was the winter of despair."
	stringB := "It was the best of times, it was the worst of times. Then it was the age of wisdom, it was the age of foolishness. It was the epoch of belief, it was the epoch of incredulity. Put here another sentence.  It was the spring of hope, it was the winter of despair."
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LcsDP(stringA, stringB)
	}
}
