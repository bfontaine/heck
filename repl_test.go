package main

import (
	"github.com/franela/goblin"
	o "github.com/onsi/gomega"
	"testing"
)

func TestRepl(t *testing.T) {
	g := goblin.Goblin(t)

	o.RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("NewRepl", func() {
		g.It("Should not return nil", func() {
			o.Expect(NewRepl()).NotTo(o.BeNil())
		})

		g.It("Should set the default prompt", func() {
			o.Expect(NewRepl().prompt).To(o.Equal(defaultPrompt))
		})
	})

	g.Describe("Repl", func() {
		g.Describe("Eval()", func() {
			var repl *Repl

			g.BeforeEach(func() { repl = NewRepl() })

			g.It("Should return an error if it can't parse the line", func() {
				repl.line = "foob$ar-+@"
				o.Expect(repl.Eval()).NotTo(o.BeNil())
			})

			g.It("Should set .lastResult", func() {
				repl.line = "(2A+1)*-4C"
				o.Expect(repl.Eval()).To(o.BeNil())
				o.Expect(repl.lastResult).To(o.Equal(int64(-3268)))
			})
		})
	})
}
