package main

import (
	"strings"

	"github.com/devalexandre/glow/components/button"
	"github.com/devalexandre/glow/components/core"
	"github.com/devalexandre/glow/components/layouts/vbox"
	"github.com/devalexandre/glow/components/multiselect"
	"github.com/devalexandre/glow/components/text"
)

func main() {
	title := text.New("Exemplo de MultiSelect")
	result := text.New("Nenhuma opção selecionada")

	// Criar opções para o multiselect
	options := []multiselect.Option{
		{Value: "go", Label: "Go", Selected: true},
		{Value: "rust", Label: "Rust", Selected: false},
		{Value: "typescript", Label: "TypeScript", Selected: true},
		{Value: "python", Label: "Python", Selected: false},
		{Value: "java", Label: "Java", Selected: false},
		{Value: "csharp", Label: "C#", Selected: false},
		{Value: "cpp", Label: "C++", Selected: false},
		{Value: "ruby", Label: "Ruby", Selected: false},
		{Value: "php", Label: "PHP", Selected: false},
		{Value: "swift", Label: "Swift", Selected: false},
	}

	// Criar o multiselect no modo tag (padrão)
	multiSelect := multiselect.New("Linguagens de Programação", options)
	multiSelect.SetPlaceholder("Selecione uma ou mais linguagens")

	// Criar um segundo multiselect no modo dropdown
	dropdownOptions := []multiselect.Option{
		{Value: "red", Label: "Red", Selected: true},
		{Value: "blue", Label: "Blue", Selected: false},
		{Value: "green", Label: "Green", Selected: false},
		{Value: "yellow", Label: "Yellow", Selected: true},
		{Value: "purple", Label: "Purple", Selected: false},
	}
	dropdownMultiSelect := multiselect.New("Cores (Dropdown)", dropdownOptions)
	dropdownMultiSelect.SetPlaceholder("Selecione uma ou mais cores")
	dropdownMultiSelect.SetAsDropdown(true) // Usar o modo dropdown tradicional

	// Registrar o callback para quando os valores mudarem
	multiSelect.OnChange(func() {
		values := multiSelect.GetValues()
		if len(values) > 0 {
			result.SetContent("Linguagens selecionadas: " + strings.Join(values, ", "))
		} else {
			result.SetContent("Nenhuma linguagem selecionada")
		}
	})

	dropdownMultiSelect.OnChange(func() {
		values := dropdownMultiSelect.GetValues()
		if len(values) > 0 {
			result.SetContent("Cores selecionadas: " + strings.Join(values, ", "))
		} else {
			result.SetContent("Nenhuma cor selecionada")
		}
	})

	// Botão para mostrar os valores atuais
	btn := button.NewPrimary("Mostrar Seleção")
	btn.OnClick(func() {
		selected := multiSelect.GetSelectedOptions()
		if len(selected) > 0 {
			labels := []string{}
			for _, opt := range selected {
				labels = append(labels, opt.Label)
			}
			result.SetContent("Linguagens selecionadas: " + strings.Join(labels, ", "))
		} else {
			result.SetContent("Nenhuma linguagem selecionada")
		}
	})

	// Botão para selecionar todas as opções
	btnSelectAll := button.NewSuccess("Selecionar Todas")
	btnSelectAll.OnClick(func() {
		values := []string{}
		for _, opt := range options {
			values = append(values, opt.Value)
		}
		multiSelect.SetValues(values)
		result.SetContent("Todas as linguagens selecionadas")
	})

	// Botão para limpar a seleção
	btnClear := button.NewDanger("Limpar Seleção")
	btnClear.OnClick(func() {
		multiSelect.SetValues([]string{})
		result.SetContent("Nenhuma linguagem selecionada")
	})

	// Criar layout
	page := vbox.New(
		title,
		multiSelect,
		dropdownMultiSelect,
		btn,
		btnSelectAll,
		btnClear,
		result,
	)

	core.AddComponent(page)
	core.Run()
}
