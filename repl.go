package main

import (
	"bufio"
	"fmt"
	"os"
)

const defaultPrompt = "> "

type Repl struct {
	reader     *bufio.Reader
	prompt     string
	line       string
	lastResult int64

	Trace bool
}

func NewRepl() *Repl {
	return &Repl{
		reader: bufio.NewReader(os.Stdin),
		prompt: defaultPrompt,
	}
}

func (r *Repl) ReadLine() (err error) {
	var l string

	if l, err = r.reader.ReadString('\n'); err == nil {
		r.line = l[:len(l)-1]
	}

	return
}

func (r *Repl) PrintPrompt() {
	fmt.Print(r.prompt)
}

func (r *Repl) Eval() (err error) {
	parser := &HeckParser{Buffer: r.line}
	parser.Init()

	if err = parser.Parse(); err == nil {

		if r.Trace {
			parser.PrintSyntaxTree()
		}

		parser.Execute()
		r.lastResult = parser.Value()
	}

	return
}

func (r *Repl) PrintLastResult() {
	hex := fmt.Sprintf("%X", r.lastResult)
	dec := fmt.Sprintf("%d", r.lastResult)

	if hex == dec {
		fmt.Println(hex)
	} else {
		fmt.Printf("%s \t(%s)\n", hex, dec)
	}
}

func (r *Repl) Loop() (err error) {
	for {
		r.PrintPrompt()

		if err = r.ReadLine(); err != nil {
			return
		}

		if err = r.Eval(); err != nil {
			return
		}

		r.PrintLastResult()
	}

	return
}
