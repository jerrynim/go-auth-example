// controllers/books.go

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jerrynim/leave/models"
	"golang.org/x/crypto/bcrypt"
)


type CreateUserBody struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Firstname string `json:"firstname" binding:"required"`
	Lastname string  `json:"lastname" binding:"required"`
  }

// POST /users
// Create new users
func CreateUser(c *gin.Context) {
	// Validate input
	var body CreateUserBody
	
	if err := c.ShouldBindJSON(&body); err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	  c.AbortWithStatus(400)
	  return
	}
	
	//* 비밀번호 bcrypt 해쉬 하기
	hash,err := bcrypt.GenerateFromPassword([]byte(body.Password),bcrypt.DefaultCost)
	
	if err !=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	  c.Abort()
	  return
	}
	// Create ㅕ
	user := models.User{Email:body.Email,Password:string(hash),
		 FirstName:body.Firstname,LastName:body.Lastname}
	models.DB.Create(&user)
  
	c.JSON(http.StatusOK, gin.H{"data": user})
  }

