package util

import (
	"errors"
	"reflect"
	"strings"

	english "github.com/go-playground/locales/en_US"
	ut "github.com/go-playground/universal-translator"
	govalidate "github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/en"
)

var globalValidator *Validator

func init() {
	globalValidator = NewValidator()
}

func ValidateStruct(val interface{}) []error {
	return globalValidator.Struct(val)
}

func ValidtateValue(val interface{}, tag string) []error {
	return globalValidator.Value(val, tag)
}

type Validator struct {
	validator  *govalidate.Validate
	translator *ut.UniversalTranslator
}

func NewValidator() *Validator {
	eng := english.New()
	translator := ut.New(eng, eng)
	transEn, _ := translator.GetTranslator("en")

	validator := govalidate.New()

	validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		tagValues := strings.Split(fld.Tag.Get("json"), ",")

		name := ""

		for _, v := range tagValues {
			if v == "" || v == "-" || v == "omitempty" {
				continue
			}
			name = v
		}

		if name == "" {
			name = fld.Name
		}

		return name
	})

	if err := en.RegisterDefaultTranslations(validator, transEn); err != nil {
		panic(err)
	}
	return &Validator{
		validator:  validator,
		translator: translator,
	}
}

func (v *Validator) Struct(val interface{}) []error {
	err := v.validator.Struct(val)

	// if val is nil or slice or array, panic will be occured.
	if _, ok := err.(*govalidate.InvalidValidationError); ok {
		panic(err)
	}

	if fieldErrs, ok := err.(govalidate.ValidationErrors); ok {
		return v.parseValidationErrors(fieldErrs)
	}
	return nil
}

func (v *Validator) Value(val interface{}, tag string) []error {
	err := v.validator.Var(val, tag)

	// if val is nil or slice or array, panic will be occured.
	if _, ok := err.(*govalidate.InvalidValidationError); ok {
		panic(err)
	}

	if fieldErrs, ok := err.(govalidate.ValidationErrors); ok {
		return v.parseValidationErrors(fieldErrs)
	}

	return nil
}

func (v *Validator) parseValidationErrors(fieldErrs govalidate.ValidationErrors) []error {
	trans, _ := v.translator.GetTranslator("en")

	errs := []error{}

	for _, field := range fieldErrs {
		errs = append(errs, errors.New(field.Translate(trans)))
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}
