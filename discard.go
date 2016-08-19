package httpaccesslog


import (
	"io"
	"net/http"
)


func DiscardAccessLogWriter(io.Writer, *ResponseWriter, *http.Request, *Trace) {
	// Nothing here.
}
