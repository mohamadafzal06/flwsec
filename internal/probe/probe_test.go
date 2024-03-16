package probe

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	_, err := load()
	//t.Log(objs.probePrograms.Probe.VerifierLog)
	require.NoError(t, err)
}

func TestPakcet(t *testing.T) {
	objs, err := load()
	require.NoError(t, err)

	in := []byte{1, 2, 3, 4, 5, 6, 7, 8, 0, 9, 1, 2, 3, 4, 5}
	// retCode should be 0, because in prog.c we return 0.
	retCode, out, err := objs.probePrograms.Probe.Test(in)
	require.NoError(t, err)
	require.Equal(t, 0, retCode)
	require.Equal(t, in, out)

}
