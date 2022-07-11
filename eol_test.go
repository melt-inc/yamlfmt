package yamlfmt

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	SimpleCRLFText = "this\r\nis some\r\ncrlf\r\ntext\r\n"
	SimpleCRText   = "this is text with a random \r thrown in"
)

func TestTransformer(t *testing.T) {
	t.Run("SimpleCRLFText", func(t *testing.T) {
		reader := NewReader(strings.NewReader(SimpleCRLFText))
		b, err := io.ReadAll(reader)
		require.NoError(t, err)
		assert.Equal(t, "this\nis some\ncrlf\ntext\n", string(b))
	})
	t.Run("SimpleCRText", func(t *testing.T) {
		reader := NewReader(strings.NewReader(SimpleCRText))
		b, err := io.ReadAll(reader)
		require.NoError(t, err)
		assert.Equal(t, SimpleCRText, string(b))
	})
}
