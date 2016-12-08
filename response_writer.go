package httpaccesslog


import (
	"net/http"
)


type ResponseWriter struct {
	w http.ResponseWriter

	StatusCode int
	BodySize   int
}


func (rw *ResponseWriter) Header() http.Header {
	return rw.w.Header()
}


func (rw *ResponseWriter) Write(p []byte) (int, error) {
	const undefinedStatusCode = 0
	if undefinedStatusCode == rw.StatusCode {
		rw.StatusCode = http.StatusOK
	}

	size, err := rw.w.Write(p)

	rw.BodySize += size

	return size, err
}


func (rw *ResponseWriter) WriteHeader(statusCode int) {
	rw.w.WriteHeader(statusCode)
	rw.StatusCode = statusCode
}


func (rw *ResponseWriter) Flush() {
	if flusher, ok := rw.w.(http.Flusher); ok {
		flusher.Flush()
	}
}


func (rw *ResponseWriter) CloseNotify() <-chan bool {
	return rw.w.(http.CloseNotifier).CloseNotify()
}
