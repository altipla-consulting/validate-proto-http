package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGoodFile(t *testing.T) {
	require.NoError(t, checkFiles("../../testdata/good"))
}

func TestBadNoInitialSlash(t *testing.T) {
	require.Error(t, checkFiles("../../testdata/bad-no-initial-slash"), "bad http mapping: bad/url")
}

func TestBadFinalSlash(t *testing.T) {
	require.Error(t, checkFiles("../../testdata/bad-final-slash"), "bad http mapping: /bad/url/")
}

func TestBadRepeated(t *testing.T) {
	require.Error(t, checkFiles("../../testdata/bad-repeated"), "repeated http mapping: /bad/url/")
}
