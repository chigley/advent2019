package main

import (
	"errors"
	"fmt"
)

const (
	opAdd  = 1
	opMult = 2
	opHalt = 99
)

var errHalt = errors.New("halt")

type computer struct {
	memory []int
	pc     int
}

func (c *computer) run() error {
	for {
		if err := c.runOp(); err == errHalt {
			return nil
		} else if err != nil {
			return err
		}
	}
}

func (c *computer) runOp() error {
	opCode, err := c.read(c.pc)
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
		return fmt.Errorf("invalid op code %d", opCode)
	}
}

func (c *computer) read(pos int) (int, error) {
	if pos >= len(c.memory) {
		return 0, fmt.Errorf("index %d is out of bounds", pos)
	}
	return c.memory[pos], nil
}

func (c *computer) readPointer(pos int) (int, error) {
	i, err := c.read(pos)
	if err != nil {
		return 0, err
	}
	return c.read(i)
}

func (c *computer) write(pos, val int) error {
	if pos >= len(c.memory) {
		return fmt.Errorf("index %d is out of bounds", pos)
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

	dst, err := c.read(c.pc + 3)
	if err != nil {
		return err
	}

	switch opCode {
	case opAdd:
		return c.write(dst, arg1+arg2)
	case opMult:
		return c.write(dst, arg1*arg2)
	default:
		return fmt.Errorf("invalid binary op code %d", opCode)
	}
}
