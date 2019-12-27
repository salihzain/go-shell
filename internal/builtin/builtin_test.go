package builtin

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestChangeDir(t *testing.T) {
	before, err := os.Getwd()
	require.NoError(t, err)

	err = os.Mkdir("test", os.ModePerm)
	require.NoError(t, err)

	err = changeDir([]string{"", "test"})
	require.NoError(t, err)

	after, err := os.Getwd()
	require.NoError(t, err)
	require.Equal(t, before+"/test", after)

	err = changeDir([]string{"", "../"})
	require.NoError(t, err)

	err = os.Remove("test")
	require.NoError(t, err)
}
