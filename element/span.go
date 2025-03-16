package element

import (
	"fmt"
	"goml/faces"
	"goml/parse"
)

type Span struct {
	Class    string `attr:"class"`
	Children faces.Element
}

func (s Span) HTML() string {
	return fmt.Sprintf("<span%s>%s</span>", s.ATTRIBUTES(), s.CHILDREN())
}

func (s Span) ATTRIBUTES() string {
	return parse.GetAttr(s)
}

func (s Span) CHILDREN() string {
	return parse.GetChild(s.Children)
}

func (s Span) VALIDATE() bool {
	//	Check if the button has a DIV element as a child
	return true
}

func (s Span) CanBeButtonChild() {}
