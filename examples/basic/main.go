package main

import (
	"time"

	"github.com/devalexandre/glow/components"
	"github.com/devalexandre/glow/components/button"
	"github.com/devalexandre/glow/components/core"
	"github.com/devalexandre/glow/components/input"
	"github.com/devalexandre/glow/components/layouts/vbox"
	"github.com/devalexandre/glow/components/text"
)

func main() {

	txt := text.New("Olá mundo com Glow! ✨")
	input := input.New("Nome", "DevAlexandre")

	btn := button.NewTargetFunc("Atualizar", txt, func() components.Component {
		txt.SetContent(input.Value + " - " + time.Now().Format("15:04:05"))
		return txt
	})

	page := vbox.New(txt, input, btn)

	core.AddComponent(page)
	core.Run()
}
