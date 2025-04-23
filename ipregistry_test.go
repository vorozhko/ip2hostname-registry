package main

import (
	"testing"
)


func TestResolve(t *testing.T) {

	tests := map[string]struct {
		ip       string
		expected []string
	}{
		"IP in registry": {
			ip:       "8.8.8.8",
			expected: []string{"dns.google."},
		},
		"IP not in registry but resolved": {
			ip:       "1.1.1.1",
			expected: []string{"one.one.one.one."},
		},
		"IP not in registry and not resolved": {
			ip:       "52.94.236.248",
			expected: []string{},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result := search(tt.ip)
			if len(result) != len(tt.expected) || (len(result) > 0 && result[0] != tt.expected[0]) {
				t.Errorf("resolve(%s) = %v; want %v", tt.ip, result, tt.expected)
			}
		})
	}
}
