package core

type ButtonState interface {
	ID() string
	MarkClicked()
}
