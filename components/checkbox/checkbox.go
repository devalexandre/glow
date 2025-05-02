package checkbox

import (
	"html/template"

	"github.com/devalexandre/glow/components/core"
	"github.com/devalexandre/glow/components/styles"
	"github.com/google/uuid"
)

// Checkbox representa um componente de caixa de seleção
type Checkbox struct {
	id       string
	Label    string
	Checked  bool
	Error    string
	actionID string
}

// New cria um novo Checkbox
func New(label string, checked bool) *Checkbox {
	return &Checkbox{
		id:      "checkbox-" + uuid.NewString(),
		Label:   label,
		Checked: checked,
	}
}

// SetError define uma mensagem de erro
func (c *Checkbox) SetError(err string) {
	c.Error = err
}

// OnChange registra uma função para ser chamada quando o valor mudar
func (c *Checkbox) OnChange(fn func()) {
	c.actionID = "action-" + c.id

	// Registrar uma ação que atualiza o valor do checkbox antes de chamar a função do usuário
	core.RegisterAction(c.actionID, func() {
		// Atualizar o valor do checkbox com o valor do formulário
		values := core.GetFormValues()
		if _, ok := values[c.id]; ok {
			c.Checked = true
		} else {
			c.Checked = false
		}

		// Chamar a função do usuário
		if fn != nil {
			fn()
		}
	})
}

// Render renderiza o componente
func (c *Checkbox) Render() template.HTML {
	checked := ""
	if c.Checked {
		checked = " checked"
	}

	hx := ""
	if c.actionID != "" {
		hx = `hx-post="/__glow/action?id=` + c.actionID + `" 
            hx-trigger="change" 
            hx-swap="innerHTML" 
            hx-target="#page-content"`
	}

	errorHTML := ""
	if c.Error != "" {
		errorHTML = `<p class="` + styles.ErrorText + `">` + c.Error + `</p>`
	}

	return template.HTML(`<div class="mb-4">
        <label class="flex items-center">
            <input type="checkbox" id="` + c.id + `" name="` + c.id + `" class="` + styles.CheckboxBase + `"` + checked + ` ` + hx + `>
            <span class="ml-2 ` + styles.Label + `">` + c.Label + `</span>
        </label>
        ` + errorHTML + `
    </div>`)
}

// ID retorna o ID do componente
func (c *Checkbox) ID() string {
	return c.id
}

// GetValue retorna o valor atual do checkbox como string ("true" ou "false")
func (c *Checkbox) GetValue() string {
	if c.Checked {
		return "true"
	}
	return "false"
}

// IsChecked retorna se o checkbox está marcado
func (c *Checkbox) IsChecked() bool {
	return c.Checked
}

// SetChecked define se o checkbox está marcado
func (c *Checkbox) SetChecked(checked bool) {
	c.Checked = checked
}
