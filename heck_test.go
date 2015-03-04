package main

import (
	"github.com/franela/goblin"
	o "github.com/onsi/gomega"
	"testing"
)

func TestHeck(t *testing.T) {
	g := goblin.Goblin(t)

	o.RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Heck", func() {

		g.Describe(".AddValue()", func() {
			var h *Heck

			g.BeforeEach(func() { h = &Heck{} })

			g.It("Should parse the value as hexadecimal", func() {
				o.Expect(h.AddValue("2A")).To(o.BeNil())
				o.Expect(h.Value()).To(o.Equal(int64(42)))
			})

			g.It("Should return an error if the value can't be parsed", func() {
				o.Expect(h.AddValue("2Ax")).NotTo(o.BeNil())
			})
		})

		g.Describe(".AddOp()", func() {
			var h *Heck

			g.BeforeEach(func() { h = &Heck{} })

			g.It("Should return ErrStackUnderflow without value", func() {
				o.Expect(h.AddOp(Add)).To(o.Equal(ErrStackUnderflow))
			})

			g.It("Should return ErrStackUnderflow if binop with one value", func() {
				o.Expect(h.AddValue("42")).To(o.BeNil())
				o.Expect(h.AddOp(Add)).To(o.Equal(ErrStackUnderflow))
			})

			g.It("Should return ErrUnknownOperator if the op doesn't exist", func() {
				var Foo OpCode
				o.Expect(h.AddOp(Foo)).To(o.Equal(ErrUnknownOperator))
			})

			g.It("Should return nil if unop and stack size >=1", func() {
				o.Expect(h.AddValue("2A")).To(o.BeNil())
				o.Expect(h.AddOp(Neg)).To(o.BeNil())
			})

			g.It("Should return nil if binop and stack size >=2", func() {
				o.Expect(h.AddValue("2A")).To(o.BeNil())
				o.Expect(h.AddValue("3B")).To(o.BeNil())
				o.Expect(h.AddOp(Add)).To(o.BeNil())
			})

			g.It("Should update the top value if unop", func() {
				o.Expect(h.AddValue("2A")).To(o.BeNil())
				o.Expect(h.AddOp(Neg)).To(o.BeNil())
				o.Expect(h.Value()).To(o.Equal(int64(-42)))
			})

			g.It("Should apply a binop operation", func() {
				o.Expect(h.AddValue("2A")).To(o.BeNil())
				o.Expect(h.AddValue("3A")).To(o.BeNil())
				o.Expect(h.AddOp(Add)).To(o.BeNil())
				o.Expect(h.Value()).To(o.Equal(int64(100)))
			})
		})

		g.Describe(".Value()", func() {
			var h *Heck

			g.BeforeEach(func() { h = &Heck{} })

			g.It("Should return 0 by default", func() {
				o.Expect(h.Value()).To(o.Equal(int64(0)))
			})

			g.It("Should return the latest value if no op was used", func() {
				o.Expect(h.AddValue("2A")).To(o.BeNil())
				o.Expect(h.AddValue("17")).To(o.BeNil()) // 23 in hexadecimal
				o.Expect(h.Value()).To(o.Equal(int64(23)))
			})

			g.It("Should not modify the internal state", func() {
				o.Expect(h.AddValue("2A")).To(o.BeNil())
				o.Expect(h.Value()).To(o.Equal(int64(42)))
				o.Expect(h.Value()).To(o.Equal(int64(42)))
			})
		})
	})

}
