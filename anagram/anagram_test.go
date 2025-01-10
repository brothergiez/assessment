package anagram

import (
	"reflect"
	"testing"
)

func TestGroupAnagram(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		expected [][]string
	}{
		{
			name:  "Basic anagram grouping",
			words: []string{"listen", "silent", "enlist", "google", "gogole", "cat", "act"},
			expected: [][]string{
				{"listen", "silent", "enlist"},
				{"google", "gogole"},
				{"cat", "act"},
			},
		},
		{
			name:     "Empty input",
			words:    []string{},
			expected: [][]string{},
		},
		{
			name:  "No anagrams",
			words: []string{"hello", "world", "golang"},
			expected: [][]string{
				{"hello"},
				{"world"},
				{"golang"},
			},
		},
		{
			name:  "All identical words",
			words: []string{"test", "test", "test"},
			expected: [][]string{
				{"test", "test", "test"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := GroupAnagram(test.words)

			resultMap := map[string][]string{}
			for _, group := range result {
				key := group[0]
				resultMap[key] = group
			}

			expectedMap := map[string][]string{}
			for _, group := range test.expected {
				key := group[0]
				expectedMap[key] = group
			}

			if !reflect.DeepEqual(resultMap, expectedMap) {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}
