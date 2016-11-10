package httpaccesslog


import (
	"errors"
	"net/http"
)


const (
	traceIDContextKey = `"trace"."id"`
)


var (
	errBadTraceID  = errors.New("Bad Trace ID")
	errNilContext  = errors.New("Nil Context")
	errNilTraceID  = errors.New("Nil Trace ID")
)


// TraceID returns the trace ID loaded in the *http.Request.
func TraceID(r *http.Request) (string, error) {

	ctx := r.Context()
	if nil == ctx {
		return "", errNilContext
	}

	traceIDInterface := ctx.Value(traceIDContextKey)
	if nil == traceIDInterface {
		return "", errNilTraceID
	}

	traceID, ok := traceIDInterface.(string)
	if !ok {
		return "", errBadTraceID
	}

	return traceID, nil
}


// MustTraceID is like TraceID, but panic()s if there is an error.
func MustTraceID(r *http.Request) string {

	traceID, err := TraceID(r)
	if nil != err {
		panic(err)
	}

	return traceID
}
