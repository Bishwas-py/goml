package parse

import (
	"goml/faces"
)

func GetChild(children faces.Element) string {
	if children == nil {
		return ""
	}
	return children.HTML()
}

func GetChildren[T faces.Element](children []T) string {
	var html string
	for _, child := range children {
		html += GetChild(child)
	}
	return html
}

func GetGroupChildren[T faces.Element](children []T) string {
	var html string
	for _, child := range children {
		html += GetChild(child)
	}
	return html
}
