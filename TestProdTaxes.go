package checkout

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckout(t *testing.T) {
	testCases := []struct {
		name     string
		state    string
		items    []Item
		expected float64
		err      error
	}{
		{
			name:  "DE: empty cart",
			state: "DE",
			items: []Item{},
			expected: 0,
			err:      nil,
		},
		{
			name: "NJ: clothing with fur",
			state: "NJ",
			items: []Item{
				{Name: "Furry Coat", Price: 50, Type: "Clothing"},
			},
			expected: 53.3,
			err:      nil,
		},
		{
			name: "PA: wic eligible food",
			state: "PA",
			items: []Item{
				{Name: "Milk", Price: 2.99, Type: "Wic Eligible food"},
				{Name: "Bread", Price: 3.99, Type: "Wic Eligible food"},
			},
			expected: 6.98,
			err:      nil,
		},
		{
			name: "Unknown state",
			state: "NY",
			items: []Item{
				{Name: "Hat", Price: 19.99, Type: "Clothing"},
				{Name: "Sunglasses", Price: 24.99, Type: "Accessories"},
			},
			expected: 0,
			err:      fmt.Errorf("state NY not supported"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Checkout(tc.state, tc.items)
			assert.Equal(t, tc.expected, actual)
			assert.Equal(t, tc.err, err)
		})
	}
}
