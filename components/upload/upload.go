package upload

import (
	"html/template"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/devalexandre/glow/components/core"
	"github.com/devalexandre/glow/components/styles"
	"github.com/google/uuid"
)

// Upload representa um componente de upload de arquivos
type Upload struct {
	id           string
	Label        string
	UploadPath   string
	AcceptTypes  string
	Multiple     bool
	Error        string
	SuccessMsg   string
	UploadedFile string
	onComplete   func(string)
}

// New cria um novo componente de upload
func New(label, uploadPath string) *Upload {
	return &Upload{
		id:          "upload-" + uuid.NewString(),
		Label:       label,
		UploadPath:  uploadPath,
		AcceptTypes: "*", // Aceita todos os tipos de arquivo por padrão
	}
}

// SetAcceptTypes define os tipos de arquivo aceitos (ex: ".jpg,.png,.pdf")
func (u *Upload) SetAcceptTypes(types string) {
	u.AcceptTypes = types
}

// SetMultiple define se o upload aceita múltiplos arquivos
func (u *Upload) SetMultiple(multiple bool) {
	u.Multiple = multiple
}

// SetError define uma mensagem de erro
func (u *Upload) SetError(err string) {
	u.Error = err
}

// SetSuccessMsg define uma mensagem de sucesso
func (u *Upload) SetSuccessMsg(msg string) {
	u.SuccessMsg = msg
}

// OnComplete registra uma função para ser chamada quando o upload for concluído
func (u *Upload) OnComplete(fn func(string)) {
	u.onComplete = fn

	// Registrar a ação para processar o upload
	core.RegisterUploadHandler(u.id, func(header *multipart.FileHeader) {
		// Salvar o arquivo no caminho especificado
		filePath, err := u.saveFile(header)
		if err != nil {
			u.Error = "Erro ao salvar o arquivo: " + err.Error()
			u.SuccessMsg = ""
			return
		}

		u.UploadedFile = filePath
		u.Error = ""
		if u.SuccessMsg == "" {
			u.SuccessMsg = "Arquivo enviado com sucesso!"
		}

		// Chamar a função do usuário
		if fn != nil {
			fn(filePath)
		}
	})
}

// saveFile salva o arquivo enviado no caminho especificado
func (u *Upload) saveFile(file *multipart.FileHeader) (string, error) {
	// Criar o diretório de upload se não existir
	if err := os.MkdirAll(u.UploadPath, 0755); err != nil {
		return "", err
	}

	// Abrir o arquivo enviado
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Gerar um nome único para o arquivo
	filename := filepath.Base(file.Filename)
	ext := filepath.Ext(filename)
	name := strings.TrimSuffix(filename, ext)
	uniqueFilename := name + "-" + uuid.NewString() + ext

	// Caminho completo do arquivo
	filePath := filepath.Join(u.UploadPath, uniqueFilename)

	// Criar o arquivo de destino
	dst, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Copiar o conteúdo do arquivo enviado para o arquivo de destino
	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}

	return filePath, nil
}

// Render renderiza o componente
func (u *Upload) Render() template.HTML {
	labelHTML := ""
	if u.Label != "" {
		labelHTML = `<label for="` + u.id + `" class="block ` + styles.Label + `">` + u.Label + `</label>`
	}

	multipleAttr := ""
	if u.Multiple {
		multipleAttr = " multiple"
	}

	errorHTML := ""
	if u.Error != "" {
		errorHTML = `<p class="` + styles.ErrorText + `">` + u.Error + `</p>`
	}

	successHTML := ""
	if u.SuccessMsg != "" {
		successHTML = `<p class="` + styles.SuccessText + `">` + u.SuccessMsg + `</p>`
	}

	fileInfoHTML := ""
	if u.UploadedFile != "" {
		fileInfoHTML = `<p class="text-sm text-dracula-comment mt-1">Arquivo: ` + filepath.Base(u.UploadedFile) + `</p>`
	}

	// Estilos para o componente de upload
	styles := `
    <style>
        #` + u.id + `-container {
            position: relative;
            width: 100%;
            margin-bottom: 1rem;
        }
        #` + u.id + `-input {
            opacity: 0;
            position: absolute;
            top: 0;
            left: 0;
            width: 100%;
            height: 100%;
            cursor: pointer;
            z-index: 10;
        }
        #` + u.id + `-button {
            display: flex;
            align-items: center;
            justify-content: center;
            padding: 0.75rem 1rem;
            background-color: #6272a4;
            color: #f8f8f2;
            border-radius: 0.375rem;
            font-size: 0.875rem;
            font-weight: 500;
            width: 100%;
            border: 1px dashed #bd93f9;
            transition: all 0.2s;
        }
        #` + u.id + `-button:hover {
            background-color: #44475a;
        }
        #` + u.id + `-button svg {
            margin-right: 0.5rem;
        }
        .file-selected #` + u.id + `-button {
            background-color: #44475a;
            border-color: #50fa7b;
        }
    </style>
    `

	// JavaScript para melhorar a interação
	javascript := `
    <script>
        document.addEventListener('DOMContentLoaded', function() {
            const input = document.getElementById('` + u.id + `-input');
            const button = document.getElementById('` + u.id + `-button');
            const container = document.getElementById('` + u.id + `-container');
            const fileNameSpan = document.getElementById('` + u.id + `-filename');
            
            input.addEventListener('change', function() {
                if (input.files.length > 0) {
                    container.classList.add('file-selected');
                    const fileCount = input.files.length;
                    if (fileCount === 1) {
                        fileNameSpan.textContent = input.files[0].name;
                    } else {
                        fileNameSpan.textContent = fileCount + ' arquivos selecionados';
                    }
                } else {
                    container.classList.remove('file-selected');
                    fileNameSpan.textContent = 'Selecione um arquivo';
                }
            });
        });
    </script>
    `

	return template.HTML(styles + javascript + `<div class="mb-4">
        ` + labelHTML + `
        <div id="` + u.id + `-container" class="` + func() string {
		if u.UploadedFile != "" {
			return "file-selected"
		}
		return ""
	}() + `">
            <div id="` + u.id + `-button">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                    <path fill-rule="evenodd" d="M3 17a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM6.293 6.707a1 1 0 010-1.414l3-3a1 1 0 011.414 0l3 3a1 1 0 01-1.414 1.414L11 5.414V13a1 1 0 11-2 0V5.414L7.707 6.707a1 1 0 01-1.414 0z" clip-rule="evenodd" />
                </svg>
                <span id="` + u.id + `-filename">` + func() string {
		if u.UploadedFile != "" {
			return filepath.Base(u.UploadedFile)
		}
		return "Selecione um arquivo"
	}() + `</span>
            </div>
            <form id="` + u.id + `-form" enctype="multipart/form-data" method="post" action="/__glow/upload?id=` + u.id + `">
                <input type="file" id="` + u.id + `-input" name="file" accept="` + u.AcceptTypes + `"` + multipleAttr + `
                    onchange="this.form.submit();" />
            </form>
        </div>
        ` + errorHTML + `
        ` + successHTML + `
        ` + fileInfoHTML + `
    </div>`)
}

// ID retorna o ID do componente
func (u *Upload) ID() string {
	return u.id
}

// GetUploadedFile retorna o caminho do arquivo enviado
func (u *Upload) GetUploadedFile() string {
	return u.UploadedFile
}
