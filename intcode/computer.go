package intcode

import (
	"errors"
	"fmt"
)

const (
	opAdd  = 1
	opMult = 2
	opHalt = 99
)

var errHalt = errors.New("intcode: halt")

type computer struct {
	program []int
	memory  []int
	pc      int
}

func New(program []int) *computer {
	return &computer{
		program: program,
	}
}

func (c *computer) Run() error {
	c.memory = append([]int(nil), c.program...)
	c.pc = 0
	for {
		if err := c.runOp(); err == errHalt {
			return nil
		} else if err != nil {
			return err
		}
	}
}

func (c *computer) runOp() error {
	opCode, err := c.Read(c.pc)
	if err != nil {
		return err
	}

	switch opCode {
	case opAdd, opMult:
		if err := c.binaryOp(opCode); err != nil {
			return err
		}
		c.pc += 4
		return nil
	case opHalt:
		return errHalt
	default:
		return fmt.Errorf("intcode: invalid op code %d", opCode)
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

func (c *computer) binaryOp(opCode int) error {
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

	switch opCode {
	case opAdd:
		return c.write(dst, arg1+arg2)
	case opMult:
		return c.write(dst, arg1*arg2)
	default:
		return fmt.Errorf("intcode: invalid binary op code %d", opCode)
	}
}
