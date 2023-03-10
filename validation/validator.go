package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/happyhippyhippo/slate-rest/envelope"
)

// Validator is a function type used to define a calling interface of
// function responsible to validate an instance of a structure and return
// an initialized response envelope with the founded error
type Validator func(val interface{}) (*envelope.Envelope, error)

// NewValidator instantiates a new validation function
func NewValidator(
	validate *validator.Validate,
	parser IParser,
) (Validator, error) {
	// check validate argument reference
	if validate == nil {
		return nil, errNilPointer("validate")
	}
	// check parser argument reference
	if parser == nil {
		return nil, errNilPointer("parser")
	}
	// return the validation method instance
	return func(value interface{}) (*envelope.Envelope, error) {
		// check the value argument reference
		if value == nil {
			return nil, errNilPointer("value")
		}
		// validate the given structure
		if errs := validate.Struct(value); errs != nil {
			// compose the response envelope with the parsed validation error
			return parser.Parse(value, errs.(validator.ValidationErrors))
		}
		return nil, nil
	}, nil
}
