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
