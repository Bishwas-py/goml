package faces

type Element interface {
	HTML() string
	ATTRIBUTES() string
	CHILDREN() string
	VALIDATE() bool
}

type ButtonChild interface {
	Element
	CanBeButtonChild()
}

type FormChild interface {
	Element
	CanBeFormChild()
}
