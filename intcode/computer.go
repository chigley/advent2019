package intcode

import (
	"errors"
	"fmt"
)

type opCode int

const (
	opAdd             opCode = 1
	opMult                   = 2
	opInput                  = 3
	opOutput                 = 4
	opJumpIfTrue             = 5
	opJumpIfFalse            = 6
	opLessThan               = 7
	opEquals                 = 8
	opAdjRelativeBase        = 9
	opHalt                   = 99
)

type paramMode int

const (
	modePosition paramMode = iota
	modeImmediate
)

var errHalt = errors.New("intcode: halt")

type Computer struct {
	program []int

	memory       []int
	pc           int
	relativeBase int
	inputs       <-chan int
	outputs      chan int
	err          error
}

func New(program []int) *Computer {
	return &Computer{
		program: program,
	}
}

func (c *Computer) Run(inputs []int) ([]int, error) {
	inChan := make(chan int)
	go func() {
		for _, in := range inputs {
			inChan <- in
		}
		close(inChan)
	}()

	var outputs []int
	for out := range c.RunInteractive(inChan, nil) {
		outputs = append(outputs, out)
	}

	if err := c.Err(); err != nil {
		return nil, err
	}

	return outputs, nil
}

func (c *Computer) RunInteractive(inputs chan int, done func()) chan int {
	c.memory = append([]int(nil), c.program...)
	c.pc = 0
	c.relativeBase = 0
	c.inputs = inputs
	c.outputs = make(chan int)

	go func() {
		for {
			c.err = c.runOp()
			if c.err != nil {
				close(c.outputs)
				if done != nil {
					done()
				}
				return
			}

		}
	}()

	return c.outputs
}

func (c *Computer) Err() error {
	if c.err == errHalt {
		return nil
	}
	return c.err
}

func (c *Computer) runOp() error {
	instr, err := c.Read(c.pc)
	if err != nil {
		return err
	}
	c.pc++

	op, encodedModes := separateInstr(instr)

	switch op {
	case opJumpIfTrue, opJumpIfFalse:
		modes, err := decodeModes(encodedModes, 2)
		if err != nil {
			return err
		}
		if err := c.binaryOp(op, modes); err != nil {
			return err
		}
		return nil
	case opAdd, opMult, opLessThan, opEquals:
		modes, err := decodeModes(encodedModes, 3)
		if err != nil {
			return err
		}
		if err := c.binaryOpWithDest(op, modes); err != nil {
			return err
		}
		return nil
	case opInput, opOutput, opAdjRelativeBase:
		modes, err := decodeModes(encodedModes, 1)
		if err != nil {
			return err
		}
		if err := c.unaryOp(op, modes); err != nil {
			return err
		}
		return nil
	case opHalt:
		return errHalt
	default:
		return fmt.Errorf("intcode: invalid op code %d", op)
	}
}

func (c *Computer) Read(pos int) (int, error) {
	if pos >= len(c.memory) {
		return 0, fmt.Errorf("intcode: index %d is out of bounds", pos)
	}
	return c.memory[pos], nil
}

func (c *Computer) readPointer(pos int) (int, error) {
	i, err := c.Read(pos)
	if err != nil {
		return 0, err
	}
	return c.Read(i)
}

func (c *Computer) write(pos, val int) error {
	if pos >= len(c.memory) {
		return fmt.Errorf("intcode: index %d is out of bounds", pos)
	}
	c.memory[pos] = val
	return nil
}

func (c *Computer) binaryOp(op opCode, modes []paramMode) error {
	if len(modes) != 2 {
		return fmt.Errorf("intcode: binary op got %d paremeter modes, expected 2", len(modes))
	}

	arg1, err := c.readArg(modes[0])
	if err != nil {
		return err
	}

	arg2, err := c.readArg(modes[1])
	if err != nil {
		return err
	}

	switch op {
	case opJumpIfTrue:
		if arg1 != 0 {
			c.pc = arg2
		}
		return nil
	case opJumpIfFalse:
		if arg1 == 0 {
			c.pc = arg2
		}
		return nil
	default:
		return fmt.Errorf("intcode: invalid binary op code %d", op)
	}
}

func (c *Computer) binaryOpWithDest(op opCode, modes []paramMode) error {
	if len(modes) != 3 {
		return fmt.Errorf("intcode: binary op got %d paremeter modes, expected 3", len(modes))
	}

	arg1, err := c.readArg(modes[0])
	if err != nil {
		return err
	}

	arg2, err := c.readArg(modes[1])
	if err != nil {
		return err
	}

	dstMode := modes[2]

	switch op {
	case opAdd:
		return c.writeArg(dstMode, arg1+arg2)
	case opMult:
		return c.writeArg(dstMode, arg1*arg2)
	case opLessThan:
		if arg1 < arg2 {
			return c.writeArg(dstMode, 1)
		}
		return c.writeArg(dstMode, 0)
	case opEquals:
		if arg1 == arg2 {
			return c.writeArg(dstMode, 1)
		}
		return c.writeArg(dstMode, 0)
	default:
		return fmt.Errorf("intcode: invalid binary op code %d", op)
	}
}

func (c *Computer) unaryOp(op opCode, modes []paramMode) error {
	if len(modes) != 1 {
		return fmt.Errorf("intcode: unary op got %d paremeter modes, expected 1", len(modes))
	}

	switch op {
	case opInput:
		return c.writeArg(modes[0], <-c.inputs)
	case opOutput:
		arg, err := c.readArg(modes[0])
		if err != nil {
			return err
		}
		c.outputs <- arg
		return nil
	case opAdjRelativeBase:
		arg, err := c.readArg(modes[0])
		if err != nil {
			return err
		}
		c.relativeBase += arg
		return nil
	default:
		return fmt.Errorf("intcode: invalid unary op code %d", op)
	}
}

func (c *Computer) readArg(mode paramMode) (int, error) {
	switch mode {
	case modePosition:
		arg, err := c.readPointer(c.pc)
		if err != nil {
			return 0, err
		}
		c.pc++
		return arg, nil
	case modeImmediate:
		arg, err := c.Read(c.pc)
		if err != nil {
			return 0, err
		}
		c.pc++
		return arg, nil
	default:
		return 0, fmt.Errorf("intcode: unrecognised parameter mode %d", mode)
	}
}

func (c *Computer) writeArg(mode paramMode, val int) error {
	if mode != modePosition {
		return fmt.Errorf("intcode: invalid parameter mode for write: %d", mode)
	}

	dst, err := c.readArg(modeImmediate)
	if err != nil {
		return err
	}
	return c.write(dst, val)
}

func separateInstr(instr int) (opCode, int) {
	op := opCode(instr % 100)
	modes := instr / 100
	return op, modes
}

func decodeModes(modes, n int) ([]paramMode, error) {
	out := make([]paramMode, n)
	for i := 0; i < n; i++ {
		out[i] = paramMode(modes % 10)
		modes = modes / 10
	}
	return out, nil
}
