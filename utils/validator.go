package utils

import (
    "regexp"
    "github.com/go-playground/validator/v10"
)

func ValidatePassword(fl validator.FieldLevel) bool {
    password := fl.Field().String()

    if len(password) < 8 || len(password) > 20 {
        return false
    }
    
    var (
        hasNumber      = regexp.MustCompile(`[0-9]`).MatchString
        hasUppercase   = regexp.MustCompile(`[A-Z]`).MatchString
        hasLowercase   = regexp.MustCompile(`[a-z]`).MatchString
        hasSpecialChar = regexp.MustCompile(`[!@#\$%\^&\*\(\)\-_+=]`).MatchString
    )

    return hasNumber(password) && hasUppercase(password) && hasLowercase(password) && hasSpecialChar(password)
}

func NewValidator() *validator.Validate {
    v := validator.New()

    v.RegisterValidation("password", ValidatePassword)

    return v
}