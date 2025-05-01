package main

import (
	"path/filepath"

	"github.com/devalexandre/glow/components/button"
	"github.com/devalexandre/glow/components/core"
	"github.com/devalexandre/glow/components/layouts/vbox"
	"github.com/devalexandre/glow/components/text"
	"github.com/devalexandre/glow/components/upload"
)

func main() {
	title := text.New("Exemplo de Upload de Arquivos")
	result := text.New("Nenhum arquivo enviado ainda")

	// Criar o componente de upload
	uploader := upload.New("Envie um arquivo", "./uploads")

	// Configurar para aceitar apenas imagens
	uploader.SetAcceptTypes(".jpg,.jpeg,.png,.gif")

	// Registrar o callback para quando o upload for concluído
	uploader.OnComplete(func(filePath string) {
		if filePath != "" {
			fileName := filepath.Base(filePath)
			result.SetContent("Arquivo enviado com sucesso: " + fileName)
		} else {
			result.SetContent("Falha ao enviar o arquivo")
		}
	})

	// Botão para limpar o resultado
	btn := button.NewDanger("Limpar")
	btn.OnClick(func() {
		result.SetContent("Nenhum arquivo enviado ainda")
	})

	// Criar layout
	page := vbox.New(
		title,
		uploader,
		result,
		btn,
	)

	core.AddComponent(page)
	core.Run()
}
