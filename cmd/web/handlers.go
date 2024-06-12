package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// Define a home handler function which writes a byte slice
// containing "Hello from Snippetbox" as the response body
func home(w http.ResponseWriter, r *http.Request) {
	// Use the Header().Add() method to add a 'Server: Go' header to the
	// response header map. The first parameter is the header name, and
	// the second paramter is the header value.
	// Important: Any changes you make to the response header map _after_ calling
	// w.WriteHeader() or w.Write() will have no effect on the headers
	// that the user receives
	w.Header().Add("Server", "Go")

	w.Write([]byte("Hello from Snippetbox\n"))
}

// Add a snippet view handler function
func snippetView(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the id wildcard from the request using r.PathValue()
	// an dtry to convert it to an integer using the strconv.Atoi() function.
	// If it can't be converted to an integer, or the value is less than 1, 
	// we ruturn 404 page not found response
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w,r)
		return
	}

	// http.ResponseWriter satisfies the io.Writer interface
	// That means you can use any standard library function which accepts
	// an io.Writer parameter to write plain-text response bodies
	fmt.Fprintf(w, "Display a specific snipped with ID %d...\n", id)
}

// Add a snippet create handler function
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet...\n"))
}

// Add a snipped create POST handler function
func snippetCreatePost(w http.ResponseWriter, r *http.Request){
	// Use the w.WriteHeader() method to send a 201 status code.
	w.WriteHeader(http.StatusCreated)

	// Then w.Write() method to write the response body as normal.
	w.Write([]byte("Save a new snippet... \n"))
}