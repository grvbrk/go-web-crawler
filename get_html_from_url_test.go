package main

import (
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
	<html>
		<body>
			<a href="/path/one">
				<span>Boot.dev</span>
			</a>
			<a href="https://other.com/path/one">
				<span>Boot.dev</span>
			</a>
		</body>
	</html>
	`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
	<html>
		<body>
			<a href="/path/one">
				<span>Boot.dev</span>
			</a>
			<a href="https://other.com/path">
				<span>Boot.dev</span>
			</a>
		</body>
	</html>
	`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path"},
		}, {
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
	<html>
		<body>
			<a href="/path">
				<span>Boot.dev</span>
			</a>
			<a href="https://other.com/path/two">
				<span>Boot.dev</span>
			</a>
		</body>
	</html>
	`,
			expected: []string{"https://blog.boot.dev/path", "https://other.com/path/two"},
		},
	}

	for i, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(test.inputBody, test.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, test.name, err)
				return
			}
			ok := reflect.DeepEqual(actual, test.expected)
			if !ok {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, test.name, test.expected, actual)
			}
		})
	}
}
