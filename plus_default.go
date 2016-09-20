package httpaccesslog

import (
	"io"
	"net/http"
)

// PlusDefault returns a WriterFunc that calls `fn` first, to write the beginning of the log,
// and then calls the DefaultAccessLogWriter.
//
// If the WriterFunc this wraps returns an error, no log will be written.
func PlusDefault(fn WriterFunc) WriterFunc {

	return func(writer io.Writer, w *ResponseWriter, r *http.Request, trace *Trace) error {
		if err := fn(writer, w, r, trace); nil != err {
			return err
		}

		writer.Write(space)

		if err := DefaultAccessLogWriter(writer, w, r, trace); nil != err {
			return err
		}

		return nil
	}

}
