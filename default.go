package httpaccesslog


import (
	"io"
	"net/http"
	"strconv"
)


var (
	space = []byte{' '}
	endl  = []byte{'\n'}
)


func DefaultAccessLogWriter(writer io.Writer, w *ResponseWriter, r *http.Request, trace *Trace) {

	io.WriteString(writer, "remote-host=")
	io.WriteString(writer, strconv.Quote(r.RemoteAddr) ) //@TODO: This is inefficient.

	writer.Write(space)

	io.WriteString(writer, "request-method=")
	io.WriteString(writer, strconv.Quote(r.Method))

	writer.Write(space)

	io.WriteString(writer, "request-uri=")
	io.WriteString(writer, strconv.Quote(r.URL.String())) //@TODO: This is inefficient.

	writer.Write(space)

	io.WriteString(writer, "request-protocol=")
	io.WriteString(writer, strconv.Quote(r.Proto)) //@TODO: This is inefficient.

	writer.Write(space)

	io.WriteString(writer, "response-status-code=\"")
	io.WriteString(writer, strconv.FormatInt(int64(w.StatusCode), 10)) //@TODO: This is inefficient.
	io.WriteString(writer, "\"")

	writer.Write(space)

	io.WriteString(writer, "response-size=\"")
	io.WriteString(writer, strconv.FormatInt(int64(w.BodySize), 10)) //@TODO: This is inefficient.
	io.WriteString(writer, "\"")

	writer.Write(space)

	io.WriteString(writer, "trace.begin-time=")
	io.WriteString(writer, strconv.Quote(trace.BeginTime.String())) //@TODO: This is inefficient.

	writer.Write(space)

	io.WriteString(writer, "trace.end-time=")
	io.WriteString(writer, strconv.Quote(trace.EndTime.String())) //@TODO: This is inefficient.

	writer.Write(endl)
}
