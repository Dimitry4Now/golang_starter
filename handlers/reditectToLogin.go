package handlers

import (
	"html/template"
	"net/http"
	"log"
)

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
    // Fetch localized strings
    data := GetLocalizedData()

	// Parse all templates at once
    tmpl := template.Must(template.ParseFiles("templates/header.html", "templates/login.html", "templates/footer.html"))

    // Execute the landing template, passing the data
    err := tmpl.ExecuteTemplate(w, "login", data)
    if err != nil {
        http.Error(w, "Error rendering landing template", http.StatusInternalServerError)
        log.Println(err)
    }
}

