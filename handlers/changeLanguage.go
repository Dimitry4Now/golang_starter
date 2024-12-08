package handlers

import (
	"net/http"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"encoding/json"
	"golang.org/x/text/language"
	"log"
)


// ChangeLanguageHandler changes the language based on the selected flag
func ChangeLanguageHandler(w http.ResponseWriter, r *http.Request) {
	lang := r.URL.Query().Get("lang")
	if lang == "" {
		http.Error(w, "Language not specified", http.StatusBadRequest)
		return
	}

	Bundle := i18n.NewBundle(language.English)
	Bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	// Load translation files
	Bundle.MustLoadMessageFile("i18n/en.json")
	Bundle.MustLoadMessageFile("i18n/mk.json")
	// Update the global localizer with the new language
	Localizer = i18n.NewLocalizer(Bundle, lang)

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

	// Debug
	log.Printf("Inside changeLanguage.go")
	log.Printf("Language selected: %s", lang)
	log.Printf("Localized Title: %s", data["Title"])
	
	// Redirect back to the referring page
	referer := r.Header.Get("Referer")
	if referer == "" {
		// If there's no referer, redirect to the homepage
		referer = "/"
	}
	http.Redirect(w, r, referer, http.StatusSeeOther)
}
