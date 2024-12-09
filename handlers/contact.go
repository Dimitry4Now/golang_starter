package handlers

import (
	"html/template"
	"net/http"
	"log"
)


// LandingHandler handles the landing page request
func ContactHandler(w http.ResponseWriter, r *http.Request) {
    // Get language from query parameter or default to English

    // Fetch localized strings
    data := GetLocalizedData() 

	// Parse all templates at once
    tmpl := template.Must(template.ParseFiles("templates/header.html", "templates/contact.html", "templates/footer.html"))

    // Execute the landing template, passing the data
    err := tmpl.ExecuteTemplate(w, "contact",data)
    if err != nil {
        http.Error(w, "Error rendering landing template", http.StatusInternalServerError)
        log.Println(err)
    }
}

