package input

import (
	"html/template"

	"github.com/devalexandre/glow/components/core"
	"github.com/devalexandre/glow/components/styles"
	"github.com/google/uuid"
)

type TextInput struct {
	id          string
	Label       string
	Value       string
	Placeholder string
	Multiline   bool
	Error       string
	width       string
	actionID    string
}

func New(label, initial string) *TextInput {
	return &TextInput{
		id:    "input-" + uuid.NewString(),
		Label: label,
		Value: initial,
		width: "w-full",
	}
}

func (t *TextInput) SetPlaceholder(p string) {
	t.Placeholder = p
}

func (t *TextInput) SetMultiline(m bool) {
	t.Multiline = m
}

func (t *TextInput) SetError(err string) {
	t.Error = err
}

func (t *TextInput) OnChange(fn func()) {
	t.actionID = "action-" + t.id

	// Registrar uma ação que atualiza o valor do input antes de chamar a função do usuário
	core.RegisterAction(t.actionID, func() {
		// Atualizar o valor do input com o valor do formulário
		values := core.GetFormValues()
		if val, ok := values[t.id]; ok {
			t.Value = val
		}

		// Chamar a função do usuário
		if fn != nil {
			fn()
		}
	})
}

func (t *TextInput) Render() template.HTML {
	labelHTML := ""
	if t.Label != "" {
		labelHTML = `<label for="` + t.id + `" class="block ` + styles.Label + `">` + t.Label + `</label>`
	}

	hx := ""
	if t.actionID != "" {
		hx = `hx-post="/__glow/action?id=` + t.actionID + `" hx-trigger="change" hx-swap="none" hx-include="input[name='` + t.id + `']"`
	}

	field := ""
	if t.Multiline {
		field = `<textarea id="` + t.id + `" name="` + t.id + `" placeholder="` + t.Placeholder + `" 
			class="` + styles.InputBase + `" ` + hx + `>` + t.Value + `</textarea>`
	} else {
		field = `<input type="text" id="` + t.id + `" name="` + t.id + `" value="` + t.Value + `" placeholder="` + t.Placeholder + `" 
			class="` + styles.InputBase + `" ` + hx + ` />`
	}

	errorHTML := ""
	if t.Error != "" {
		errorHTML = `<p class="` + styles.ErrorText + `">` + t.Error + `</p>`
	}

	return template.HTML(`<div id="` + t.id + `" class="mb-4">` + labelHTML + field + errorHTML + `</div>`)
}

func (t *TextInput) ID() string {
	return t.id
}

// GetValue retorna o valor atual do input
func (t *TextInput) GetValue() string {
	return t.Value
}
