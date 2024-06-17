package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// Define a new command-line flag with the name 'addr', a default value of ":4000"
	// and som eshort help text explaingin what the flag controls. The value of the
	// flag will be stored in the addr variable at runtime.
	addr := flag.String("addr", ":4000", "HTTP network address")

	// Importantly, we use the flag.Parse() function to parse the command-line flag.
	// This reads in the command-line flag value and assigns it to the addr variable.
	// You need to call this *before* you use the addr variable otherwise it will 
	// always contain the default value of ":4000". If any errors are encountered
	// during parsing the applicatioin will be terminated.
	flag.Parse()

	// Use the http.NewServeMux() function to initialize
	// a new servemux, then register the home function
	// as the handler for the "/" URL pattern.
	mux := http.NewServeMux()

	// Create a file server which serves files out of the "./ui/static" directory.
	// Note that the path given to the http.Dir function is relatice to the directory root.
	fileServer := http.FileServer(http.Dir("./ui/static"))

	// Use the mux.Handle() functioin to register the file server as the handler for
	// all URL paths that start with "/static/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// Prefix the route patterns with the required HTTP method
	// (for now, we will restrict all three routes to acting on GET requests).
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	// Create the new route, which is restricted to POST request only.
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	// The value returned from the flag.String() function is a pointer to the flag
	// value, not the value itself. So in this code, that means the addr variable
	// is actually a pointer, and we neet to dereference it (i.e. prefix it with
	// * symbol) before using it. Note that we're using the log.Printf() function
	// to interpolate the address with the log message.
	log.Printf("starting server on %s", *addr)

	// Use the http.ListenAndServe() function to start the web server
	// We pass in two parameters:
	// the TCP network address to listen on (in this case ":4000")
	// and the servemux we just created. If http.ListenAndServe() returns
	// an error we use the log.Fatal() function to log the error message and exit.
	// Note that any error returned by http.ListenAndServe() is always non-nil
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
