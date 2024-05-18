package users

import (
	"database/sql"
	b64 "encoding/base64"
	"net/http"

	"example.com/api/db"
	"github.com/gin-gonic/gin"
)

type Users struct {
	user_id  int    `json:"user_id"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func UsersHello(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"msg": "HelloFromUsers",
	})
}

func Signup(c *gin.Context) {
	var jsonData Users
	c.ShouldBindJSON(&jsonData)
	encoded_pass := b64.StdEncoding.EncodeToString([]byte(jsonData.Password))
	sql_statement := "INSERT INTO users (username, password) VALUES (?, ?)"
	_, err := db.DB.Exec(sql_statement, jsonData.Username, encoded_pass)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Could not sign up user",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User added successfully",
	})
}

func Login(c *gin.Context) {
	var jsonData Users
	var user Users
	c.ShouldBindJSON(&jsonData)
	encoded_pass := b64.StdEncoding.EncodeToString([]byte(jsonData.Password))
	row := db.DB.QueryRow("SELECT username, password FROM users where username = ? and password = ?", jsonData.Username, encoded_pass)
	err_scan := row.Scan(&user.Username, &user.Password)
	if sql.ErrNoRows != err_scan {
		c.JSON(http.StatusOK, gin.H{
			"username": &user.Username,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
	}
}
