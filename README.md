# Glow ✨

Framework server-driven para criação de interfaces web rápidas e responsivas usando Go, HTMX e TailwindCSS.

## Como usar

```go
page := vbox.New(
    text.New("Olá mundo!"),
    button.New("Clique aqui", "/api/action"),
)
page.Render(w)
```

- Sem HTML manual
- Sem JavaScript
- 100% Golang
- HTMX para requisições dinâmicas
- TailwindCSS para responsividade

## Exemplo

```bash
go run main.go
```
