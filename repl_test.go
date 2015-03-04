package main

import (
	"fmt"
	"github.com/franela/goblin"
	o "github.com/onsi/gomega"
	"testing"
)

func checkParsing(g *goblin.G, line string) {
	repl := NewRepl()
	repl.line = line
	// we don't use o.BeNil here because the parsing errors take too much place
	// when they're pretty-printed
	o.Expect(repl.Eval() == nil).To(o.BeTrue())
}

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

			// these were generated using:
			// $ abnfgen -y 10 heck.abnf

			lines := []string{
				"   d%          (   a   )   /       a",
				"   (   \tA-  A+  A   )   /   -   (   A   )",
				"\t\t\tdaA",
				"~          2aa-    ~\tA-  -   (   A   )   +   A",
			}

			for i, line := range lines {
				desc := fmt.Sprintf("Should parse valid expressions (%d)", i+1)
				g.It(desc, func() {
					checkParsing(g, line)
				})
			}
		})
	})
}
