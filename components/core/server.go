package core

import (
	"mime/multipart"
	"net/http"
	"sync"

	"github.com/pterm/pterm"
)

var (
	actions = make(map[string]func())
	mu      sync.RWMutex
)

func RegisterAction(id string, fn func()) {
	mu.Lock()
	actions[id] = fn
	mu.Unlock()
}

// Adicione esta função para obter os valores do formulário
func GetFormValues() map[string]string {
	return formValues
}

// Variável global para armazenar os valores do formulário
var formValues = make(map[string]string)

func HandleAction(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	// Parse do formulário para obter os valores
	r.ParseForm()

	// Limpar e atualizar os valores do formulário
	formValues = make(map[string]string)
	for key, values := range r.Form {
		if len(values) > 0 {
			formValues[key] = values[0]
		}
	}

	mu.RLock()
	fn, ok := actions[id]
	mu.RUnlock()

	if !ok {
		http.Error(w, "ação não encontrada", http.StatusNotFound)
		return
	}

	// Execute a ação
	fn()

	// Retorna apenas o conteúdo dinâmico
	main := ""
	for _, c := range componentsList {
		main += string(c.Render())
	}

	// Adiciona cabeçalho para garantir que o HTMX processe corretamente
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(main))
}

var (
	uploadHandlers = make(map[string]func(*multipart.FileHeader))
	uploadMu       sync.RWMutex
)

// RegisterUploadHandler registra um manipulador para processar uploads de arquivos
func RegisterUploadHandler(id string, fn func(*multipart.FileHeader)) {
	uploadMu.Lock()
	uploadHandlers[id] = fn
	uploadMu.Unlock()
}

// HandleUpload processa o upload de arquivos
func HandleUpload(w http.ResponseWriter, r *http.Request) {
	// Verificar se o método é POST
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	// Obter o ID do manipulador
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID não fornecido", http.StatusBadRequest)
		return
	}

	// Verificar se o manipulador existe
	uploadMu.RLock()
	handler, exists := uploadHandlers[id]
	uploadMu.RUnlock()

	if !exists {
		http.Error(w, "Manipulador não encontrado", http.StatusNotFound)
		return
	}

	// Processar o upload do arquivo (máximo 32MB)
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		http.Error(w, "Erro ao processar o formulário: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Obter o arquivo enviado
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Erro ao obter o arquivo: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Chamar o manipulador
	handler(header)

	// Redirecionar para a página principal
	html := renderAllComponents()
	w.Write([]byte(html))
}

func Run() {
	http.HandleFunc("/__glow/action", HandleAction)
	http.HandleFunc("/__glow/upload", HandleUpload)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html := renderAllComponents()
		w.Write([]byte(html))
	})

	paddedBox := pterm.DefaultBox.
		WithLeftPadding(4).
		WithRightPadding(4).
		WithTopPadding(1).
		WithBottomPadding(1)

	title := pterm.LightGreen("Glow App")
	paddedBox.
		WithTitle(title).
		WithTextStyle(pterm.NewStyle(pterm.FgGreen)).
		Println("Servidor iniciado na porta 8080")

	http.ListenAndServe(":8080", nil)
}
