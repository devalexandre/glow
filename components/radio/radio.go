package radio

import (
	"html/template"

	"github.com/devalexandre/glow/components/core"
	"github.com/devalexandre/glow/components/styles"
	"github.com/google/uuid"
)

// RadioOption representa uma opção no grupo de radio buttons
type RadioOption struct {
	Value    string
	Label    string
	Selected bool
}

// RadioGroup representa um grupo de radio buttons
type RadioGroup struct {
	id         string
	Name       string
	Label      string
	Options    []RadioOption
	Value      string
	Error      string
	actionID   string
	horizontal bool
}

// New cria um novo RadioGroup
func New(name, label string, options []RadioOption) *RadioGroup {
	r := &RadioGroup{
		id:      "radio-" + uuid.NewString(),
		Name:    name,
		Label:   label,
		Options: options,
	}

	// Define o valor inicial como o valor da primeira opção selecionada
	for _, opt := range options {
		if opt.Selected {
			r.Value = opt.Value
			break
		}
	}

	return r
}

// SetError define uma mensagem de erro
func (r *RadioGroup) SetError(err string) {
	r.Error = err
}

// SetHorizontal define se os radio buttons devem ser exibidos horizontalmente
func (r *RadioGroup) SetHorizontal(horizontal bool) {
	r.horizontal = horizontal
}

// OnChange registra uma função para ser chamada quando o valor mudar
func (r *RadioGroup) OnChange(fn func()) {
	r.actionID = "action-" + r.id

	// Registrar uma ação que atualiza o valor do radio antes de chamar a função do usuário
	core.RegisterAction(r.actionID, func() {
		// Atualizar o valor do radio com o valor do formulário
		values := core.GetFormValues()
		if val, ok := values[r.Name]; ok {
			r.Value = val

			// Atualizar o estado Selected das opções
			for i := range r.Options {
				r.Options[i].Selected = (r.Options[i].Value == val)
			}
		}

		// Chamar a função do usuário
		if fn != nil {
			fn()
		}
	})
}

// Render renderiza o componente
func (r *RadioGroup) Render() template.HTML {
	labelHTML := ""
	if r.Label != "" {
		labelHTML = `<div class="` + styles.Label + ` mb-2">` + r.Label + `</div>`
	}

	hx := ""
	if r.actionID != "" {
		hx = `hx-post="/__glow/action?id=` + r.actionID + `" 
            hx-trigger="change" 
            hx-swap="innerHTML" 
            hx-target="#page-content"`
	}

	containerClass := "space-y-2" // Vertical por padrão
	if r.horizontal {
		containerClass = "flex space-x-4" // Horizontal
	}

	optionsHTML := ""
	for _, opt := range r.Options {
		checked := ""
		if opt.Selected {
			checked = " checked"
		}

		optionID := r.id + "-" + opt.Value

		optionsHTML += `<label class="flex items-center">
            <input type="radio" id="` + optionID + `" name="` + r.Name + `" value="` + opt.Value + `" 
                class="` + styles.RadioBase + `"` + checked + ` ` + hx + `>
            <span class="ml-2 ` + styles.Label + `">` + opt.Label + `</span>
        </label>`
	}

	errorHTML := ""
	if r.Error != "" {
		errorHTML = `<p class="` + styles.ErrorText + `">` + r.Error + `</p>`
	}

	return template.HTML(`<div id="` + r.id + `" class="mb-4">
        ` + labelHTML + `
        <div class="` + containerClass + `">
            ` + optionsHTML + `
        </div>
        ` + errorHTML + `
    </div>`)
}

// ID retorna o ID do componente
func (r *RadioGroup) ID() string {
	return r.id
}

// GetValue retorna o valor atual do radio group
func (r *RadioGroup) GetValue() string {
	return r.Value
}

// SetValue define o valor do radio group
func (r *RadioGroup) SetValue(value string) {
	r.Value = value

	// Atualizar o estado Selected das opções
	for i := range r.Options {
		r.Options[i].Selected = (r.Options[i].Value == value)
	}
}

// GetSelectedOption retorna a opção selecionada
func (r *RadioGroup) GetSelectedOption() *RadioOption {
	for i, opt := range r.Options {
		if opt.Selected {
			return &r.Options[i]
		}
	}
	return nil
}
