package core

import (
	"github.com/devalexandre/glow/components"
	"github.com/devalexandre/glow/components/styles"
)

var (
	componentsList []components.Component
	sidebarLayout  components.Component
)

func AddComponent(c components.Component) {
	componentsList = append(componentsList, c)
}

func Render(c components.Component) {
	AddComponent(c)
}

func SetSidebar(s components.Component) {
	sidebarLayout = s
}

func renderAllComponents() string {
	main := ""
	for _, c := range componentsList {
		main += string(c.Render())
	}

	side := ""
	if sidebarLayout != nil {
		side = string(sidebarLayout.Render())
	}

	return wrapPage(main, side)
}

func wrapPage(content string, sidebar string) string {
	return `
<!DOCTYPE html>
<html lang="en" class="dark">
<head>
    <meta charset="UTF-8">
    <title>Glow App</title>
    <script src="https://unpkg.com/htmx.org@1.9.10"></script>
    <script src="https://cdn.tailwindcss.com"></script>
    <script>
      tailwind.config = {
        darkMode: 'class',
        theme: {
          extend: {}
        }
      }
    </script>
</head>
<body class="h-screen w-screen overflow-hidden ` + styles.BgDark + ` ` + styles.TextPrimary + `">
  <div class="flex">
    ` + sidebar + `
    <main class="flex-1 p-6 overflow-y-auto">
      <div id="page-content">` + content + `</div>
    </main>
  </div>
</body>
</html>`
}
