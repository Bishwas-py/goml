package main

import (
	"fmt"
	"goml/element"
)

func main() {
	button := element.Button{
		Disabled: true,
		Count:    2,
		Children: element.Div{},
	}
	fmt.Printf("HTML: \n %s!\n", button.HTML())
}
