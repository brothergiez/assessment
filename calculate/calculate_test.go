package calculate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTotal(t *testing.T) {
	tests := []struct {
		name     string
		items    []map[string]interface{}
		expected int
		err      string
	}{
		{
			name: "Valid input with string price and quantity",
			items: []map[string]interface{}{
				{"price": "10", "quantity": "2"},
				{"price": "20", "quantity": "3"},
			},
			expected: 80,
			err:      "",
		},
		{
			name: "Valid input with integer price and quantity",
			items: []map[string]interface{}{
				{"price": 15, "quantity": 4},
				{"price": 5, "quantity": 6},
			},
			expected: 90,
			err:      "",
		},
		{
			name: "Mixed valid input types",
			items: []map[string]interface{}{
				{"price": "10", "quantity": 2},
				{"price": 20, "quantity": "3"},
			},
			expected: 80,
			err:      "",
		},
		{
			name: "Missing price key",
			items: []map[string]interface{}{
				{"quantity": 2},
			},
			expected: 0,
			err:      "price not found in item",
		},
		{
			name: "Invalid price type",
			items: []map[string]interface{}{
				{"price": 10.5, "quantity": 2},
			},
			expected: 0,
			err:      "invalid type for price",
		},
		{
			name: "Invalid quantity value",
			items: []map[string]interface{}{
				{"price": "10", "quantity": "invalid"},
			},
			expected: 0,
			err:      "error converting quantity to integer",
		},
		{
			name:     "Empty input",
			items:    []map[string]interface{}{},
			expected: 0,
			err:      "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			total, err := CalculateTotal(test.items)

			if test.err != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), test.err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, test.expected, total)
		})
	}
}
