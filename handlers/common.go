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
        "NoAcct":      Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "no_account"}),
        "CreateHere":      Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "create_here"}),
        "HeaderTitle":     Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "header_title"}),
        "Message":         Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "landing_message"}),
        "Home":            Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "home"}),
        "About":           Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "about"}),
        "Contact":         Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "contact"}),
        "FooterCopyright": Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "footer_copyright"}),
        "FooterPrivacyPolicy": Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "footer_privacy_policy"}),
        "FooterTermsService":  Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "footer_terms_service"}),
        "Project":  Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "project"}),
        "ContactUs":  Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "contact_us"}),
        "Subject":  Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "subject"}),
        "SubjectLabel":  Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "subject_label"}),
        "Submit":  Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "submit"}),
        "ContactMessage":  Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "contact_message"}),
        "Mission":  Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "mission"}),
        "Vision":  Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "vision"}),
        "Values":  Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "values"}),
        "Team":  Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "team"}),
        "ContactMessageLabel":  Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "contact_message_label"}),
    }
}
