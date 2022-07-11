package yamlfmt

import (
	"io"

	"golang.org/x/text/transform"
)

type CRLFTransformer struct {
	lastCharacterWasCR bool
}

func (t *CRLFTransformer) Transform(dst, src []byte, atEOF bool) (j int, i int, err error) {
	// If the previous character from the last batch was a CR and the first
	// character this batch is a LF, then the CRLF spanned the boundary.
	// However, since we didn't emit a CR in the previous batch,
	// if the first character is _not_ a LF, we need to emit the CR now.
	if t.lastCharacterWasCR && src[i] != '\n' {
		dst[j] = '\r'
		j++
	}

	// Check if we can exit early if this was the last call
	if atEOF && len(src) == 0 {
		return j, i, nil
	}

	// for all but last character
	for i < len(src)-1 && j < len(dst) {
		// check if it's a carriage return
		if src[i] == '\r' {
			// look ahead to see if it's followed by linefeed
			if src[i+1] == '\n' {
				// if so, jump over the carriage return
				i++
				continue
			}
		}

		dst[j] = src[i]
		i++
		j++
	}

	// For the last character, if it's a carriage return, we can't peek ahead
	// so flag it for the next call. Note that we don't emit the CR as we can't
	// be sure at this stage that we should keep it.
	if !atEOF && src[i] == '\r' {
		t.lastCharacterWasCR = true
		i++
		return j, i, nil
	}

	t.lastCharacterWasCR = false
	if j < len(dst) {
		dst[j] = src[i]
		i++
		j++
		return j, i, nil
	}

	return j, i, transform.ErrShortDst
}

func (t *CRLFTransformer) Reset() {
	t.lastCharacterWasCR = false
}

// NewReader returns a reader that transforms all CRLF in r to LF
func NewReader(r io.Reader) io.Reader {
	return transform.NewReader(r, new(CRLFTransformer))
}
