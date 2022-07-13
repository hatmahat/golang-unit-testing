package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CalculatorTestSuite struct {
	suite.Suite
	calcService Calculator
}

func (suite *CalculatorTestSuite) SetupTest() {
	suite.calcService = Calculator{}
}

func (suite *CalculatorTestSuite) TestCalculator_Add() {
	testTable := []struct {
		num1 int
		num2 int
		res  int
	}{
		{4, 4, 8},
		{-4, 5, -1},
	}
	for _, table := range testTable {
		suite.calcService.Num1 = table.num1
		suite.calcService.Num2 = table.num2
		actual, err := suite.calcService.Add()
		if err != nil {
			assert.NotNil(suite.T(), err)
			assert.Equal(suite.T(), table.res, actual)
		} else {
			assert.Nil(suite.T(), err)
			assert.Equal(suite.T(), table.res, actual)
		}
	}
	// suite.calcService.Num1 = 4
	// suite.calcService.Num2 = 4
	// actual, err := suite.calcService.Add()
	// expected := 8
	// assert.Equal(suite.T(), expected, actual, "it should return 8")
}

func (suite *CalculatorTestSuite) TestCalculator_Sub() {
	suite.calcService.Num1 = 5
	suite.calcService.Num2 = 4
	actual := suite.calcService.Sub()
	expected := 1
	assert.Equal(suite.T(), expected, actual, "it should return 1")
}

func TestCalculatorTestSuite(t *testing.T) {
	suite.Run(t, new(CalculatorTestSuite))
}

// func TestCalculator_Add(t *testing.T) {
// 	// t.Run("Calculator Add Operator Testing", func(t *testing.T) {
// 	// 	calc := Calculator{
// 	// 		Num1: 2,
// 	// 		Num2: 3,
// 	// 	}
// 	// 	expected := 5
// 	// 	actual := calc.Add()
// 	// 	if expected != actual {
// 	// 		t.Error("The result should be", expected)
// 	// 	}
// 	// })
// 	myCalc := Calculator{Num1: 4, Num2: 4}
// 	assert.Equal(t, 8, myCalc.Add(), "it should return 8")
// }

// func TestCalculator_Sub(t *testing.T) {
// 	// t.Run("Calculator Substraction operator testing", func(t *testing.T) {
// 	// 	calc := Calculator{
// 	// 		Num1: 2,
// 	// 		Num2: 3,
// 	// 	}
// 	// 	expected := -1
// 	// 	actual := calc.Sub()
// 	// 	if expected != actual {
// 	// 		t.Error("The result should be", expected)
// 	// 	}
// 	// })
// 	myCalc := Calculator{Num1: 5, Num2: 4}
// 	assert.Equal(t, 1, myCalc.Sub(), "it should return 1")
// }

// func TestMainApp_Add(t *testing.T) {
// 	os.Args = []string{"test", "-num1", "2", "-num2", "3", "-opr", "add"}
// 	main()
// }

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
