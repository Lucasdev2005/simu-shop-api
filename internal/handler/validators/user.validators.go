package validators

import (
	"simushop/internal/enum"

	"github.com/go-playground/validator/v10"
)

func ValidUserType() (string, validator.Func) {
	return "userType", func(fl validator.FieldLevel) bool {
		var types = map[string]bool{
			enum.UserTypeEnum.BUYER:  true,
			enum.UserTypeEnum.SELLER: true,
		}

		return types[fl.Field().String()]
	}
}

func UptateValidUserType() (string, validator.Func) {
	return "updateUserType", func(fl validator.FieldLevel) bool {
		var types = map[string]bool{
			enum.UserTypeEnum.BUYER:  true,
			enum.UserTypeEnum.SELLER: true,
			"":                       true,
		}

		return types[fl.Field().String()]
	}
}
