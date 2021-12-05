package humanerrors

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

var tmpl = template.Must(template.New("").Parse(strings.TrimSpace(`
Oh no! {{ .Message }}
{{- if .Advice }}

To fix this, you can try:
{{- range .Advice }}
 - {{ . }}
{{- end -}}
{{- with .Cause }}

This was caused by:
{{ .Error }}
{{ end -}}
{{ end -}}
`)))

type HumaneError interface {
	Error() string
	Message() string
	Advice() []string
	Cause() error
}

type humaneError struct {
	message string
	advice  []string
	cause   error
}

func New(message string, advice ...string) HumaneError {
	return &humaneError{message, advice, nil}
}

func NewWithCause(cause error, message string, advice ...string) HumaneError {
	return &humaneError{message, advice, cause}
}

func (e *humaneError) Error() string {
	b := bytes.NewBufferString("")
	err := tmpl.Execute(b, e)
	if err != nil {
		return fmt.Sprintf("%s (error building nice message: %s)", e.message, err.Error())
	}

	return b.String()
}

func (e *humaneError) Message() string {
	return e.message
}

func (e *humaneError) Advice() []string {
	return e.advice
}

func (e *humaneError) Cause() error {
	return e.cause
}
