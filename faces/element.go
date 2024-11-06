package faces

type Element interface {
	HTML() string
	ATTRIBUTES() string
	CHILDREN() string
}
