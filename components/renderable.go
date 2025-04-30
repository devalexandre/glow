package components

import "net/http"

type Renderable interface {
    Render(w http.ResponseWriter)
}
