package element

import (
	"fmt"
	"goml/faces"
	"goml/parse"
)

type Span struct {
	Class    string `attr:"class"`
	Children []faces.Element
}

func (s Span) HTML() string {
	return fmt.Sprintf("<span%s>%s</span>", s.ATTRIBUTES(), s.CHILDREN())
}

func (s Span) ATTRIBUTES() string {
	return parse.GetAttr(s)
}

func (s Span) CHILDREN() string {
	return parse.GetChildren(s.Children)
}

func (s Span) VALIDATE() bool {
	return true
}

func (s Span) CanBeButtonChild() {}
func (s Span) CanBeFormChild()   {}
