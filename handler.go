package httpaccesslog


import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)


const (
	traceIDContextKey = `"trace"."id"`
)


var (
	notFoundHandler = http.HandlerFunc(http.NotFound)
)


type WriterFunc func(io.Writer, *ResponseWriter, *http.Request, *Trace)error


// Handler is "middleware" that provides logging and tracing capabilities for your http.Handler.
//
//
// By default access logs will be sent to STDOUT. (I.e., os.Stdout.)
//
// However, one can sent the access logs to any io.Writer, by setting the Writer field.
//
//
// A trace ID will be automagically generated with each HTTP request.
//
// The trace ID will be returned to the HTTP client as the value of the "X-Trace" HTTP
// response header.
//
// The sub-handler can get the value of the trace ID looking up the `"trace"."id"` in
// the context of the *http.Request passed as a parameter to the ServeHTTP method.
type Handler struct {
	Subhandler      http.Handler
	Writer          io.Writer
	AccessLogWriter WriterFunc
}


// ServeHTTP makes Handler fit the http.Handler interface.
func (handler Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var trace Trace
	trace.initialize()


	subhandler := handler.Subhandler
	if nil == subhandler {
		subhandler = notFoundHandler
	}

	writer := handler.Writer
	if nil == writer {
		writer = os.Stdout
	}


	var w2 ResponseWriter = ResponseWriter{w:w}


	ctx := r.Context()
	if nil == ctx {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	newCTX := context.WithValue(ctx, traceIDContextKey, string(trace.ID[:]))


	r2 := r.WithContext(newCTX)


	subhandler.ServeHTTP(&w2, r2)

	trace.EndTime = time.Now()

	writeLog := handler.AccessLogWriter
	if nil == writeLog {
		writeLog = DefaultAccessLogWriter
	}
	writeLog(writer, &w2, r, &trace)
}
