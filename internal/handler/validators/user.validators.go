package validators

import "github.com/go-playground/validator/v10"

func ValidUserType() (string, validator.Func) {
	return "userType", func(fl validator.FieldLevel) bool {
		var types = map[string]bool{
			"seller": true,
			"buyer":  true,
		}

		return types[fl.Field().String()]
	}
}
