package element

import (
	"fmt"
	"goml/faces"
	"goml/parse"
)

const TAGform = "form"

type Form struct {
	Action       string `attr:"action"`
	Method       string `attr:"method"`
	Enctype      string `attr:"enctype"`
	Name         string `attr:"name"`
	Target       string `attr:"target"`
	NoValidate   bool   `attr:"novalidate"`
	Autocomplete string `attr:"autocomplete"`
	Class        string `attr:"class"`
	Id           string `attr:"id"`
	Children     []faces.FormChild
}

func (f Form) HTML() string {
	return fmt.Sprintf("<%s%s>%s</%s>", TAGform, f.ATTRIBUTES(), f.CHILDREN(), TAGform)
}

func (f Form) ATTRIBUTES() string {
	return parse.GetAttr(f)
}

func (f Form) CHILDREN() string {
	return parse.GetChildren(f.Children)
}

func (f Form) VALIDATE() bool {
	return true
}
