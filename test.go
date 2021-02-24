package main

import (
	"fmt"
	"math/rand"

	L "./basic"
	FC "./flow-control"
	MI "./methods-and-interfaces"
	MT "./more-types"
)

func main() {
	L.Say()
	fmt.Println("My favorite number is: ", rand.Intn(10))
	fmt.Println("Sum of 3 & 4 is: ", L.Add(3, 4))
	a, b := L.Swap(55, 11)
	fmt.Println("Swap ", a, b)
	x, y := L.Split(44)
	fmt.Println("Split 44 into ", x, y)

	L.Show()
	L.ShowVariablesWithInitializer()
	L.ShortVariableDeclaration()
	L.ShowBasicTypes()
	L.ShowTypeConversion()

	FC.ShowForLoop()
	FC.ShowIf()
	FC.ShowIfElse()
	FC.ShowMySqrt()
	FC.ShowSwitch()
	FC.ShowDefer()

	MT.ShowPointer()
	MT.ShowStruct()
	MT.ShowArray()
	MT.ShowSlice()
	MT.ShowMap()
	MT.ShowFunctionValues()

	MI.ShowMethods()
	MI.ShowInterface()
	MI.ShowStringer()
	MI.ShowError()
	MI.ShowReader()
}
