package validation

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
	"github.com/matheusdcunha/aprendendo-go/meu-primeiro-crud-go/src/config/resterr"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validationError error) *resterr.RestErr {

	var jsonErr *json.UnsupportedTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validationError, &jsonErr) {
		return resterr.NewBadRequestError("invalid field type")
	}

	if errors.As(validationError, &jsonValidationError) {
		errorsCauses := []resterr.Causes{}

		for _, e := range validationError.(validator.ValidationErrors) {
			cause := resterr.Causes{
				Message: strings.ToLower(e.Translate(transl)),
				Field:   strings.ToLower(e.Field()),
			}

			errorsCauses = append(errorsCauses, cause)
		}
		return resterr.NewBadRequestValidationError("some fields are invalid", errorsCauses)
	}

	return resterr.NewBadRequestError("error trying to convert fields")
}
