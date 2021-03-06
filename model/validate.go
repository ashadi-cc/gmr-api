package model

import (
	"log"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

func init() {
	en := en.New()
	uni = ut.New(en, en)
	validate = validator.New()

	var ok bool
	trans, ok = uni.GetTranslator("en")
	if ok {
		err := en_translations.RegisterDefaultTranslations(validate, trans)
		if err != nil {
			log.Fatalln(err)
		}
	}

}
