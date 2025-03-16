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
		Children: []faces.ButtonChild{
			element.Span{
				Class: "bg-red-500 h-7 w-7",
				Children: []faces.Element{
					element.T("Hello"),
				},
			},
		},
	}

	body := element.Body{
		Children: []faces.Element{
			button,
		},
	}
	fmt.Printf("HTML: \n %s!\n", body.HTML())
}
