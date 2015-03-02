package main

import (
	"flag"
	"fmt"
	"io"
)

func main() {
	trace := flag.Bool("trace", false, "print the AST for each expression")
	flag.Parse()

	r := NewRepl()

	r.Trace = *trace

	if err := r.Loop(); err != nil && err != io.EOF {
		fmt.Printf("ERROR: %s\n", err)
	}
}
