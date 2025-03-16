package element

import (
	"fmt"
	"goml/faces"
	"goml/parse"
)

const TAGBody = "body"

type Body struct {
	Class    string `attr:"class"`
	Children []faces.Element
}

func (d Body) HTML() string {
	return fmt.Sprintf("<%s%s>%s</%s>", TAGBody, d.ATTRIBUTES(), d.CHILDREN(), TAGBody)
}

func (d Body) ATTRIBUTES() string {
	return parse.GetAttr(d)
}

func (d Body) CHILDREN() string {
	return parse.GetChildren(d.Children)
}

func (d Body) VALIDATE() bool {
	return true
}
