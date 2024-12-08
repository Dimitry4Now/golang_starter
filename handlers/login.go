package handlers

	import (
		"fmt"
		"net/http"
	)

	func LoginHandler(w http.ResponseWriter, r *http.Request) {
	    if r.Method == http.MethodPost {
		    username := r.FormValue("username")
		    if username == "" {
			    fmt.Fprint(w, "Please provide a username!")
			    return
		    }
		    fmt.Fprintf(w, "Hello, %s! You are logged in.", username)
	    }
    }
