package text

import (
	"html/template"
	"time"

	"github.com/devalexandre/glow/components"
	"github.com/devalexandre/glow/components/styles"
	"github.com/google/uuid"
)

type Text struct {
	id      string
	Content string
}

func New(content string) *Text {
	t := &Text{
		id:      "text-" + uuid.NewString(),
		Content: content,
	}

	return t
}

func (t *Text) Render() template.HTML {
	return template.HTML(`<div id="` + t.id + `" class="` + styles.TextPrimary + `">` + t.Content + `</div>`)
}

func (t *Text) ID() string {
	return t.id
}

func (t *Text) SetContent(content string) {
	t.Content = content
}

// SetContentWithValue atualiza o conteúdo do texto com o valor de um componente
func (t *Text) SetContentWithValue(valueComponent components.ValueComponent, format string) {
	value := valueComponent.GetValue()
	if format != "" {
		t.Content = value + " - " + time.Now().Format(format)
	} else {
		t.Content = value
	}
}
