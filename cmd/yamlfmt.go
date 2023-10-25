package main

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/jamesrom/yamlfmt"
	flag "github.com/spf13/pflag"
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

	// if no flags given, read from stdin
	if len(flag.Args()) < 1 {
		err := yamlfmt.Ffmt(os.Stdin, os.Stdout, yamlfmt.WithCompactSequenceStyle(*compact), yamlfmt.WithIndentSize(*tabsize))
		if err != nil {
			panic(err)
		}
	}

	for _, filePath := range flag.Args() {
		in, err := os.Open(filePath)
		if err != nil {
			in.Close()
			continue
		}
		defer in.Close()

		// use an in-memory buffer for the output
		var buf bytes.Buffer
		var writer io.Writer
		if *stdout {
			writer = io.MultiWriter(&buf, os.Stdout)
		} else if *inplace {
			writer = &buf
		} else {
			writer = os.Stdout
		}

		// do it
		err = yamlfmt.Ffmt(
			in,
			writer,
			yamlfmt.WithCompactSequenceStyle(*compact),
			yamlfmt.WithIndentSize(*tabsize),
		)
		in.Close()
		if err != nil {
			panic(err)
		}

		if *inplace {
			out, err := os.OpenFile(filePath, os.O_WRONLY, os.ModePerm)
			defer out.Close()
			if err != nil {
				panic(err)
			}
			n, err := io.Copy(out, &buf)
			if err != nil {
				panic(err)
			}
			out.Truncate(n)
			if err != nil {
				panic(err)
			}
			out.Close()
		}
	}
}
