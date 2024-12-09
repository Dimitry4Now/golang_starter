package handlers

import "github.com/nicksnyder/go-i18n/v2/i18n"

// Shared localizer
// Exported variables
var Bundle *i18n.Bundle
var Localizer *i18n.Localizer

// Function to return the localized strings map
func GetLocalizedData() map[string]string {
    return map[string]string{
        "Title":           Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "login_title"}),
        "Username":        Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "username_label"}),
        "Password":        Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "password_label"}),
        "Login":           Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "login_button"}),
        "CreateAcct":      Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "create_account"}),
        "HeaderTitle":     Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "header_title"}),
        "Message":         Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "landing_message"}),
        "Home":            Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "home"}),
        "About":           Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "about"}),
        "Contact":         Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "contact"}),
        "FooterCopyright": Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "footer_copyright"}),
        "FooterPrivacyPolicy": Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "footer_privacy_policy"}),
        "FooterTermsService":  Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "footer_terms_service"}),
    }
}
