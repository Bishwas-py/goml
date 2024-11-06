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
