package utils

import (
	"fmt"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

func init() {
	enLocale := en.New()
	uni = ut.New(enLocale, enLocale)
	trans, _ = uni.GetTranslator("en")
	validate = validator.New()
	enTranslations.RegisterDefaultTranslations(validate, trans)
}

func Validate[T any](data T) map[string]string {
	err := validate.Struct(data)
	if err == nil {
		return nil
	}

	res := map[string]string{}
	if err != nil {
		fmt.Print(err)
		for _, v := range err.(validator.ValidationErrors) {
			res[v.StructField()] = v.Translate(trans)
		}
	}
	return res
}
