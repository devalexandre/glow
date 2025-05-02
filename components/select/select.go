package selectbox

import (
	"html/template"

	"github.com/devalexandre/glow/components/core"
	"github.com/devalexandre/glow/components/styles"
	"github.com/google/uuid"
)

// Option representa uma opção no select
type Option struct {
	Value    string
	Label    string
	Selected bool
}

// SelectBox representa um componente de seleção dropdown
type SelectBox struct {
	id          string
	Label       string
	Options     []Option
	Value       string
	Placeholder string
	Error       string
	actionID    string
}

// New cria um novo SelectBox
func New(label string, options []Option) *SelectBox {
	s := &SelectBox{
		id:      "select-" + uuid.NewString(),
		Label:   label,
		Options: options,
	}

	// Define o valor inicial como o valor da primeira opção selecionada
	for _, opt := range options {
		if opt.Selected {
			s.Value = opt.Value
			break
		}
	}

	return s
}

// SetPlaceholder define o texto de placeholder
func (s *SelectBox) SetPlaceholder(p string) {
	s.Placeholder = p
}

// SetError define uma mensagem de erro
func (s *SelectBox) SetError(err string) {
	s.Error = err
}

// OnChange registra uma função para ser chamada quando o valor mudar
func (s *SelectBox) OnChange(fn func()) {
	s.actionID = "action-" + s.id

	// Registrar uma ação que atualiza o valor do select antes de chamar a função do usuário
	core.RegisterAction(s.actionID, func() {
		// Atualizar o valor do select com o valor do formulário
		values := core.GetFormValues()
		if val, ok := values[s.id]; ok {
			s.Value = val

			// Atualizar o estado Selected das opções
			for i := range s.Options {
				s.Options[i].Selected = (s.Options[i].Value == val)
			}
		}

		// Chamar a função do usuário
		if fn != nil {
			fn()
		}
	})
}

// Render renderiza o componente
func (s *SelectBox) Render() template.HTML {
	labelHTML := ""
	if s.Label != "" {
		labelHTML = `<label for="` + s.id + `" class="block ` + styles.Label + `">` + s.Label + `</label>`
	}

	hx := ""
	if s.actionID != "" {
		// Usar a mesma configuração do botão para garantir que a página seja atualizada
		hx = `hx-post="/__glow/action?id=` + s.actionID + `" 
			hx-trigger="change" 
			hx-swap="innerHTML" 
			hx-target="#page-content"`
	}

	optionsHTML := ""
	if s.Placeholder != "" {
		selected := ""
		if s.Value == "" {
			selected = " selected"
		}
		optionsHTML += `<option value="" disabled` + selected + `>` + s.Placeholder + `</option>`
	}

	for _, opt := range s.Options {
		selected := ""
		if opt.Selected {
			selected = " selected"
		}
		optionsHTML += `<option value="` + opt.Value + `"` + selected + `>` + opt.Label + `</option>`
	}

	errorHTML := ""
	if s.Error != "" {
		errorHTML = `<p class="` + styles.ErrorText + `">` + s.Error + `</p>`
	}

	return template.HTML(`<div class="mb-4">
        ` + labelHTML + `
        <select id="` + s.id + `" name="` + s.id + `" class="` + styles.SelectBase + `" ` + hx + `>
            ` + optionsHTML + `
        </select>
        ` + errorHTML + `
    </div>`)
}

// ID retorna o ID do componente
func (s *SelectBox) ID() string {
	return s.id
}

// GetValue retorna o valor atual do select
func (s *SelectBox) GetValue() string {
	return s.Value
}

// SetValue define o valor do select
func (s *SelectBox) SetValue(value string) {
	s.Value = value

	// Atualizar o estado Selected das opções
	for i := range s.Options {
		s.Options[i].Selected = (s.Options[i].Value == value)
	}
}

// AddOption adiciona uma nova opção ao select
func (s *SelectBox) AddOption(value, label string, selected bool) {
	option := Option{
		Value:    value,
		Label:    label,
		Selected: selected,
	}

	// Se a nova opção estiver selecionada, atualiza o valor do select
	if selected {
		s.Value = value

		// Desmarcar outras opções
		for i := range s.Options {
			s.Options[i].Selected = false
		}
	}

	s.Options = append(s.Options, option)
}
