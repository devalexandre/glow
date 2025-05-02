package main

import (
	"github.com/devalexandre/glow/components/button"
	"github.com/devalexandre/glow/components/core"
	"github.com/devalexandre/glow/components/layouts/vbox"
	"github.com/devalexandre/glow/components/radio"
	"github.com/devalexandre/glow/components/text"
)

func main() {
	title := text.New("Exemplo de Radio Button")
	result := text.New("Nenhuma opção selecionada")

	// Criar opções para o radio group
	options := []radio.RadioOption{
		{Value: "go", Label: "Go", Selected: true},
		{Value: "rust", Label: "Rust", Selected: false},
		{Value: "typescript", Label: "TypeScript", Selected: false},
		{Value: "python", Label: "Python", Selected: false},
		{Value: "java", Label: "Java", Selected: false},
	}

	// Criar o radio group
	radioGroup := radio.New("language", "Linguagem de Programação", options)

	// Versão horizontal
	horizontalOptions := []radio.RadioOption{
		{Value: "light", Label: "Claro", Selected: true},
		{Value: "dark", Label: "Escuro", Selected: false},
		{Value: "system", Label: "Sistema", Selected: false},
	}
	horizontalRadio := radio.New("theme", "Tema", horizontalOptions)
	horizontalRadio.SetHorizontal(true)

	// Registrar o callback para quando o valor mudar
	radioGroup.OnChange(func() {
		result.SetContent("Linguagem selecionada: " + radioGroup.GetValue())
	})

	horizontalRadio.OnChange(func() {
		result.SetContent("Tema selecionado: " + horizontalRadio.GetValue())
	})

	// Botão para mostrar o valor atual
	btn := button.NewPrimary("Mostrar Seleção")
	btn.OnClick(func() {
		selected := radioGroup.GetSelectedOption()
		if selected != nil {
			result.SetContent("Linguagem: " + selected.Label + " (" + selected.Value + "), Tema: " + horizontalRadio.GetValue())
		} else {
			result.SetContent("Nenhuma linguagem selecionada")
		}
	})

	// Criar layout
	page := vbox.New(
		title,
		radioGroup,
		horizontalRadio,
		btn,
		result,
	)

	core.AddComponent(page)
	core.Run()
}
