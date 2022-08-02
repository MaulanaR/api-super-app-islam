package app

import (
	idLocale "github.com/go-playground/locales/id"
	idTranslation "github.com/go-playground/validator/v10/translations/id"
	"grest.dev/grest/validator"
)

func configureValidator() {
	validator.Configure()
	validator.AddTranslator("id", idLocale.New(), idTranslation.RegisterDefaultTranslations)
	validator.AddTranslator("id-ID", idLocale.New(), idTranslation.RegisterDefaultTranslations)
}

func IsValid(v interface{}, tag string) bool {
	return validator.IsValid(v, tag)
}

func Validate(lang string, v interface{}) error {
	return validator.Validate(lang, v)
}
