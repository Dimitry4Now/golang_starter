package handlers

import (
	"html/template"
	"net/http"
	"log"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)


// LandingHandler handles the landing page request
func LandingHandler(w http.ResponseWriter, r *http.Request) {
    // Get language from query parameter or default to English

    // Fetch localized strings
    data := map[string]string{
        "Title":      Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "login_title"}),
        "Username":   Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "username_label"}),
        "Password":   Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "password_label"}),
        "Login":      Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "login_button"}),
        "CreateAcct": Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "create_account"}),
    	"HeaderTitle": Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "header_title"}),
		"Message":	Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "landing_message"}),
		"Home":       Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "home"}),
        "About":      Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "about"}),
        "Contact":    Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "contact"}),
        "FooterCopyright":      Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "footer_copyright"}),
        "FooterPrivacyPolicy":  Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "footer_privacy_policy"}),
        "FooterTermsService":   Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "footer_terms_service"}),
	}

	// Parse all templates at once
    tmpl := template.Must(template.ParseFiles("templates/header.html", "templates/landing.html", "templates/footer.html"))

    // Execute the landing template, passing the data
    err := tmpl.ExecuteTemplate(w, "landing", data)
    if err != nil {
        http.Error(w, "Error rendering landing template", http.StatusInternalServerError)
        log.Println(err)
    }
}

