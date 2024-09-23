package pkg

import (
    "regexp"
    "github.com/go-playground/validator/v10"
)

func ValidatePassword(fl validator.FieldLevel) bool {
    password := fl.Field().String()

    // Password must be between 8 and 20 characters long.
    if len(password) < 8 || len(password) > 20 {
        return false
    }


    // Check if the password contains at least one number, one uppercase letter,
    // one lowercase letter, and one special character.
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

    // Register the custom validation function for the "password"
    v.RegisterValidation("password", ValidatePassword)

    return v
}
