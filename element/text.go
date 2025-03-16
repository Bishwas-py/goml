package element

import (
	"html"
)

type Text struct {
	Content string
}

func (t Text) HTML() string {
	return html.EscapeString(t.Content)
}

func (t Text) ATTRIBUTES() string {
	return ""
}

func (t Text) CHILDREN() string {
	return ""
}

func (t Text) VALIDATE() bool {
	return true
}

func (t Text) CanBeButtonChild() {}

func T(content string) Text {
	return Text{Content: content}
}
