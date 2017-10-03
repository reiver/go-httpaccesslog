/*
Package httpaccesslog provides HTTP "middleware" that provides access log generation capabilities.

Simple Example:

	var subhandler http.Handler
	
	// ...
	
	var httpOverlordHandler http.Handler = httpaccesslog.Handler{
		Subhandler: subhandler,
		Writer:     os.Stdout,
	}

*/
package httpaccesslog

