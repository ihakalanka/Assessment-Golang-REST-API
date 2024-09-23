package dto

type RegisterUserInput struct {
    Name     string `json:"name" validate:"required"`
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password,omitempty" validate:"required,password"`
    Role     string `json:"role"`
}

type UpdateUserInput struct {
    Name     string `json:"name" validate:"required,min=3,max=100"`
    Email    string `json:"email" validate:"required,email"`   
    Password string `json:"password,omitempty" validate:"required,password"`
    Role     string `json:"role"`
}

type UserDTO struct {
    ID        uint   `json:"id"`
    Name     string `json:"name"`
    Email     string `json:"email"`
    Role     string `json:"role"`
}