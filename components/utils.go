package components

// StringWriter permite capturar a saída do `template.Execute`
type StringWriter struct {
	Str *string
}

func (w *StringWriter) Write(p []byte) (int, error) {
	*w.Str += string(p)
	return len(p), nil
}
