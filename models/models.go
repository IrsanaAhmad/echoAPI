package models

import (
    "github.com/gofrs/uuid"
)

type Book struct {
    ID       uuid.UUID `json:"id"`
    Author   string    `json:"author"`
    Post     string    `json:"post"`
    Category string    `json:"category"`
}

type User struct {
    ID       uuid.UUID `json:"id"`
    Username string    `json:"username"`
    Password string    `json:"password"`
    RoleID   uuid.UUID `json:"role_id"`
}

type Role struct {
    ID   uuid.UUID `json:"id"`
    Name string    `json:"name"`
}
