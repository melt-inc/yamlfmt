package yamlfmt

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	Simple = `---
stack: "asd"
test: 'this is a test'
# this is a comment
and:  'so is this'`
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

func TestEOL(t *testing.T) {
	s, err := Sfmt("test: 'this is a test'\r\nand:  'so is this'")
	require.NoError(t, err)
	assert.Equal(t, "test: 'this is a test'\nand: 'so is this'\n", s)
}

func TestMultidocument(t *testing.T) {
	s, err := Sfmt(`
---
test: 'this is a test'
---
and:  'so is this'
`)
	require.NoError(t, err)
	assert.Equal(t, "test: 'this is a test'\n---\nand: 'so is this'\n", s)
}

func TestComment(t *testing.T) {
	s, err := Sfmt(Simple)
	require.NoError(t, err)
	assert.Equal(t, `stack: "asd"
test: 'this is a test'
# this is a comment
and: 'so is this'
`, s)

	reader := strings.NewReader(Simple)
	writer := new(bytes.Buffer)
	err = Ffmt(reader, writer, WithCompactSequenceStyle(false))
	require.NoError(t, err)
	assert.Equal(t, `stack: "asd"
test: 'this is a test'
# this is a comment
and: 'so is this'
`, writer.String())
}

func TestFileRead(t *testing.T) {
	reader, err := os.Open("examples/comment.yaml")
	require.NoError(t, err)
	writer := new(bytes.Buffer)
	err = Ffmt(reader, writer, WithCompactSequenceStyle(false))
	require.NoError(t, err)
	assert.Equal(t, `stack: "asd"
test: 'this is a test'
# this is a comment
and: 'so is this'
`, writer.String())
}
