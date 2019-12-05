package intcode

import (
	"errors"
	"fmt"
)

type opCode int

const (
	opAdd    opCode = 1
	opMult          = 2
	opInput         = 3
	opOutput        = 4
	opHalt          = 99
)

type paramMode int

const (
	modePosition paramMode = iota
	modeImmediate
)

var errHalt = errors.New("intcode: halt")

type computer struct {
	program []int

	memory  []int
	pc      int
	inputs  []int
	outputs []int
}

func New(program []int) *computer {
	return &computer{
		program: program,
	}
}

func (c *computer) Run(inputs []int) ([]int, error) {
	c.memory = append([]int(nil), c.program...)
	c.pc = 0
	c.inputs = append([]int(nil), inputs...)
	c.outputs = nil
	for {
		if err := c.runOp(); err == errHalt {
			return c.outputs, nil
		} else if err != nil {
			return nil, err
		}
	}
}

func (c *computer) runOp() error {
	instr, err := c.Read(c.pc)
	if err != nil {
		return err
	}

	op, encodedModes := separateInstr(instr)

	switch op {
	case opAdd, opMult:
		modes, err := parseModes(encodedModes, 2)
		if err != nil {
			return err
		}
		if err := c.binaryOp(op, modes); err != nil {
			return err
		}
		c.pc += 4
		return nil
	case opInput, opOutput:
		modes, err := parseModes(encodedModes, 1)
		if err != nil {
			return err
		}
		if err := c.unaryOp(op, modes); err != nil {
			return err
		}
		c.pc += 2
		return nil
	case opHalt:
		return errHalt
	default:
		return fmt.Errorf("intcode: invalid op code %d", op)
	}
}

func (c *computer) Read(pos int) (int, error) {
	if pos >= len(c.memory) {
		return 0, fmt.Errorf("intcode: index %d is out of bounds", pos)
	}
	return c.memory[pos], nil
}

func (c *computer) readPointer(pos int) (int, error) {
	i, err := c.Read(pos)
	if err != nil {
		return 0, err
	}
	return c.Read(i)
}

func (c *computer) write(pos, val int) error {
	if pos >= len(c.memory) {
		return fmt.Errorf("intcode: index %d is out of bounds", pos)
	}
	c.memory[pos] = val
	return nil
}

func (c *computer) binaryOp(op opCode, modes []paramMode) error {
	arg1, err := c.readPointer(c.pc + 1)
	if err != nil {
		return err
	}

	arg2, err := c.readPointer(c.pc + 2)
	if err != nil {
		return err
	}

	dst, err := c.Read(c.pc + 3)
	if err != nil {
		return err
	}

	switch op {
	case opAdd:
		return c.write(dst, arg1+arg2)
	case opMult:
		return c.write(dst, arg1*arg2)
	default:
		return fmt.Errorf("intcode: invalid binary op code %d", op)
	}
}

func (c *computer) unaryOp(op opCode, modes []paramMode) error {
	arg, err := c.Read(c.pc + 1)
	if err != nil {
		return err
	}

	switch op {
	case opInput:
		if len(c.inputs) == 0 {
			return errors.New("intcode: input instruction has no input to read")
		}

		var input int
		input, c.inputs = c.inputs[0], c.inputs[1:]

		return c.write(arg, input)
	case opOutput:
		val, err := c.Read(arg)
		if err != nil {
			return err
		}
		c.outputs = append(c.outputs, val)
		return nil
	default:
		return fmt.Errorf("intcode: invalid unary op code %d", op)
	}
}

func separateInstr(instr int) (opCode, int) {
	op := opCode(instr % 100)
	modes := instr / 100
	return op, modes
}

func parseModes(modes, n int) ([]paramMode, error) {
	out := make([]paramMode, n)
	for i := 0; i < n; i++ {
		mode := modes % 10
		if mode > 2 {
			return nil, fmt.Errorf("intcode: unrecognised parameter mode %d", mode)
		}
		out[i] = paramMode(mode)
		modes = modes / 10
	}
	return out, nil
}
