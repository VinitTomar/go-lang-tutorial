package basic

import "fmt"

var i, j int = 1, 2

// ShowVariablesWithInitializer ...
func ShowVariablesWithInitializer() {
	var c, python, java = true, false, "no!"

	fmt.Println(c, python, java)
}

// ShortVariableDeclaration ...
func ShortVariableDeclaration() {
	c, python, java := true, false, "no!"

	fmt.Println(c, python, java)
}
