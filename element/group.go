package element

import (
	"fmt"
	"goml/faces"
	"goml/parse"
)

type Group struct {
	Tag      string
	Id       string `attr:"id"`
	Class    string `attr:"class"`
	Style    string `attr:"style"`
	Children []faces.Element
}

func (d Group) HTML() string {
	if d.Tag == "" {
		return fmt.Sprintf(d.CHILDREN())
	}
	return fmt.Sprintf("<%s%s>%s</%s>", d.Tag, d.ATTRIBUTES(), d.CHILDREN(), d.Tag)
}

func (d Group) ATTRIBUTES() string {
	return parse.GetAttr(d)
}

func (d Group) CHILDREN() string {
	return parse.GetGroupChildren(d.Children)
}
