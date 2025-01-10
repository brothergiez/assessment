package main

import (
	"os/exec"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func sortAnagramGroups(groups [][]string) [][]string {
	for _, group := range groups {
		sort.Strings(group)
	}
	sort.Slice(groups, func(i, j int) bool {
		return groups[i][0] < groups[j][0]
	})
	return groups
}

func parseAnagramOutput(output string) [][]string {
	output = strings.TrimSpace(output)
	output = strings.TrimPrefix(output, "Grouped Anagrams: ")
	output = strings.Trim(output, "[]")
	groups := strings.Split(output, "] [")
	var result [][]string
	for _, group := range groups {
		group = strings.Trim(group, "[]")
		result = append(result, strings.Split(group, " "))
	}
	return result
}

func TestMainAnagram(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go", "anagram", "--words=listen,silent,enlist,google,gogole,cat,act")
	output, err := cmd.CombinedOutput()

	assert.NoError(t, err)

	expected := [][]string{
		{"listen", "silent", "enlist"},
		{"google", "gogole"},
		{"cat", "act"},
	}

	actual := parseAnagramOutput(string(output))
	expected = sortAnagramGroups(expected)
	actual = sortAnagramGroups(actual)

	assert.Equal(t, expected, actual, "Anagram groups do not match.")
}
func TestMainBooking(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go", "booking")
	output, err := cmd.CombinedOutput()

	expectedSubstrings := []string{
		"Room 'room1' with capacity 5 added successfully.",
		"Memesan room1 mulai jam 10:00 selama 2 jam",
		"Error: booking conflict detected",
	}

	assert.NoError(t, err)
	for _, substr := range expectedSubstrings {
		assert.Contains(t, string(output), substr)
	}
}

func TestMainCalculate(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go", "calculate")
	output, err := cmd.CombinedOutput()

	expectedOutput := "Total Price: 35\n"
	assert.NoError(t, err)
	assert.Equal(t, expectedOutput, string(output))
}

func TestMainInvalidCommand(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go", "invalid")
	output, err := cmd.CombinedOutput()

	expectedOutput := "Unknown command. Use 'anagram' or 'booking'.\n"
	assert.NoError(t, err)
	assert.Equal(t, expectedOutput, string(output))
}
