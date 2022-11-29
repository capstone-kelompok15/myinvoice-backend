package validatorutils

import (
	"fmt"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	validator_en "github.com/go-playground/validator/v10/translations/en"
)

type Validator struct {
	*validator.Validate
	ut.Translator
}

func New() (*Validator, error) {
	validator := validator.New()

	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	err := validator_en.RegisterDefaultTranslations(validator, trans)
	if err != nil {
		return nil, err
	}

	return &Validator{
		Validate:   validator,
		Translator: trans,
	}, nil
}

func (v Validator) TranslateValidatorError(err error) (errs []string) {
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(v.Translator))
		errs = append(errs, translatedErr.Error())
	}

	return errs
}
