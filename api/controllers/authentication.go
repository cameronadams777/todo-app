package controllers

import (
	"api/config"
	"api/database"
	"api/models"
	"fmt"

	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	type LoginInput struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var input LoginInput

	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Error on login request.", "data": err})
		return
	}

	var user models.User

	database.DB.Where(&models.User{Email: input.Email}).First(&user)

	if user.ID == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "User not found.", "data": nil})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "error": "Incorrect email and password combination."})
		return
	}

	expirationTime := time.Now().Add(time.Hour * 24)

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: expirationTime.Unix(),
	})

	token, err := claims.SignedString([]byte(config.Config("AUTH_TOKEN_SECRET")))

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Error on login request.", "data": err})
		return
	}

	c.SetCookie("ucid", token, int(expirationTime.Unix()), "", "", false, true)

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "User logged in.", "data": token})
}

func Register(c *gin.Context) {
	type RegisterInput struct {
		FirstName            string `json:"firstName" binding:"required"`
		LastName             string `json:"lastName" binding:"required"`
		Email                string `json:"email" binding:"required"`
		Password             string `json:"password" binding:"required"`
		PasswordConfirmation string `json:"passwordConfirmation" binding:"required"`
	}

	var input RegisterInput

	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Error on register request.", "data": err})
		return
	}

	if input.Password != input.PasswordConfirmation {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Password and Confirmation do not match.", "data": nil})
		return
	}

	var existingUser models.User

	database.DB.Where(&models.User{Email: input.Email}).Find(&existingUser)

	if existingUser.ID != 0 {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Error on register request.", "data": nil})
		return
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 14)

	user := models.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  string(password),
	}

	database.DB.Create(&user)

	expirationTime := time.Now().Add(time.Hour * 24)

	// TODO: Migrate to RegisteredClaims
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: expirationTime.Unix(),
	})

	token, err := claims.SignedString([]byte(config.Config("AUTH_TOKEN_SECRET")))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Error on register request.", "data": err})
		return
	}

	c.SetCookie("ucid", token, int(expirationTime.Unix()), "", "", false, true)

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "User registered", "data": gin.H{
		"token":     token,
		"firstName": input.FirstName,
		"lastName":  input.LastName,
		"email":     input.Email,
	}})
}
