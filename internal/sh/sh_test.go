package sh

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSplitLine(t *testing.T) {
	line := "echo \"hello world\""
	args, argsOK := splitLine(line)
	require.Equal(t, true, argsOK)
	require.Equal(t, 2, len(args))
	require.Equal(t, "echo", args[0])
	require.Equal(t, "\"hello world\"", args[1])

	line = "gcc -o test test.c"
	args, argsOK = splitLine(line)
	require.Equal(t, true, argsOK)
	require.Equal(t, 4, len(args))
	require.Equal(t, "gcc", args[0])
	require.Equal(t, "-o", args[1])
	require.Equal(t, "test", args[2])
	require.Equal(t, "test.c", args[3])

	line = ""
	args, argsOK = splitLine(line)
	require.Equal(t, false, argsOK)
	require.Equal(t, 0, len(args))
}

func TestExecute(t *testing.T) {
	err := execute([]string{})
	require.Error(t, err)

	err = execute([]string{"ls", "-a"})
	require.NoError(t, err)

	// command doesn't exist
	err = execute([]string{"clean", "dishes"})
	require.Error(t, err)
}
