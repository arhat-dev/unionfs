package unionfs

import (
	"io/fs"
	"testing"
	"testing/fstest"

	"github.com/stretchr/testify/assert"
)

func TestUnionFS_Open(t *testing.T) {
	tests := []struct {
		name string
		path string

		expectedName string
		expectDir    bool
		expectErr    error
	}{
		{
			name:         "File Exists Directly",
			path:         "foo",
			expectedName: "foo",
		},
		{
			name:      "File Not Exists",
			path:      "bar",
			expectErr: fs.ErrNotExist,
		},
		{
			name:         "Dir Exists In Underlay FS",
			path:         "foo/dir",
			expectedName: "foo/dir",
			expectDir:    true,
		},
		{
			name:         "File Exists In Underlay FS",
			path:         "foo/dir/foo",
			expectedName: "foo/dir/foo",
			expectDir:    false,
		},
		{
			name:         "Path Is Dir",
			path:         "some",
			expectedName: "some",
			expectDir:    true,
		},
	}

	rmfs := New()
	rmfs.Map("foo", "bar", fstest.MapFS{
		"bar": &fstest.MapFile{
			Data: []byte("bar"),
		},
		"bar/dir/foo": &fstest.MapFile{
			Data: []byte("foo"),
		},
	})

	rmfs.Map("some/data", "some/data", fstest.MapFS{
		"data": &fstest.MapFile{
			Data: []byte("data"),
		},
	})

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f, err := rmfs.Open(test.path)
			if test.expectErr != nil {
				assert.ErrorIs(t, err, test.expectErr)
				return
			}

			if !assert.NoError(t, err) {
				return
			}

			fInfo, err := f.Stat()
			assert.NoError(t, err)
			assert.Equal(t, test.expectedName, fInfo.Name())
			assert.Equal(t, test.expectDir, fInfo.IsDir())
		})
	}
}
