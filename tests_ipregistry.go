package main

import (
	"testing"
)

// Mock the search function for testing
func mockSearch(name string) []string {
	return map[string][]string{
		"8.8.8.8":      {"dns.google"},
		"104.21.56.30": {"vorozhko.net"},
		"1.1.1.1":      {"one.one.one.one"},
	}[name]
}

func TestResolve(t *testing.T) {
	// Mock the search function
	originalSearch := search
	search = mockSearch
	defer func() { search = originalSearch }()

	tests := map[string]struct {
		ip       string
		expected []string
	}{
		"IP in registry": {
			ip:       "8.8.8.8",
			expected: []string{"dns.google"},
		},
		"IP not in registry but resolved": {
			ip:       "1.1.1.1",
			expected: []string{"one.one.one.one"},
		},
		"IP not in registry and not resolved": {
			ip:       "192.168.1.1",
			expected: []string{},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result := resolve(tt.ip)
			if len(result) != len(tt.expected) || (len(result) > 0 && result[0] != tt.expected[0]) {
				t.Errorf("resolve(%s) = %v; want %v", tt.ip, result, tt.expected)
			}
		})
	}
}
