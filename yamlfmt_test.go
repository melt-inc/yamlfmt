package yamlfmt

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSfmt(t *testing.T) {
	s, err := Sfmt(`
test: 'this is a test'
and:  'so is this'
`)
	require.NoError(t, err)
	assert.Equal(t, "test: 'this is a test'\nand: 'so is this'\n", s)
}

func TestFfmt(t *testing.T) {
	reader := strings.NewReader(`
test: 'this is a test'
and:  'so is this'
`)
	writer := new(bytes.Buffer)
	err := Ffmt(reader, writer)
	require.NoError(t, err)
	assert.Equal(t, "test: 'this is a test'\nand: 'so is this'\n", writer.String())
}
