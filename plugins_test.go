package protoeasy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExtraPlugin(t *testing.T) {
	testExtraPlugin(t, "foo", "--foo_out=")
	testExtraPlugin(t, "foo=hello=one", "--foo_out=hello=one:")
	testExtraPlugin(t, "foo=hello=one,woot=two", "--foo_out=hello=one,woot=two:")
}

func testExtraPlugin(t *testing.T, flagString string, flagMinusOutDirPath string) {
	plugin := newExtraPlugin(flagString)
	flags, err := plugin.Flags(nil, "", "/path/to/out")
	require.NoError(t, err)
	require.Len(t, flags, 1)
	require.Equal(t, fmt.Sprintf("%s/path/to/out", flagMinusOutDirPath), flags[0])
}
