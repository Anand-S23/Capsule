package models

type RegisterDto struct {
    Name     string
    Email    string
    Phone    string
    Password string
}

type LoginDto struct {
    Email    string
    Password string
}

