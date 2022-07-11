package yamlfmt

import (
	"io"

	yaml "sigs.k8s.io/yaml/goyaml.v3"
)

type Option func(*yaml.Encoder)

func WithIndentSize(size int) Option {
	return func(e *yaml.Encoder) {
		e.SetIndent(size)
	}
}

func WithCompactSequenceStyle(compact bool) Option {
	return func(e *yaml.Encoder) {
		if compact {
			e.CompactSeqIndent()
		} else {
			e.DefaultSeqIndent()
		}
	}
}

func NewEncoder(w io.Writer, opts ...Option) *yaml.Encoder {
	encoder := yaml.NewEncoder(w)
	for _, opt := range opts {
		opt(encoder)
	}
	return encoder
}
