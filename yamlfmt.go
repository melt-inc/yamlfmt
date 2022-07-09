package yamlfmt

import (
	"bytes"
	"io"
	"strings"

	yaml "sigs.k8s.io/yaml/goyaml.v3"
)

type decoder interface {
	Decode(v interface{}) (err error)
}
type encoder interface {
	Encode(v interface{}) (err error)
}

func Sfmt(in string, opt ...Option) (string, error) {
	decoder := yaml.NewDecoder(strings.NewReader(in))
	writer := new(bytes.Buffer)
	encoder := NewEncoder(writer, opt...)
	defer encoder.Close()
	err := Transcode(decoder, encoder)
	return writer.String(), err
}

func Ffmt(in io.Reader, out io.Writer, opt ...Option) error {
	decoder := yaml.NewDecoder(in)
	encoder := NewEncoder(out, opt...)
	defer encoder.Close()
	return Transcode(decoder, encoder)
}

func Transcode(in decoder, out encoder) error {
	var n yaml.Node
	if err := in.Decode(&n); err != nil {
		return err
	}
	if err := out.Encode(&n); err != nil {
		return err
	}
	return nil
}
