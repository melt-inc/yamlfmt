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

// Sfmt formats the input yaml string and returns it
func Sfmt(in string, opt ...Option) (string, error) {
	decoder := yaml.NewDecoder(NewReader(strings.NewReader(in)))
	writer := new(bytes.Buffer)
	encoder := NewEncoder(writer, opt...)
	defer encoder.Close()
	err := Transcode(decoder, encoder)
	return writer.String(), err
}

// Ffmt reads unformatted yaml from in and writes formatted yaml to out
func Ffmt(in io.Reader, out io.Writer, opt ...Option) error {
	decoder := yaml.NewDecoder(NewReader(in))
	encoder := NewEncoder(out, opt...)
	defer encoder.Close()
	return Transcode(decoder, encoder)
}

// Transcode decodes and encodes every yaml document from in to out
func Transcode(in decoder, out encoder) (err error) {
	for err == nil {
		err = transcodeNextNode(in, out)
	}
	if err == io.EOF {
		return nil
	}
	return err
}

func transcodeNextNode(in decoder, out encoder) error {
	var n yaml.Node
	if err := in.Decode(&n); err != nil {
		return err
	}
	if err := out.Encode(&n); err != nil {
		return err
	}
	return nil
}
