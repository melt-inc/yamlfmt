package main

import (
	"fmt"
	"io"
	"os"

	"github.com/jamesrom/yamlfmt"
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
	encoder := yamlfmt.NewEncoder(
		os.Stdout,
		yamlfmt.WithCompactSequenceFormat(*compact),
		yamlfmt.WithIndentSize(*tabsize),
	)
	defer encoder.Close()

	for err == nil {
		err = yamlfmt.Transcode(decoder, encoder)
	}
	if err != io.EOF {
		panic(err)
	}
}
