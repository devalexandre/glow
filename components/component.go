package components

import "html/template"

type Component interface {
	Render() template.HTML
	ID() string
}

// ValueComponent representa um componente que possui um valor que pode ser obtido
type ValueComponent interface {
	Component
	GetValue() string
}
