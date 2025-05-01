package multiselect

import (
	"html/template"
	"strings"

	"github.com/devalexandre/glow/components/core"
	"github.com/devalexandre/glow/components/styles"
	"github.com/google/uuid"
)

// Option representa uma opção no multiselect
type Option struct {
	Value    string
	Label    string
	Selected bool
}

// MultiSelect representa um componente de seleção múltipla
type MultiSelect struct {
	id          string
	Label       string
	Options     []Option
	Values      []string
	Placeholder string
	Error       string
	actionID    string
	removeID    string
	asDropdown  bool
}

// New cria um novo MultiSelect
func New(label string, options []Option) *MultiSelect {
	m := &MultiSelect{
		id:      "multiselect-" + uuid.NewString(),
		Label:   label,
		Options: options,
		Values:  []string{},
	}
	// Define os valores iniciais com base nas opções selecionadas
	for _, opt := range options {
		if opt.Selected {
			m.Values = append(m.Values, opt.Value)
		}
	}
	return m
}

// SetPlaceholder define o texto de placeholder
func (m *MultiSelect) SetPlaceholder(p string) {
	m.Placeholder = p
}

// SetError define uma mensagem de erro
func (m *MultiSelect) SetError(err string) {
	m.Error = err
}

// SetAsDropdown define se o multiselect deve ser renderizado como um dropdown tradicional
func (m *MultiSelect) SetAsDropdown(asDropdown bool) {
	m.asDropdown = asDropdown
}

// OnChange registra uma função para ser chamada quando os valores mudarem
func (m *MultiSelect) OnChange(fn func()) {
	m.actionID = "action-" + m.id
	m.removeID = "remove-" + m.id

	// Registrar uma ação que atualiza os valores do multiselect antes de chamar a função do usuário
	core.RegisterAction(m.actionID, func() {
		// Atualizar os valores do multiselect com os valores do formulário
		values := core.GetFormValues()
		if m.asDropdown {
			// Limpar os valores atuais
			m.Values = []string{}
			// O formulário envia os valores selecionados como um array com o mesmo nome
			// Por exemplo: multiselect-123=valor1&multiselect-123=valor2
			for key, val := range values {
				if strings.HasPrefix(key, m.id) {
					m.Values = append(m.Values, val)
				}
			}
		} else {
			// Modo tag select
			if val, ok := values[m.id]; ok && val != "" {
				// Verificar se o valor já está selecionado
				alreadySelected := false
				for _, v := range m.Values {
					if v == val {
						alreadySelected = true
						break
					}
				}
				// Adicionar o valor se não estiver já selecionado
				if !alreadySelected {
					m.Values = append(m.Values, val)
				}
			}
		}
		// Atualizar o estado Selected das opções
		for i := range m.Options {
			m.Options[i].Selected = false
			for _, val := range m.Values {
				if m.Options[i].Value == val {
					m.Options[i].Selected = true
					break
				}
			}
		}
		// Chamar a função do usuário
		if fn != nil {
			fn()
		}
	})

	// Registrar uma ação para remover uma tag (apenas para o modo tag select)
	core.RegisterAction(m.removeID, func() {
		// Obter o valor a ser removido
		values := core.GetFormValues()
		if val, ok := values["value"]; ok {
			// Remover o valor da lista
			newValues := []string{}
			for _, v := range m.Values {
				if v != val {
					newValues = append(newValues, v)
				}
			}
			m.Values = newValues
			// Atualizar o estado Selected das opções
			for i := range m.Options {
				if m.Options[i].Value == val {
					m.Options[i].Selected = false
					break
				}
			}
		}
		// Chamar a função do usuário
		if fn != nil {
			fn()
		}
	})
}

// Render renderiza o componente
func (m *MultiSelect) Render() template.HTML {
	labelHTML := ""
	if m.Label != "" {
		labelHTML = `<label for="` + m.id + `" class="block ` + styles.Label + `">` + m.Label + `</label>`
	}
	errorHTML := ""
	if m.Error != "" {
		errorHTML = `<p class="` + styles.ErrorText + `">` + m.Error + `</p>`
	}

	// Renderizar como dropdown tradicional
	if m.asDropdown {
		hx := ""
		if m.actionID != "" {
			hx = `hx-post="/__glow/action?id=` + m.actionID + `" 
                hx-trigger="change" 
                hx-swap="innerHTML" 
                hx-target="#page-content"`
		}
		optionsHTML := ""
		if m.Placeholder != "" {
			selected := ""
			if len(m.Values) == 0 {
				selected = " selected"
			}
			optionsHTML += `<option value="" disabled` + selected + `>` + m.Placeholder + `</option>`
		}
		for _, opt := range m.Options {
			selected := ""
			if opt.Selected {
				selected = " selected"
			}
			optionsHTML += `<option value="` + opt.Value + `"` + selected + `>` + opt.Label + `</option>`
		}
		return template.HTML(`<div class="mb-4">
            ` + labelHTML + `
            <select id="` + m.id + `" name="` + m.id + `" multiple class="` + styles.MultiSelectBase + `" ` + hx + `>
                ` + optionsHTML + `
            </select>
            ` + errorHTML + `
        </div>`)
	}

	// Renderizar como tag select integrado
	// Criar os chips para os valores selecionados
	tagsHTML := ""
	for _, val := range m.Values {
		// Encontrar o label correspondente ao valor
		label := val
		for _, opt := range m.Options {
			if opt.Value == val {
				label = opt.Label
				break
			}
		}
		// Criar o chip com botão de remover
		tagsHTML += `<span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-dracula-purple text-dracula-foreground mr-2 mb-1">
            ` + label + `
            <button type="button" class="ml-1.5 inline-flex items-center justify-center h-4 w-4 rounded-full text-dracula-foreground hover:bg-dracula-pink focus:outline-none tag-button"
                hx-post="/__glow/action?id=` + m.removeID + `" 
                hx-trigger="click" 
                hx-swap="innerHTML" 
                hx-target="#page-content"
                hx-vals='{"value":"` + val + `"}'
                onclick="event.stopPropagation()">
                <span class="sr-only">Remover</span>
                <svg class="h-3 w-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                </svg>
            </button>
        </span>`
	}

	// Criar o dropdown para selecionar novas opções
	optionsHTML := ""
	if m.Placeholder != "" {
		optionsHTML += `<option value="" selected disabled>` + m.Placeholder + `</option>`
	}
	for _, opt := range m.Options {
		// Não mostrar opções já selecionadas no dropdown
		if !opt.Selected {
			optionsHTML += `<option value="` + opt.Value + `">` + opt.Label + `</option>`
		}
	}

	// No modo tag, adicionamos o evento change diretamente ao select para adicionar automaticamente
	selectHx := ""
	if m.actionID != "" {
		selectHx = `hx-post="/__glow/action?id=` + m.actionID + `" 
            hx-trigger="change" 
            hx-swap="innerHTML" 
            hx-target="#page-content"`
	}

	// Placeholder para mostrar quando não há tags selecionadas
	placeholderHTML := ""
	if len(m.Values) == 0 && m.Placeholder != "" {
		placeholderHTML = `<span class="text-dracula-comment">` + m.Placeholder + `</span>`
	}

	// Estilos ajustados para evitar sobreposição e permitir interação
	styles := `
    <style>
        #` + m.id + `-container {
            position: relative;
            display: flex;
            align-items: center;
            justify-content: space-between; /* Distribui os elementos */
            width: 100%;
            border: 1px solid #6272a4;
            border-radius: 0.375rem;
            background-color: #44475a;
            min-height: 42px;
            cursor: pointer;
            overflow: hidden; /* Evita que tags ultrapassem o contêiner */
        }
        #` + m.id + `-tags {
            display: flex;
            flex-wrap: wrap;
            gap: 4px;
            padding: 0.5rem;
            flex: 0 1 auto; /* Não cresce, mas ocupa espaço necessário */
            max-width: 70%; /* Deixa espaço para o dropdown */
            z-index: 20;
        }
        #` + m.id + `-select {
            position: absolute;
            top: 0;
            left: 0;
            width: 100%;
            height: 100%;
            opacity: 0;
            z-index: 10;
            cursor: pointer;
            pointer-events: auto; /* Garante que o select receba cliques */
        }
        #` + m.id + `-container .tag-button {
            z-index: 30;
        }
    </style>
    `

	// JavaScript para melhorar a interação
	javascript := `
    <script>
        document.addEventListener('DOMContentLoaded', function() {
            const container = document.getElementById('` + m.id + `-container');
            const select = document.getElementById('` + m.id + `-select');
            const buttons = container.querySelectorAll('.tag-button');

            // Abrir dropdown apenas ao clicar em áreas vazias do container
            container.addEventListener('click', function(e) {
                if (!e.target.closest('.tag-button')) {
                    select.focus();
                    select.click();
                }
            });

            // Evitar propagação de clique nos botões de remoção
            buttons.forEach(button => {
                button.addEventListener('click', function(e) {
                    e.stopPropagation();
                });
            });
        });
    </script>
    `

	// Retornar o componente completo
	return template.HTML(styles + javascript + `<div class="mb-4">
        ` + labelHTML + `
        <div id="` + m.id + `-container">
            <div id="` + m.id + `-tags">
                ` + tagsHTML + `
                ` + placeholderHTML + `
            </div>
            <select id="` + m.id + `-select" name="` + m.id + `" ` + selectHx + `>
                ` + optionsHTML + `
            </select>
        </div>
        ` + errorHTML + `
    </div>`)
}

// ID retorna o ID do componente
func (m *MultiSelect) ID() string {
	return m.id
}

// GetValues retorna os valores atuais do multiselect
func (m *MultiSelect) GetValues() []string {
	return m.Values
}

// GetValuesString retorna os valores atuais do multiselect como uma string separada por vírgulas
func (m *MultiSelect) GetValuesString() string {
	return strings.Join(m.Values, ", ")
}

// SetValues define os valores do multiselect
func (m *MultiSelect) SetValues(values []string) {
	m.Values = values
	// Atualizar o estado Selected das opções
	for i := range m.Options {
		m.Options[i].Selected = false
		for _, val := range values {
			if m.Options[i].Value == val {
				m.Options[i].Selected = true
				break
			}
		}
	}
}

// GetSelectedOptions retorna as opções selecionadas
func (m *MultiSelect) GetSelectedOptions() []Option {
	var selected []Option
	for _, opt := range m.Options {
		if opt.Selected {
			selected = append(selected, opt)
		}
	}
	return selected
}

// AddOption adiciona uma nova opção ao multiselect
func (m *MultiSelect) AddOption(value, label string, selected bool) {
	option := Option{
		Value:    value,
		Label:    label,
		Selected: selected,
	}
	// Se a nova opção estiver selecionada, adiciona ao array de valores
	if selected {
		m.Values = append(m.Values, value)
	}
	m.Options = append(m.Options, option)
}
