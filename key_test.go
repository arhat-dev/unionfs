package unionfs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathKeyFunc(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected []string
	}{
		{
			name:     "Empty Path",
			path:     "",
			expected: nil,
		},
		{
			name:     "Root Path",
			path:     "/",
			expected: []string{"/"},
		},
		{
			name:     "Absolute Path",
			path:     "/a/b/c/d/e/f",
			expected: []string{"/", "/a", "/b", "/c", "/d", "/e", "/f"},
		},
		{
			name:     "Relative Path Single Element",
			path:     "test",
			expected: []string{"test"},
		},
		{
			name:     "Relative Path Multiple Elements",
			path:     "test/foo/foo",
			expected: []string{"test", "/foo", "/foo"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var keys []string
			var (
				k    string
				next int
			)

			for k, next = PathKeyFunc(test.path, 0); next > 0; k, next = PathKeyFunc(test.path, next) {
				keys = append(keys, k)
			}

			assert.EqualValues(t, test.expected, keys)
			assert.Equal(t, -1, next)
		})
	}
}
