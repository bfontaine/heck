package main

import (
	"errors"
	"strconv"
)

type OpCode uint8

type Unop struct {
	Execute func(int64) int64
}

type Binop struct {
	Execute func(int64, int64) int64
}

const (
	_ OpCode = iota

	begin_binop
	Add
	Sub
	Mult
	Div
	Mod
	end_binop

	begin_unop
	Neg
	Tilde
	end_unop
)

func isBinop(n OpCode) bool { return begin_binop < n && n < end_binop }
func isUnop(n OpCode) bool  { return begin_unop < n && n < end_unop }

func makeUnop(exe func(int64) int64) Unop          { return Unop{Execute: exe} }
func makeBinop(exe func(int64, int64) int64) Binop { return Binop{Execute: exe} }

var Unops = map[OpCode]Unop{
	Neg:   makeUnop(func(a int64) int64 { return -a }),
	Tilde: makeUnop(func(a int64) int64 { return ^a }),
}
var Binops = map[OpCode]Binop{
	Add:  makeBinop(func(a, b int64) int64 { return a + b }),
	Sub:  makeBinop(func(a, b int64) int64 { return a - b }),
	Mult: makeBinop(func(a, b int64) int64 { return a * b }),
	Div:  makeBinop(func(a, b int64) int64 { return a / b }),
	Mod:  makeBinop(func(a, b int64) int64 { return a % b }),
}

var (
	ErrStackUnderflow  = errors.New("Stack underflow")
	ErrUnknownOperator = errors.New("Unknown operator")
)

type Heck struct {
	stack     []int64
	stackSize int
}

func (h *Heck) pop() (int64, error) {
	if h.stackSize == 0 {
		return 0, ErrStackUnderflow
	}

	h.stackSize--

	n := h.stack[h.stackSize]
	h.stack = h.stack[:h.stackSize]

	return n, nil
}

func (h *Heck) push(n int64) {
	h.stack = append(h.stack, n)
	h.stackSize++
}

// AddOp adds a new operator to the current expression. The values used by the
// operator should already be in the stack.
func (h *Heck) AddOp(op OpCode) (err error) {
	if isBinop(op) {
		if b, ok := Binops[op]; !ok {
			err = ErrUnknownOperator
		} else {
			var e1, e2 int64

			if e2, err = h.pop(); err != nil {
				return
			}

			if e1, err = h.pop(); err != nil {
				return
			}

			h.push(b.Execute(e1, e2))
		}
	} else if isUnop(op) {
		if u, ok := Unops[op]; !ok {
			err = ErrUnknownOperator
		} else {
			e, _ := h.pop()
			h.push(u.Execute(e))
		}
	} else {
		err = ErrUnknownOperator
	}

	return
}

// AddValue adds a new value to the current expression.
func (h *Heck) AddValue(s string) error {
	if v, err := strconv.ParseInt(s, 16, 64); err != nil {
		return err
	} else {
		h.push(v)
	}

	return nil
}

// Value returns the current expression's value
func (h *Heck) Value() int64 {
	if h.stackSize < 1 {
		return 0
	}
	return h.stack[h.stackSize-1]
}
