package vbox

import (
	"html/template"

	"github.com/devalexandre/glow/components"
	"github.com/devalexandre/glow/components/styles"
	"github.com/google/uuid"
)

type VBox struct {
	id       string
	Children []components.Component
}

func New(children ...components.Component) *VBox {
	return &VBox{
		id:       "vbox-" + uuid.NewString(),
		Children: children,
	}
}

func (v *VBox) Render() template.HTML {
	content := `<div id="` + v.id + `" class="` + styles.VBoxBase + `">`
	for _, child := range v.Children {
		content += string(child.Render())
	}
	content += `</div>`

	return template.HTML(content)
}

func (v *VBox) ID() string {
	return v.id
}
