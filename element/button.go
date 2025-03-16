package element

import (
	"fmt"
	"goml/faces"
	"goml/parse"
)

type Button struct {
	Autofocus      bool   `attr:"autofocus"`
	Disabled       bool   `attr:"disabled"`
	Form           string `attr:"form"`
	FormAction     string `attr:"form_action"`
	FormEnctype    string `attr:"form_enctype"`
	FormMethod     string `attr:"form_method"`
	FormNoValidate bool   `attr:"form_no_validate"`
	FormTarget     string `attr:"form_target"`
	Name           string `attr:"name"`
	Type           string `attr:"type"`
	Value          string `attr:"value"`
	Count          int    `attr:"count"`
	Children       []faces.ButtonChild
}

func (b Button) HTML() string {
	return fmt.Sprintf("<button%s>%s</button>", b.ATTRIBUTES(), b.CHILDREN())
}

func (b Button) ATTRIBUTES() string {
	return parse.GetAttr(b)
}

func (b Button) CHILDREN() string {
	return parse.GetChildren(b.Children)
}

func (b Button) VALIDATE() bool {
	return true
}
