package validate

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"strings"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

func Struct(s interface{}) error {
	en := en.New()
	uni = ut.New(en, en)

	trans, _ := uni.GetTranslator("en")

	validate = validator.New()
	en_translations.RegisterDefaultTranslations(validate, trans)

	err := validate.Struct(s)

	var messages []string
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			messages = append(messages, notification(e.Field(), e.Translate(trans)))
		}

		return errors.New(strings.Join(messages, "|"))
	}

	return nil
}

func ErrorMessage(field string, message string) error {
	return errors.New(notification(field, message))
}

func notification(field string, message string) string {
	m := struct {
		Message string `json:"message"`
		Field   string `json:"field"`
	}{
		Message: message,
		Field:   field,
	}

	out, _ := json.Marshal(m)

	return string(out)
}
