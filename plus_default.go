package httpaccesslog


import (
	"io"
	"net/http"
)


func PlusDefault(fn WriterFunc) WriterFunc {

	return func(writer io.Writer, w *ResponseWriter, r *http.Request, trace *Trace) {
		fn(writer, w, r, trace)

		writer.Write(space)

		DefaultAccessLogWriter(writer, w, r, trace)
	}

}
