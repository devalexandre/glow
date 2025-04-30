package main

import (
	"github.com/devalexandre/glow/components/button"
	"github.com/devalexandre/glow/components/core"
	"github.com/devalexandre/glow/components/input"
	"github.com/devalexandre/glow/components/layouts/sidebar"
	"github.com/devalexandre/glow/components/layouts/vbox"
	"github.com/devalexandre/glow/components/text"
)

func main() {
	txt := text.New("Olá mundo com Glow! ✨")
	input := input.New("Nome", "DevAlexandre")
	btn := button.New("Atualizar")

	core.SetSidebar(sidebar.New(input, btn))

	box := vbox.New(txt)

	btn.OnClick(func() {
		txt.SetContentWithValue(input, "15:04:05")
	})

	core.AddComponent(box)
	core.Run()
}
