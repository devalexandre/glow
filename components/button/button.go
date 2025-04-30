package button

import (
	"html/template"

	"github.com/devalexandre/glow/components/core"
	"github.com/devalexandre/glow/components/styles"
	"github.com/google/uuid"
)

type Button struct {
	id      string
	Label   string
	onClick func()
}

func New(label string) *Button {
	return &Button{
		id:    "btn-" + uuid.NewString(),
		Label: label,
	}
}

func (b *Button) OnClick(fn func()) {
	b.onClick = fn
	core.RegisterAction(b.id, fn)
}

func (b *Button) Render() template.HTML {
	return template.HTML(`
	<button id="` + b.id + `" class="` + styles.ButtonPrimary + `" 
		hx-post="/__glow/action?id=` + b.id + `" 
		hx-trigger="click" 
		hx-swap="innerHTML" 
		hx-target="#page-content"
		hx-include="input,textarea">` + b.Label + `</button>`)
}

func (b *Button) ID() string {
	return b.id
}
