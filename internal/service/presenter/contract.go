package presenter

import "io"

type Renderer interface {
	ExecuteTemplate(io.Writer, string, any) error
}
