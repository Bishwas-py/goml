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

func GetGroupChildren(children []faces.Element) string {
	var html string
	for _, child := range children {
		html += GetChild(child)
	}
	return html
}
