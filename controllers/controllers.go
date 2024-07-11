package controllers

import (
	"context"
	"echoAPI/config"
	"echoAPI/models"
	"log"
	"net/http"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
)

func GetBooks(c echo.Context) error {
    author := c.QueryParam("author")
    category := c.QueryParam("category")
    log.Printf("Author: %s, Category: %s", author, category)

    query := "SELECT id, author, post, category FROM books"
    var books []models.Book

    conditions := []string{}
    params := []interface{}{}
    if author != "" {
        conditions = append(conditions, "author=$1")
        params = append(params, author)
    }
    if category != "" {
        conditions = append(conditions, "category=$2")
        params = append(params, category)
    }

    if len(conditions) > 0 {
        query += " WHERE " + strings.Join(conditions, " AND ")
    }

    log.Printf("Query: %s", query)

    rows, err := config.DB.Query(context.Background(), query, params...)
    if err != nil {
        log.Printf("Error querying books: %v", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch books"})
    }
    defer rows.Close()

    for rows.Next() {
        var book models.Book
        err := rows.Scan(&book.ID, &book.Author, &book.Post, &book.Category)
        if err != nil {
            log.Printf("Error scanning book: %v", err)
            return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch books"})
        }
        books = append(books, book)
    }

    log.Printf("Books: %v", books)
    return c.JSON(http.StatusOK, books)
}

// func Login(c echo.Context) error {
//     username := c.FormValue("username")
//     password := c.FormValue("password")

//     var user models.User
//     query := "SELECT id, username, password FROM users WHERE username=$1"
//     err := config.DB.QueryRow(context.Background(), query, username).Scan(&user.ID, &user.Username, &user.Password)
//     if err != nil {
//         return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid username or password"})
//     }

//     if user.Password != password {
//         return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid username or password"})
//     }

//     sessionID := uuid.Must(uuid.NewV7()).String()
//     cookie := new(http.Cookie)
//     cookie.Name = "session_id"
//     cookie.Value = sessionID
//     cookie.Path = "/"
//     c.SetCookie(cookie)

//     return c.JSON(http.StatusOK, map[string]string{"message": "Logged in successfully"})
// }

// func Logout(c echo.Context) error {
//     cookie := new(http.Cookie)
//     cookie.Name = "session_id"
//     cookie.Value = ""
//     cookie.Path = "/"
//     cookie.MaxAge = -1
//     c.SetCookie(cookie)

//     return c.JSON(http.StatusOK, map[string]string{"message": "Logged out successfully"})
// }

func GetBookByID(c echo.Context) error {
    id := c.Param("id")

    var book models.Book
    query := "SELECT id, author, post, category FROM books WHERE id=$1"
    err := config.DB.QueryRow(context.Background(), query, id).Scan(&book.ID, &book.Author, &book.Post, &book.Category)
    if err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "Book not found"})
    }

    return c.JSON(http.StatusOK, book)
}

// POST /api/books
func CreateBook(c echo.Context) error {
    book := new(models.Book)
    if err := c.Bind(book); err != nil {
        return err
    }

    book.ID = uuid.Must(uuid.NewV7())
    query := "INSERT INTO books (id, author, post, category) VALUES ($1, $2, $3, $4)"
    _, err := config.DB.Exec(context.Background(), query, book.ID, book.Author, book.Post, book.Category)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create book"})
    }

    return c.JSON(http.StatusCreated, book)
}

// PUT /api/books/:id
func UpdateBook(c echo.Context) error {
    id := c.Param("id")
    book := new(models.Book)
    if err := c.Bind(book); err != nil {
        return err
    }

    query := "UPDATE books SET author=$1, post=$2, category=$3 WHERE id=$4"
    _, err := config.DB.Exec(context.Background(), query, book.Author, book.Post, book.Category, id)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update book"})
    }

    return c.JSON(http.StatusOK, book)
}

// DELETE /api/books/:id
func DeleteBook(c echo.Context) error {
    id := c.Param("id")

    query := "DELETE FROM books WHERE id=$1"
    _, err := config.DB.Exec(context.Background(), query, id)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete book"})
    }

    return c.NoContent(http.StatusNoContent)
}
