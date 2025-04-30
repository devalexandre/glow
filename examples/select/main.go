package main

import (
	"github.com/devalexandre/glow/components/button"
	"github.com/devalexandre/glow/components/core"
	"github.com/devalexandre/glow/components/layouts/vbox"
	selectbox "github.com/devalexandre/glow/components/select"
	"github.com/devalexandre/glow/components/text"
)

func main() {
	title := text.New("Exemplo de SelectBox")
	result := text.New("Nenhuma opção selecionada")

	// Criar opções para o select
	options := []selectbox.Option{
		{Value: "go", Label: "Go", Selected: true},
		{Value: "rust", Label: "Rust", Selected: false},
		{Value: "typescript", Label: "TypeScript", Selected: false},
		{Value: "python", Label: "Python", Selected: false},
		{Value: "java", Label: "Java", Selected: false},
	}

	// Criar o select
	selectBox := selectbox.New("Linguagem de Programação", options)
	selectBox.SetPlaceholder("Selecione uma linguagem")

	// Registrar o callback para quando o valor mudar
	selectBox.OnChange(func() {
		result.SetContent("Linguagem selecionada: " + selectBox.GetValue())
	})

	// Botão para mostrar o valor atual
	btn := button.NewPrimary("Mostrar Seleção")
	btn.OnClick(func() {
		result.SetContentWithValue(selectBox, "")
	})

	// Criar layout
	page := vbox.New(
		title,
		selectBox,
		btn,
		result,
	)

	core.AddComponent(page)
	core.Run()
}
