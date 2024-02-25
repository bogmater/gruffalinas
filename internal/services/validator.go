package services

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

// Validator provides validation mainly validating structs within the web context
type Validator struct {
	// validator stores the underlying validator
	Vdtor *validator.Validate
	Trans *ut.Translator
}

// NewValidator creats a new Validator
func NewValidator() *Validator {

	v := validator.New(validator.WithRequiredStructEnabled())
	en := en.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(v, trans)

	return &Validator{Vdtor: v, Trans: &trans}
}

// Validate validates a struct
func (v *Validator) Validate(i any) error {
	if err := v.Vdtor.Struct(i); err != nil {
		return err
	}
	return nil
}

func (v *Validator) Struct(s interface{}) validator.ValidationErrorsTranslations {
	err := v.Vdtor.Struct(s)
	errs := err.(validator.ValidationErrors).Translate(*v.Trans)

	return errs
}
