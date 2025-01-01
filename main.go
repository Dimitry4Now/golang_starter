package main

import (
	"net/http"
	"go_htmx_essentials/handlers"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"encoding/json"
	"go_htmx_essentials/logging"
)

func main() {
	logging.InitLogger()

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	handlers.Bundle = i18n.NewBundle(language.English)
    handlers.Bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
    handlers.Bundle.MustLoadMessageFile("i18n/en.json")
    handlers.Bundle.MustLoadMessageFile("i18n/mk.json")

	if handlers.Localizer == nil {
    	handlers.Localizer = i18n.NewLocalizer(handlers.Bundle, "en")
	}

	logging.Logger.Info("Server Started Successfully...")

	http.HandleFunc("/", handlers.LandingHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/loginForm", handlers.RedirectHandler)
	http.HandleFunc("/change-language", handlers.ChangeLanguageHandler)
	http.HandleFunc("/about", handlers.AboutUsHandler)
	http.HandleFunc("/contact", handlers.ContactHandler)
	
	logging.Logger.Fatal(http.ListenAndServe(":8000", nil))
}
