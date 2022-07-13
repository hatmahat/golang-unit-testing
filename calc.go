package main

import "errors"

type Calculator struct {
	Num1 int
	Num2 int
}

func (c *Calculator) Add() (int, error) {
	if c.Num1 < 0 || c.Num2 < 0 {
		return -1, errors.New("numbers can't be negative")
	}
	result := c.Num1 + c.Num2
	return result, nil
}

func (c *Calculator) Sub() int {
	return c.Num1 - c.Num2
}
