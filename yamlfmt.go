package main

import (
	"fmt"
	"io"
	"os"

	flag "github.com/spf13/pflag"
	yaml "sigs.k8s.io/yaml/goyaml.v3"
)

var (
	inplace *bool = flag.BoolP("in-place", "i", false, "Replace the file in-place (ignored when reading from stdin)")
	stdout  *bool = flag.BoolP("stdout", "o", false, "Write to stdout (default to stdout unless --in-place is enabled)")
	tabsize *int  = flag.IntP("tabsize", "t", 2, "Number of spaces for indentation")
	compact *bool = flag.BoolP("compact", "c", true, "Compact sequence style")
	help    *bool = flag.BoolP("help", "h", false, "")
)

func init() {
	// flag.Lookup("stdout").DefValue = "enabled unless --in-place is enabled"
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: yamlfmt [flags] [path ...]\n\n")
	fmt.Fprintf(os.Stderr, "If no paths given, read from stdin.\n\n")
	flag.CommandLine.SortFlags = false
	flag.PrintDefaults()
}

func main() {
	flag.Parse()

	if *help {
		usage()
		return
	}

	var err error
	decoder := yaml.NewDecoder(os.Stdin)
	encoder := yaml.NewEncoder(os.Stdout)
	encoder.SetIndent(*tabsize)
	encoder.CompactSeqIndent()
	defer encoder.Close()

	for err == nil {
		err = transcode(decoder, encoder)
	}
	if err != io.EOF {
		panic(err)
	}

	// type foo struct {
	// 	Bing string
	// 	Bang []string
	// 	Bong []int
	// }
	// x := foo{
	// 	Bing: "foo",
	// 	Bang: []string{"bar", "baz"},
	// 	Bong: []int{1, 2, 3},
	// }
	// out, err := yaml.Marshal(&x)
	// if err != nil {
	// 	panic(err)
	// }
	// os.Stdout.Write(out)
}

type decoder interface {
	Decode(v interface{}) (err error)
}
type encoder interface {
	Encode(v interface{}) (err error)
}

func transcode(in decoder, out encoder) error {
	var n yaml.Node
	if err := in.Decode(&n); err != nil {
		return err
	}
	if err := out.Encode(&n); err != nil {
		return err
	}
	return nil
}
