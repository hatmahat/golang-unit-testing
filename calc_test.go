package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculator_Add(t *testing.T) {
	// t.Run("Calculator Add Operator Testing", func(t *testing.T) {
	// 	calc := Calculator{
	// 		Num1: 2,
	// 		Num2: 3,
	// 	}
	// 	expected := 5
	// 	actual := calc.Add()
	// 	if expected != actual {
	// 		t.Error("The result should be", expected)
	// 	}
	// })
	myCalc := Calculator{Num1: 4, Num2: 4}
	assert.Equal(t, 8, myCalc.Add(), "it should return 8")
}

func TestCalculator_Sub(t *testing.T) {
	// t.Run("Calculator Substraction operator testing", func(t *testing.T) {
	// 	calc := Calculator{
	// 		Num1: 2,
	// 		Num2: 3,
	// 	}
	// 	expected := -1
	// 	actual := calc.Sub()
	// 	if expected != actual {
	// 		t.Error("The result should be", expected)
	// 	}
	// })
	myCalc := Calculator{Num1: 5, Num2: 4}
	assert.Equal(t, 1, myCalc.Sub(), "it should return 1")
}

func TestMainApp_Add(t *testing.T) {
	os.Args = []string{"test", "-num1", "2", "-num2", "3", "-opr", "add"}
	main()
}

/*
go test
go test -cover
go test -v
go test ./...
go test ./... -v
go test ./... -v -coverprofile=cover.out

go test ./... -cover -coverprofile=c.out
go tool cover -html=c.out -o "coverage.html"

go tool cover -html=cover.out

*/
