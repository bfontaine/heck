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
func isUnop(n OpCode) bool  { return begin_unop < n && n < end_binop }

func MakeUnop(exe func(int64) int64) Unop          { return Unop{Execute: exe} }
func MakeBinop(exe func(int64, int64) int64) Binop { return Binop{Execute: exe} }

var Unops = map[OpCode]Unop{
	Neg:   MakeUnop(func(a int64) int64 { return -a }),
	Tilde: MakeUnop(func(a int64) int64 { return ^a }),
}
var Binops = map[OpCode]Binop{
	Add:  MakeBinop(func(a, b int64) int64 { return a + b }),
	Sub:  MakeBinop(func(a, b int64) int64 { return a - b }),
	Mult: MakeBinop(func(a, b int64) int64 { return a * b }),
	Div:  MakeBinop(func(a, b int64) int64 { return a / b }),
	Mod:  MakeBinop(func(a, b int64) int64 { return a % b }),
}

var (
	errStackUnderflow  = errors.New("Stack underflow")
	errUnknownOperator = errors.New("Unknown operator")
)

type Heck struct {
	Stack     []int64
	StackSize int
}

func (h *Heck) pop() (int64, error) {
	if h.StackSize == 0 {
		return 0, errStackUnderflow
	}

	h.StackSize--

	n := h.Stack[h.StackSize]
	h.Stack = h.Stack[:h.StackSize]

	return n, nil
}

func (h *Heck) push(n int64) {
	h.Stack = append(h.Stack, n)
	h.StackSize++
}

func (h *Heck) AddOp(op OpCode) (err error) {
	if isBinop(op) {
		if b, ok := Binops[op]; !ok {
			err = errUnknownOperator
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
			err = errUnknownOperator
		} else {
			e, _ := h.pop()
			h.push(u.Execute(e))
		}
	} else {
		err = errUnknownOperator
	}

	return
}

func (h *Heck) AddValue(s string) error {
	if v, err := strconv.ParseInt(s, 16, 64); err != nil {
		return err
	} else {
		h.push(v)
	}

	return nil
}

func (h *Heck) Value() int64 {
	return h.Stack[h.StackSize-1]
}

func (o OpCode) String() (s string) {
	switch o {
	case Add:
		s = "+"
	case Sub, Neg:
		s = "-"
	case Mult:
		s = "*"
	case Div:
		s = "/"
	case Mod:
		s = "%"
	case Tilde:
		s = "~"
	default:
		s = "{?}"
	}

	return
}
