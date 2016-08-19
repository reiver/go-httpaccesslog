package httpaccesslog


import (
	"io"
	"net/http"
	"os"
	"time"
)


var (
	notFoundHandler = http.HandlerFunc(http.NotFound)
)


type WriterFunc func(io.Writer, *ResponseWriter, *http.Request, *Trace)


type Handler struct {
	Subhandler      http.Handler
	Writer          io.Writer
	AccessLogWriter WriterFunc
}


func (handler Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var trace Trace
	trace.BeginTime = time.Now()


	subhandler := handler.Subhandler
	if nil == subhandler {
		subhandler = notFoundHandler
	}

	writer := handler.Writer
	if nil == writer {
		writer = os.Stdout
	}

	var w2 ResponseWriter = ResponseWriter{w:w}

	subhandler.ServeHTTP(&w2, r)

	trace.EndTime = time.Now()

	writeLog := handler.AccessLogWriter
	if nil == writeLog {
		writeLog = DefaultAccessLogWriter
	}
	writeLog(writer, &w2, r, &trace)
}
