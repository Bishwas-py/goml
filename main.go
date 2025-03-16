package main

import (
	"fmt"
	"goml/element"
	"goml/faces"
)

func main() {
	// Create a form with a button
	form := element.Form{
		Action:     "/submit",
		Method:     "post",
		NoValidate: true,
		Class:      "w-full max-w-lg mx-auto p-6 bg-white rounded-lg shadow-md",
		Id:         "contact-form",
		Children: []faces.FormChild{
			element.Div{
				Class: "mb-4",
				Children: []faces.Element{
					element.Button{
						Type:     "submit",
						Disabled: false,
						Class:    "w-full py-2 px-4 bg-blue-500 text-white font-semibold rounded-lg shadow-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-400 focus:ring-opacity-75",
						Children: []faces.ButtonChild{
							element.T("Submit"),
						},
					},
				},
			},
			element.Div{
				Class: "mt-4 text-center",
				Children: []faces.Element{
					element.Span{
						Class: "text-sm text-gray-500",
						Children: []faces.Element{
							element.T("By submitting, you agree to our terms."),
						},
					},
				},
			},
		},
	}

	// Create another button outside the form
	button := element.Button{
		Type:     "button",
		Disabled: true,
		Count:    3,
		Class:    "mt-4 py-2 px-4 bg-gray-300 text-gray-500 font-semibold rounded-lg shadow-md cursor-not-allowed",
		Children: []faces.ButtonChild{
			element.Span{
				Class: "flex items-center justify-center",
				Children: []faces.Element{
					element.T("Cancel"),
				},
			},
		},
	}

	// Create a container div
	container := element.Div{
		Class: "container mx-auto px-4 py-8",
		Children: []faces.Element{
			element.Div{
				Class: "text-center mb-8",
				Children: []faces.Element{
					element.Span{
						Class: "text-3xl font-bold text-gray-800",
						Children: []faces.Element{
							element.T("Contact Us"),
						},
					},
				},
			},
			form,
			button,
			element.Div{
				Class: "mt-12 pt-6 border-t border-gray-200 text-center",
				Children: []faces.Element{
					element.Span{
						Class: "text-sm text-gray-500",
						Children: []faces.Element{
							element.T("© 2025 My Company"),
						},
					},
				},
			},
		},
	}

	// Add everything to the body
	body := element.Body{
		Class: "bg-gray-100 min-h-screen font-sans antialiased text-gray-900",
		Children: []faces.Element{
			container,
		},
	}

	// Output the generated HTML
	fmt.Printf("Generated HTML:\n%s\n", body.HTML())
}
