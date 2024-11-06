package element

import (
	"fmt"
	"goml/faces"
	"goml/parse"
)

type Div struct {
	Class    string
	Children faces.Element
}

func (d Div) HTML() string {
	return fmt.Sprintf("<div%s>%s</div>", d.ATTRIBUTES(), d.CHILDREN())
}

func (d Div) ATTRIBUTES() string {
	return parse.GetAttr(d)
}

func (d Div) CHILDREN() string {
	return parse.GetChild(d.Children)
}
