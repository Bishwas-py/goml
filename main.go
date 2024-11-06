package main

import (
	"fmt"
	"goml/element"
	"goml/faces"
)

func main() {
	button := element.Button{
		Disabled: true,
		Count:    2,
		Children: element.Group{
			Tag:   "div",
			Class: "flex gap-2",
			Children: []faces.Element{
				element.Div{
					Class: "bg-red-500",
				},
				element.Div{
					Class: "bg-blue-500",
				},
			},
		},
	}
	fmt.Printf("HTML: \n %s!\n", button.HTML())
}
