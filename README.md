# go-httpaccesslog

Package **httpaccesslog** provides tools for creating HTTP access-logs, for the Go programming language.

Package **httpaccesslog** provides HTTP "middleware" that provides HTTP access-log generation capabilities.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-httpaccesslog

[![GoDoc](https://godoc.org/github.com/reiver/go-httpaccesslog?status.svg)](https://godoc.org/github.com/reiver/go-httpaccesslog)

## Example

Here is a simple example:

```golang
var subhandler http.Handler

// ...

var httpOverlordHandler http.Handler = httpaccesslog.Handler{
	Subhandler: subhandler,
	Writer:     os.Stdout,
}
```

## Import

To import package **httpaccesslog** use `import` code like the following:
```
import "github.com/reiver/go-httpaccesslog"
```

## Installation

To install package **httpaccesslog** do the following:
```
GOPROXY=direct go get github.com/reiver/go-httpaccesslog
```

## Author

Package **httpaccesslog** was written by [Charles Iliya Krempeaux](http://reiver.link)
