package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNumber(t *testing.T) {
	assert := assert.New(t)

	var actual string = FizzBuzz(1)
	assert.Equal(actual, "1")
}

func TestFizz(t *testing.T) {
	assert := assert.New(t)

	var actual string = FizzBuzz(3)
	assert.Equal(actual, "fizz")
}

func TestBuzz(t *testing.T) {
	assert := assert.New(t)

	var actual string = FizzBuzz(5)
	assert.Equal(actual, "buzz")
}

func TestFizzBuzz(t *testing.T) {
	assert := assert.New(t)

	var actual string = FizzBuzz(15)
	assert.Equal(actual, "fizzbuzz")
}