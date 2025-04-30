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
	style   string
}

func New(label string) *Button {
	return &Button{
		id:    "btn-" + uuid.NewString(),
		Label: label,
		style: styles.ButtonPrimary, // Estilo padrão
	}
}

func (b *Button) OnClick(fn func()) {
	b.onClick = fn
	core.RegisterAction(b.id, fn)
}

func (b *Button) Render() template.HTML {
	return template.HTML(`
	<button id="` + b.id + `" class="` + b.style + `" 
		hx-post="/__glow/action?id=` + b.id + `" 
		hx-trigger="click" 
		hx-swap="innerHTML" 
		hx-target="#page-content"
		hx-include="input,textarea">` + b.Label + `</button>`)
}

func (b *Button) ID() string {
	return b.id
}

// NewPrimary cria um botão com a cor primária (roxo)
func NewPrimary(label string) *Button {
	btn := New(label)
	btn.style = styles.ButtonPrimary
	return btn
}

// NewDanger cria um botão com a cor de perigo (vermelho)
func NewDanger(label string) *Button {
	btn := New(label)
	btn.style = styles.ButtonDanger
	return btn
}

// NewSuccess cria um botão com a cor de sucesso (verde)
func NewSuccess(label string) *Button {
	btn := New(label)
	btn.style = styles.ButtonSuccess
	return btn
}

// NewWarning cria um botão com a cor de aviso (laranja)
func NewWarning(label string) *Button {
	btn := New(label)
	btn.style = styles.ButtonWarning
	return btn
}

// NewInfo cria um botão com a cor de informação (ciano)
func NewInfo(label string) *Button {
	btn := New(label)
	btn.style = styles.ButtonInfo
	return btn
}

// NewLight cria um botão com a cor clara (foreground)
func NewLight(label string) *Button {
	btn := New(label)
	btn.style = styles.ButtonLight
	return btn
}

// NewDark cria um botão com a cor escura (current-line)
func NewDark(label string) *Button {
	btn := New(label)
	btn.style = styles.ButtonDark
	return btn
}

// NewSecondary cria um botão com a cor secundária (rosa)
func NewSecondary(label string) *Button {
	btn := New(label)
	btn.style = styles.ButtonSecondary
	return btn
}

// NewYellow cria um botão com a cor amarela
func NewYellow(label string) *Button {
	btn := New(label)
	btn.style = styles.ButtonYellow
	return btn
}
