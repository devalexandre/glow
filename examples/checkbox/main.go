package main

import (
	"github.com/devalexandre/glow/components/button"
	"github.com/devalexandre/glow/components/checkbox"
	"github.com/devalexandre/glow/components/core"
	"github.com/devalexandre/glow/components/layouts/vbox"
	"github.com/devalexandre/glow/components/text"
)

func main() {
	title := text.New("Exemplo de Checkbox")
	result := text.New("Status: Não marcado")

	// Criar o checkbox
	cb := checkbox.New("Aceito os termos e condições", false)

	// Registrar o callback para quando o valor mudar
	cb.OnChange(func() {
		if cb.IsChecked() {
			result.SetContent("Status: Marcado")
		} else {
			result.SetContent("Status: Não marcado")
		}
	})

	// Botão para inverter o estado do checkbox
	btn := button.NewPrimary("Inverter Estado")
	btn.OnClick(func() {
		cb.SetChecked(!cb.IsChecked())
		if cb.IsChecked() {
			result.SetContent("Status: Marcado")
		} else {
			result.SetContent("Status: Não marcado")
		}
	})

	// Criar layout
	page := vbox.New(
		title,
		cb,
		btn,
		result,
	)

	core.AddComponent(page)
	core.Run()
}
