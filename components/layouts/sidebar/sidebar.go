package sidebar

import (
	"html/template"

	"github.com/devalexandre/glow/components"
	"github.com/devalexandre/glow/components/styles"
	"github.com/google/uuid"
)

type Sidebar struct {
	id       string
	Children []components.Component
}

func New(children ...components.Component) *Sidebar {
	return &Sidebar{
		id:       "sidebar-" + uuid.NewString(),
		Children: children,
	}
}

func (s *Sidebar) Render() template.HTML {
	content := `<aside id="` + s.id + `" class="` + styles.SidebarBase + `">`
	for _, c := range s.Children {
		content += string(c.Render())
	}
	content += `</aside>`
	return template.HTML(content)
}

func (s *Sidebar) ID() string {
	return s.id
}
